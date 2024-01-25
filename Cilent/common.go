package Cilent

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	math_rand "math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"wechatwebapi/Cilent/device"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//Hkdf_Expand
func Hkdf_Expand(h func() hash.Hash, prk, info []byte, outLen int) []byte {
	out := []byte{}
	T := []byte{}
	i := byte(1)
	for len(out) < outLen {
		block := append(T, info...)
		block = append(block, i)

		h := hmac.New(h, prk)
		h.Write(block)

		T = h.Sum(nil)
		out = append(out, T...)
		i++
	}
	return out[:outLen]
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.New(`Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func readAllIntoMemory(filename string) (content []byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[math_rand.Intn(len(letters))]
	}
	return string(b)
}

func NewAES_GCMEncrypter(key []byte, plaintext []byte, nonce []byte, aad []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, aad)
	return ciphertext
}

func NewAES_GCMDecrypter(key []byte, ciphertext []byte, nonce []byte, aad []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		return nil
	}
	return plaintext
}

func RqtCalcData(srcdata []byte) int {
	h := md5.New()
	h.Write([]byte(srcdata))
	md5sign := hex.EncodeToString(h.Sum(nil))
	key, _ := hex.DecodeString("6a664d5d537c253f736e48273a295e4f")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(md5sign))
	my_sign := string(mac.Sum(nil))
	randvalue := 1
	index := 0
	temp0 := 0
	temp1 := 0
	temp2 := 0
	for index = 0; index+2 < 20; index++ {
		temp0 = (temp0&0xff)*0x83 + int(my_sign[index])
		temp1 = (temp1&0xff)*0x83 + int(my_sign[index+1])
		temp2 = (temp2&0xff)*0x83 + int(my_sign[index+2])

	}
	result := (temp2<<16)&0x7f0000 | temp0&0x7f | (randvalue&0x1f|0x20)<<24 | ((temp1 & 0x7f) << 8)
	return result

}

//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w, _ := zlib.NewWriterLevel(&in, zlib.DefaultCompression)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func PackSpecialCgi(Data PackSpecialCgiData) []byte {
	header := new(bytes.Buffer)
	header.Write([]byte{0xbf})
	header.Write([]byte{0x02}) //加密模式占坑,默认不压缩走12

	encryptdata := HybridEcdhEncrypt(HybridEcdhEncryptData{
		Src:                        Data.Reqdata,
		Externkey:                  Data.Extenddata,
		HybridEcdhPubkey:           Data.HybridEcdhPubkey,
		HybridEcdhPrivkey:          Data.HybridEcdhPrivkey,
		HybridEcdhInitServerPubKey: Data.HybridEcdhInitServerPubKey,
	})

	cookielen := len(Data.Cookies)
	header.Write([]byte{byte((Data.Encrypttype << 4) + cookielen)})
	binary.Write(header, binary.BigEndian, int32(Data.ClientVersion))
	if Data.Uin != 0 {
		binary.Write(header, binary.BigEndian, int32(Data.Uin))
	} else {
		header.Write([]byte{0x00, 0x00, 0x00, 0x00})
	}

	if len(Data.Cookies) == 0xF {
		header.Write(Data.Cookies)
	}

	header.Write(proto.EncodeVarint(uint64(Data.Cgi)))
	header.Write(proto.EncodeVarint(uint64(len(Data.Reqdata))))
	header.Write(proto.EncodeVarint(uint64(len(Data.Reqdata))))
	header.Write([]byte{0x90, 0x4E, 0x0D, 0x00, 0xFF})
	header.Write(proto.EncodeVarint(uint64(RqtCalcData(encryptdata))))
	header.Write([]byte{0x00})
	lens := len(header.Bytes())<<2 + 2
	header.Bytes()[1] = byte(lens)
	header.Write(encryptdata)
	return header.Bytes()
}

func AesGcmEncryptWithCompress(key []byte, plaintext []byte, nonce []byte, aad []byte) []byte {
	compressData := DoZlibCompress(plaintext)
	//nonce := []byte(randSeq(12)) //获取随机密钥
	encrypt_data := NewAES_GCMEncrypter(key, compressData, nonce, aad)
	outdata := encrypt_data[:len(encrypt_data)-16]
	retdata := new(bytes.Buffer)
	retdata.Write(outdata)
	retdata.Write(nonce)
	retdata.Write(encrypt_data[len(encrypt_data)-16:])
	return retdata.Bytes()
}

func HkdfExpand(h func() hash.Hash, prk, info []byte, outLen int) []byte {
	out := []byte{}
	T := []byte{}
	i := byte(1)
	for len(out) < outLen {
		block := append(T, info...)
		block = append(block, i)

		h := hmac.New(h, prk)
		h.Write(block)

		T = h.Sum(nil)
		out = append(out, T...)
		i++
	}
	return out[:outLen]
}

func CreateDeviceId(s string) string {
	if s == "" {
		s = RandSeq(15)
	}

	h := md5.New()
	h.Write([]byte(s))
	md5string := hex.EncodeToString(h.Sum(nil))
	return "49" + md5string[2:]
}

func AesGcmDecryptWithUncompress(key []byte, ciphertext []byte, aad []byte) []byte {
	ciphertextinput := ciphertext[:len(ciphertext)-0x1c]
	endatanonce := ciphertext[len(ciphertext)-0x1c : len(ciphertext)-0x10]
	data := new(bytes.Buffer)
	data.Write(ciphertextinput)
	data.Write(ciphertext[len(ciphertext)-0x10 : len(ciphertext)])
	decrypt_data := NewAES_GCMDecrypter(key, data.Bytes(), endatanonce, aad)
	if len(decrypt_data) > 0 {
		return DoZlibUnCompress(decrypt_data)
	} else {
		return []byte{}
	}

}

func AesEncrypt(RequestSerialize []byte, key []byte) []byte {
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	blockSize := block.BlockSize()
	RequestSerialize = PKCS5Padding(RequestSerialize, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(key))
	crypted := make([]byte, len(RequestSerialize))
	blockMode.CryptBlocks(crypted, RequestSerialize)

	return crypted
}

func AesDecrypt(body []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(key))
	origData := make([]byte, len(body))
	blockMode.CryptBlocks(origData, body)
	origData = PKCS5UnPadding(origData)
	return origData
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return nil
	}
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func XorEncodeStr(msg string, key uint8) string {
	ml := len(msg)
	pwd := ""
	for i := 0; i < ml; i++ {

		pwd += string(key ^ uint8(msg[i]))

	}
	return pwd
}

func GetClientSeqId(DeviceId string) string {
	return DeviceId + "-" + strconv.Itoa(int(time.Now().Unix()))
}

func CompressAndAes(RequestSerialize []byte, aeskey []byte) []byte {
	compressed := DoZlibCompress(RequestSerialize)
	return AesEncrypt(compressed, aeskey)
}

/*func Pack(Pack PackData) []byte {
	var body []byte
	if Pack.UseCompress {
		switch Pack.Cgi {
		case 138:
			Pack.EncryptType = 13
			mNonce := []byte(RandSeq(12)) //获取随机密钥
			body = AesGcmEncryptWithCompress(Pack.Clientsessionkey,Pack.Reqdata, mNonce, nil)
		default:
			body = CompressAndAes(Pack.Reqdata,Pack.Sessionkey)
		}
	} else {
		switch Pack.Cgi {
		case 138:
			Pack.EncryptType = 13
			mNonce := []byte(RandSeq(12)) //获取随机密钥
			body = AesGcmEncryptWithCompress(Pack.Clientsessionkey,Pack.Reqdata, mNonce, nil)
		default:
			body = AesEncrypt(Pack.Reqdata,Pack.Sessionkey)
		}
	}

	loginecdhkeylen := int32(len(Pack.Loginecdhkey))

	header := new(bytes.Buffer)
	header.Write([]byte{0xbf})
	header.Write([]byte{0x00})
	header.Write([]byte{((Pack.EncryptType << 4) + 0xf)})
	binary.Write(header, binary.BigEndian, int32(Wx_client_version))
	binary.Write(header, binary.BigEndian, int32(Pack.Uin))
	header.Write(Pack.Cookie)
	header.Write(proto.EncodeVarint(uint64(Pack.Cgi)))
	header.Write(proto.EncodeVarint(uint64(len(Pack.Reqdata))))
	header.Write(proto.EncodeVarint(uint64(len(body))))
	header.Write([]byte{0x00, 0x0d}) //占坑
	uinbyte := new(bytes.Buffer)
	binary.Write(uinbyte, binary.BigEndian, Pack.Uin)
	m1 := md5.New()
	m1.Write(uinbyte.Bytes())
	m1.Write(Pack.Loginecdhkey[:loginecdhkeylen])
	md5str := m1.Sum(nil)

	lenprotobuf := new(bytes.Buffer)
	binary.Write(lenprotobuf, binary.BigEndian, int32(len(Pack.Reqdata)))
	m2 := md5.New()
	m2.Write(lenprotobuf.Bytes())
	m2.Write(Pack.Loginecdhkey[:loginecdhkeylen])
	m2.Write(md5str)

	md5str = m2.Sum(nil)
	adler32buffer := new(bytes.Buffer)
	adler32buffer.Write(md5str)
	adler32buffer.Write(Pack.Reqdata)
	//header.Write(proto.EncodeVarint(uint64(LOGIN_RSA_VER)))
	adler32 := crc32.ChecksumIEEE(adler32buffer.Bytes())
	header.Write(proto.EncodeVarint(uint64(adler32)))
	header.Write([]byte{0xFF})                                  //占坑
	header.Write(proto.EncodeVarint(uint64(RqtCalcData(body)))) //占坑
	header.Write([]byte{0x00})                                  //占坑
	if Pack.UseCompress {
		lens := (len(header.Bytes()) << 2) + 1
		header.Bytes()[1] = byte(lens)
	} else {
		lens := (len(header.Bytes()) << 2) + 2
		header.Bytes()[1] = byte(lens)
	}
	header.Write(body)
	return header.Bytes()

}*/

func Pack(src []byte, cgi int, uin uint32, sessionkey, cookies, clientsessionkey, loginecdhkey []byte, encryptType uint8, use_compress bool) []byte {
	len_proto_compressed := len(src)
	var body []byte
	if use_compress {
		if cgi == 138 {
			encryptType = 13
			mNonce := []byte(RandSeq(12)) //获取随机密钥
			body = AesGcmEncryptWithCompress(clientsessionkey, src, mNonce, nil)
		} else {
			body = CompressAndAes(src, sessionkey)
		}
	} else {
		if cgi == 138 {
			encryptType = 13
			mNonce := []byte(RandSeq(12)) //获取随机密钥
			body = AesGcmEncryptWithCompress(clientsessionkey, src, mNonce, nil)
		} else {
			body = AesEncrypt(src, sessionkey)
		}
	}

	loginecdhkeylen := int32(len(loginecdhkey))

	header := new(bytes.Buffer)
	header.Write([]byte{0xbf})
	header.Write([]byte{0x00})
	header.Write([]byte{((encryptType << 4) + 0xf)})
	binary.Write(header, binary.BigEndian, int32(Wx_client_version))
	binary.Write(header, binary.BigEndian, int32(uin))
	header.Write(cookies)
	header.Write(proto.EncodeVarint(uint64(cgi)))

	if use_compress {
		header.Write(proto.EncodeVarint(uint64(len_proto_compressed)))
		header.Write(proto.EncodeVarint(uint64(len(body))))
	} else {
		header.Write(proto.EncodeVarint(uint64(len_proto_compressed)))
		header.Write(proto.EncodeVarint(uint64(len_proto_compressed)))
	}

	header.Write([]byte{0x00, 0x0d}) //占坑
	uinbyte := new(bytes.Buffer)
	binary.Write(uinbyte, binary.BigEndian, uin)
	m1 := md5.New()
	m1.Write(uinbyte.Bytes())
	m1.Write(loginecdhkey[:loginecdhkeylen])
	md5str := m1.Sum(nil)

	lenprotobuf := new(bytes.Buffer)
	binary.Write(lenprotobuf, binary.BigEndian, int32(len(src)))
	m2 := md5.New()
	m2.Write(lenprotobuf.Bytes())
	m2.Write(loginecdhkey[:loginecdhkeylen])
	m2.Write(md5str)

	md5str = m2.Sum(nil)
	adler32buffer := new(bytes.Buffer)
	adler32buffer.Write(md5str)
	adler32buffer.Write(src)
	//header.Write(proto.EncodeVarint(uint64(LOGIN_RSA_VER)))
	adler32 := crc32.ChecksumIEEE(adler32buffer.Bytes())
	header.Write(proto.EncodeVarint(uint64(adler32)))
	header.Write([]byte{0xFF})                                  //占坑
	header.Write(proto.EncodeVarint(uint64(RqtCalcData(body)))) //占坑
	header.Write([]byte{0x00})                                  //占坑
	if use_compress {
		lens := (len(header.Bytes()) << 2) + 1
		header.Bytes()[1] = byte(lens)
	} else {
		lens := (len(header.Bytes()) << 2) + 2
		header.Bytes()[1] = byte(lens)
	}
	header.Write(body)
	return header.Bytes()
}

func UnpackBusinessPacket(src []byte, key []byte, uin uint32, cookie *[]byte) []byte {
	if len(src) < 0x20 { //这里需要处理断线重连
		return nil
	} else {
		var nCur int64
		var bfbit byte
		srcreader := bytes.NewReader(src)
		binary.Read(srcreader, binary.BigEndian, &bfbit)
		if bfbit == byte(0xbf) {
			nCur += 1
		}
		nLenHeader := src[nCur] >> 2
		bUseCompressed := src[nCur] & 0x3
		nCur += 1
		nLenCookie := src[nCur] & 0xf
		nCur += 1
		nCur += 4
		srcreader.Seek(nCur, io.SeekStart)
		binary.Read(srcreader, binary.BigEndian, &uin)
		nCur += 4
		cookie_temp := src[nCur : nCur+int64(nLenCookie)]
		*cookie = cookie_temp
		nCur += int64(nLenCookie)
		cgidata := src[nCur:]
		_, nSize := proto.DecodeVarint(cgidata)
		nCur += int64(nSize)
		LenProtobufData := src[nCur:]
		_, nLenProtobuf := proto.DecodeVarint(LenProtobufData)
		nCur += int64(nLenProtobuf)
		body := src[nLenHeader:]
		if bUseCompressed == 1 {
			protobufData := DecompressAndAesDecrypt(body, key)
			return protobufData
		} else {
			protobufData := AesDecrypt(body, key)
			return protobufData
		}

	}
	return nil
}

func UnpackBusinessPacketWithAesGcm(src []byte, uin uint32, cookie *[]byte, Serversessionkey []byte) []byte {
	if len(src) < 0x20 { //这里需要处理断线重连
		return nil
	} else {
		var nCur int64
		var bfbit byte
		srcreader := bytes.NewReader(src)
		binary.Read(srcreader, binary.BigEndian, &bfbit)
		if bfbit == byte(0xbf) {
			nCur += 1
		}
		nLenHeader := src[nCur] >> 2
		nCur += 1
		nLenCookie := src[nCur] & 0xf
		nCur += 1
		nCur += 4
		srcreader.Seek(nCur, io.SeekStart)
		binary.Read(srcreader, binary.BigEndian, &uin)
		nCur += 4
		cookie_temp := src[nCur : nCur+int64(nLenCookie)]
		*cookie = cookie_temp
		nCur += int64(nLenCookie)
		cgidata := src[nCur:]
		_, nSize := proto.DecodeVarint(cgidata)
		nCur += int64(nSize)
		LenProtobufData := src[nCur:]
		_, nLenProtobuf := proto.DecodeVarint(LenProtobufData)
		nCur += int64(nLenProtobuf)
		body := src[nLenHeader:]
		protobufdata := AesGcmDecryptWithUncompress(Serversessionkey, body, nil)
		return protobufdata

	}
	return nil
}

func DecompressAndAesDecrypt(body []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	if len(body)%aes.BlockSize != 0 {
		log.Error(fmt.Sprintf("crypto/cipher: data is not a multiple of the block size，[BodyLength：%v] [AesLength：%v]", len(body), aes.BlockSize))
		return nil
	}

	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(body))
	blockMode.CryptBlocks(origData, body)
	origData = PKCS5UnPadding(origData)
	origData = DoZlibUnCompress(origData)
	return origData
}

/*func Compress_rsa(data []byte) []byte {
	strOut := new(bytes.Buffer)
	var publicKey rsa.PublicKey
	s, _ := hex.DecodeString(rsapublicKey)
	publicKey.N = new(big.Int).SetBytes(s)
	rsaLen := len(rsapublicKey) / 8
	if len(data) > (rsaLen - 12) {
		blockCnt := 1
		if ((len(data) / (rsaLen - 12)) + (len(data) % (rsaLen - 12))) == 0 {
			blockCnt = 0
		}

		for i := 0; i < blockCnt; i++ {
			blockSize := rsaLen - 12
			if i == blockCnt-1 {
				blockSize = len(data) - i*blockSize
			}
			temp := data[(i * (rsaLen - 12)):(i*(rsaLen-12) + blockSize)]
			encrypted, _ := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, temp)
			strOut.Write(encrypted)
		}
		return strOut.Bytes()
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, data)
	if err != nil {
		return []byte{}
	}
	return encrypted
}

//RSAEncrypt Rsa加密
func RSAEncrypt(data []byte) []byte {
	m := rsapublicKey

	M := new(big.Int)
	M.SetString(m, 16)

	pub := rsa.PublicKey{}
	pub.E = 65537
	pub.N = M

	out, _ := rsa.EncryptPKCS1v15(rand.Reader, &pub, data)
	return out
}*/

//PKCS7Padding 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS7UnPadding 去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func Get62Key(Key string) string {
	if len(Key) < 344 {
		return device.MD5ToLower(RandSeq(15))
	}
	K, _ := hex.DecodeString(Key[134:198])
	return string(K)
}

func Get62Data(DevicelId string) string {
	return "62706c6973743030d4010203040506090a582476657273696f6e58246f626a65637473592461726368697665725424746f7012000186a0a2070855246e756c6c5f1020" + hex.EncodeToString([]byte(DevicelId)) + "5f100f4e534b657965644172636869766572d10b0c54726f6f74800108111a232d32373a406375787d0000000000000101000000000000000d0000000000000000000000000000007f"
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Get16Data() string {
	value, _ := RandomHex(8)
	return "A" + value
}

func GetFileMD5Hash(Data []byte) string {
	hash := md5.New()
	hash.Write(Data)
	retVal := hash.Sum(nil)
	return hex.EncodeToString(retVal)
}
