package comm

import (
	//"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"wechatwebapi/Mmtls"
	"wechatwebapi/models"
)

var RedisClient *redis.Client

type LoginData struct {
	Uin                        uint32
	Wxid                       string
	Pwd                        string
	Uuid                       string
	Aeskey                     []byte
	NotifyKey                  []byte
	Deviceid_str               string
	Deviceid_byte              []byte
	DeviceType                 string
	ClientVersion              int
	DeviceName                 string
	NickName                   string
	Alais                      string
	Mobile                     string
	Mmtlsip                    string
	MmtlsHost                  string
	Sessionkey                 []byte
	Sessionkey_2               []byte
	Autoauthkey                []byte
	Autoauthkeylen             int32
	Clientsessionkey           []byte
	Serversessionkey           []byte
	HybridEcdhPrivkey          []byte
	HybridEcdhPubkey           []byte
	HybridEcdhInitServerPubKey []byte
	Loginecdhkey               []byte
	Cooike                     []byte
	AuthTicket                 string
	Proxy                      models.ProxyInfo
	MmtlsKey                   *Mmtls.MmtlsClient
}

func RedisInitialize() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "",           // redis密码，没有则留空
		DB:       0,            // 默认数据库，默认是0
	})

	return RedisClient
}

func CreateLoginData(data LoginData, key string, Expiration int64) error {
	var ExpTime time.Duration
	if key == "" {
		key = data.Uuid
	}

	if Expiration > 0 {
		ExpTime = time.Second * time.Duration(Expiration)
	} else {
		ExpTime = 0
	}

	JsonData, _ := json.Marshal(&data)
	//ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	err := RedisClient.Set( key, string(JsonData), ExpTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetKeyJsonData(Key string) (ret string, err error) {
	//ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	val, _ := RedisClient.Get(Key).Result()
	if val == "" {
		return ret, errors.New(fmt.Sprintf("[Key:%v]数据不存在", Key))
	}
	return val, nil
}

func GetLoginata(key string) (*LoginData, error) {
	P, err := GetKeyJsonData(key)
	if err != nil {
		return &LoginData{}, err
	}
	D := &LoginData{}
	err = json.Unmarshal([]byte(P), D)
	if err != nil {
		return &LoginData{}, err
	}

	return D, nil
}
