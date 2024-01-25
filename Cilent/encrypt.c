#ifdef _WIN32
#include <windows.h>

#ifndef false
#define false   (0)
#endif
#ifndef true
#define true   (1)
#endif

#else
#include <dlfcn.h>
#ifndef false
#define false   (0)
#endif
#ifndef true
#define true   (1)
#endif

#endif
#include "encrypt.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h> /* memset */
#include <unistd.h> /* close */


#define nullptr NULL
typedef short int int16_t;
void getsectable(char sectable[64], char *encryptrecordbuf, char *saedattablekey) {
       int localindex3; // r9
       int localindex; // r12
       char* outbuffer; // r8
       char* ptrsattable_1; // lr
       int localindex2; // r5
       char* ptrsattables; // r6
       char tempbyte; // r3
       localindex3 = 0;
       do
       {
           localindex = 0;
           outbuffer = sectable;
           ptrsattable_1 = saedattablekey;
           do
           {
               localindex2 = 0;
               ptrsattables = ptrsattable_1;
               do
               {
                   unsigned int  temp = 4 * *(unsigned char*)(encryptrecordbuf + 4 * localindex3 + localindex);
                   tempbyte = *(unsigned char *)(ptrsattables +temp);
                   ptrsattables++;
                   *(unsigned char *)(outbuffer + localindex2) = tempbyte;
                   localindex2 += 16;
               }
               while ( localindex2 != 0x40 );
               ++localindex;
               outbuffer += 4;
               ptrsattable_1 += 0x400;
           }
           while ( localindex != 4 );
           ++localindex3;
           ++sectable;
           saedattablekey += 0x1000;
       }
       while ( localindex3 != 4 );
   }


   void circshift(char * srcbuffer,int pos){
       char v2; // r1
       char v3; // r3
       char v4; // r2
       char v5; // r3
       char v6; // r1
       switch ( pos )
       {
           case 3:
               v4 = *((char *)srcbuffer + 3);
               v5 = *srcbuffer;
               *((char *)srcbuffer + 3) = *((char *)srcbuffer + 2);
               *(char *)((char *)srcbuffer + 2) = *((char *)srcbuffer + 1);
               *(char *)((char *)srcbuffer + 1) = v5;
               *(char *)srcbuffer = v4;
               break;
           case 2:
               v6 = *srcbuffer;
               *srcbuffer = srcbuffer[2];
               srcbuffer[2] = v6;

               v6 = *(char *)((char *)srcbuffer + 1);
               srcbuffer[1] = srcbuffer[3];
               *(char *)((char *)srcbuffer + 3) = v6;

               break;
           case 1:
               v2 = *(char *)srcbuffer;
               v3 = *(char *)((char *)srcbuffer + 1);
               *(char *)((char *)srcbuffer + 1) = *((char *)srcbuffer + 2);
               *((char *)srcbuffer + 2) = *((char *)srcbuffer + 3);
               *srcbuffer = v3;
               *((char *)srcbuffer + 3) = v2;
               break;
       }
   }
   void shiftrows(char *encryptrecordbuf) {
       circshift((encryptrecordbuf + 4), 1);
       circshift((encryptrecordbuf + 8), 2);
       circshift((encryptrecordbuf + 12), 3);
   }


   void getsecvalue(char *encryptrecordbuf, char sectable[64], char *saedatatable) {
       char *ptrtable64_offset2; // r3
       int outlocalindex; // r12
       int localindex; // r8
       char *ptrtable64_offset2_1; // lr
       char *saetable; // r11
       int innerindex; // r6
       unsigned char tempbytevalue; // r5
       char *outbufferptr; // r3
       char *out16buf_1; // r10
       char *ptrtable64_offset2_3; // r3
       unsigned int v13; // r4
       unsigned int v14; // r4
       unsigned char v15; // r2
       unsigned int v16; // r5
       char *ptrtable64_offset2_2; // [sp+0h] [bp-20h]
       char *saetableptr; // [sp+4h] [bp-1Ch]

       ptrtable64_offset2 = sectable + 2;
       outlocalindex = 0;
       do
       {
           localindex = 0;
           ptrtable64_offset2_1 = ptrtable64_offset2;
           saetable = saedatatable;
           ptrtable64_offset2_2 = ptrtable64_offset2;
           saetableptr = saedatatable;
           do
           {
               innerindex = 0;
               tempbytevalue = sectable[16 * outlocalindex + 3 + 4 * localindex];
               outbufferptr = &encryptrecordbuf[4 * outlocalindex];
               out16buf_1 = &outbufferptr[localindex];
               outbufferptr[localindex] = tempbytevalue;
               ptrtable64_offset2_3 = ptrtable64_offset2_1;
               do
               {
                   v13 = (unsigned char)(*ptrtable64_offset2_3 & 0xF0) | ((tempbytevalue & 0xF0u) >> 4);
                   if ( (v13 & 0x80u) != 0 ) {             // 大于128
                       v14 = (unsigned int) (unsigned char) saetable[(v13 & 0x7F) + 0x200 + innerindex] >> 4;
                   }
                   else {
                       v14 = saetable[v13 + 0x200 + innerindex] & 0xF;
                   }
                   v15 = tempbytevalue & 0xF | 16 * *ptrtable64_offset2_3;

                   if ( (v15 & 0x80u) != 0 ) {               // 大于128
                       v16 = (unsigned int) (unsigned char) saetable[(v15 & 0x7F) + 0x280 + innerindex] >> 4;
                   }
                   else {
                       v16 = saetable[v15 + 0x280 + innerindex] & 0xF;
                   }
                   tempbytevalue = v16 | 16 * v14;
                   innerindex -= 0x100;
                   --ptrtable64_offset2_3;
                   *out16buf_1 = tempbytevalue;
               }
               while ( innerindex != 0xFFFFFD00 );
               ++localindex;
               ptrtable64_offset2_1 += 4;
               saetable += 0x300;
           }
           while ( localindex != 4 );
           ++outlocalindex;
           ptrtable64_offset2 = ptrtable64_offset2_2 + 0x10;
           saedatatable = saetableptr + 0xC00;
       }
       while ( outlocalindex != 4 );
   }
   void getsecvaluefinnal(char *inbuffer, char *a3) {
       int v3; // r12
       char *v4; // lr
       char* v5; // r3
       int v6; // r0
       char* v7; // r4
       char v8; // r5
       int v9; // r0
       char v14[16] = {0}; // [sp+0h] [bp-1Ch]
       int v15; // [sp+10h] [bp-Ch]

       v3 = 0;
       v4 = v14;
       v5 = inbuffer;
       do
       {
           v6 = 0;
           v7 = a3;
           do
           {
               v8 = *(char *)(v7 + *(unsigned char *)(v5 + v6));
               v7 += 0x100;
               *((char *)v4 + v6++) = v8;
           }
           while ( v6 != 4 );
           ++v3;
           v4+=4;
           a3 += 1024;
           v5 += 4;
       }
       while ( v3 != 4 );
       memcpy(inbuffer,v14,0x10);
   }
int nativewcswbaes(char* library,char* srcbuffer, unsigned int srcbuffersize,
                                              char* iv, unsigned int ivsize,
                                              char* saedattablekey, unsigned int saedattablekeylen,
                                              char* saedattablevalue, unsigned int saedattablevaluelen,
                                              char * saetablefinal,unsigned int saetablefinallen,
                                              unsigned int encodetype,
                                              char* psaeoutbuffer, unsigned int * outbufferlen) {
   //先判断原始输入
       //sae.dat解析完成的两个字段是否都有
       unsigned int  encryptdatalen = 0;
       char * encryptdatabuf = nullptr;
       char * outencryptdatabuf = nullptr;
       char * encryptrecordbuf = nullptr;
       char sectable[64] = {0};
       char * destbufferptr = nullptr;
       char * destbuffer = nullptr;
       if (srcbuffer && saedattablekey && saedattablevalue)
       {
           if (encodetype & 0x3060 || encodetype & 4095) //如果加密方式是3
           {
               encryptdatalen = 16 - (srcbuffersize & 0xF) + srcbuffersize;
               encryptdatabuf = (char*)malloc(encryptdatalen);
               //先申请一段加密完成以后的大小内存
               memcpy(encryptdatabuf, srcbuffer, srcbuffersize);
               //如果不是0x10的整数倍就后面补齐
               if ( 16 != (srcbuffersize & 0xF) ){
                   memset((char *)encryptdatabuf + srcbuffersize, 16 - (srcbuffersize & 0xF), 16 - (srcbuffersize & 0xF));
               }
               //申请一块同样大小的内存作为轮询数据
               outencryptdatabuf = (char *)malloc(encryptdatalen);
               memset(outencryptdatabuf, 0, encryptdatalen);
               //申请一个0x10的内存做中间加密换算用
               encryptrecordbuf = (char *)malloc(0x10);
               memset(encryptrecordbuf, 0, 0x10);

               if (encryptdatalen)
               {
                   char* xorivbuffer = iv; //首先与IV进行xor
                   unsigned int uncryptdatalen = encryptdatalen; //还有多少没有加密
                   char * intdatadstptr_2_offse_1 = (char*)(outencryptdatabuf + 1);
                   char * bytedatadstptr_2_offse_1 = nullptr;
                   char * bytedatadstptr_2_offse_1_1 = nullptr;
                   char * outencryptdatabuf_temp = (char*)outencryptdatabuf;
                   while (true ) {
                       int localindex = 0;
                       int islarge16 = 0;
                       do {
                           bytedatadstptr_2_offse_1 = (char *)intdatadstptr_2_offse_1;
                           bytedatadstptr_2_offse_1_1 = bytedatadstptr_2_offse_1;
                           *(outencryptdatabuf_temp + localindex++) = encryptdatabuf[localindex] ^ *(unsigned int *)(xorivbuffer + localindex);
                           islarge16 = localindex >= 0xF;
                           if ( localindex <= 0xF )
                           {
                               intdatadstptr_2_offse_1 = (char*)(bytedatadstptr_2_offse_1 + 1);
                               islarge16 = localindex >= uncryptdatalen;
                           }
                       } while (!islarge16);
                       if ( localindex <= 0xF )
                       {
                           do
                           {
                               *bytedatadstptr_2_offse_1++ = *(char *)(xorivbuffer + localindex++);
                           }
                           while ( localindex != 16 );
                       }
                       int temprecordindex = 0 ;
                       char * outencryptdatabuf_temp_record = outencryptdatabuf_temp;
                       char * encryptrecordbuf_temp = encryptrecordbuf;
                       do
                       {
                           int localindex2 = 0;
                           do
                           {
                               encryptrecordbuf_temp[localindex2] = *(char *)(outencryptdatabuf_temp_record + 4 * localindex2);
                               ++localindex2;
                           }
                           while ( localindex2 != 4 );
                           ++temprecordindex;
                           encryptrecordbuf_temp += 4;
                           ++outencryptdatabuf_temp_record;
                       }
                       while ( temprecordindex != 4 );
                       destbufferptr = outencryptdatabuf_temp;
                       if ( ((unsigned char)encodetype & 5) == 5 ) {
                           //这里需要申请一块内存先

                           //mzd_t * ptrAmzd = mzd_init(0x80,0x80);
                           //sub_EE952((encryptrecordbuf, outencryptdatabuf_temp, , v1058);
                       }
                       if ( (unsigned char)encodetype & 8 ) {
                           //sub_EE68A(&v1116, &encryptrecordbuf, v1057);
                       }
                       destbufferptr = outencryptdatabuf_temp;
                       if ( (unsigned char)encodetype & 0x10 )
                       {
                           //sub_EE796(&encryptrecordbuf, &v1116, v1064);
                       }
                       shiftrows(encryptrecordbuf);      //行移位

                       unsigned int indexstep = 0xFFFDC000;
                       //printf("saedattablekey : %p \r\n",saedattablekey);
                       //printf("saedattablevalue : %p \r\n",saedattablevalue);
                       char * psaedattablekey = saedattablekey;
                       char * psaedattablevalue = saedattablevalue;
                       do
                       {
                           if ( (unsigned char)encodetype & 0x20 )
                               getsectable(sectable,encryptrecordbuf,psaedattablekey);// 密钥扩展
                           if ( (unsigned char)encodetype & 0x40 )
                               getsecvalue(encryptrecordbuf, sectable, psaedattablevalue);
                           if ( (unsigned char)encodetype & 0x80 )
                           {
                               //sub_EE632((int)&v1115, (int)&encryptrecordbuf, v1104 + v692 + 0x3F000);
                           }
                           if ( (unsigned char)encodetype & 0x100 )
                           {
                               //sub_EE556((int)&encryptrecordbuf, (int)&v1115, v691);
                           }
                           shiftrows(encryptrecordbuf);      //行移位
                           psaedattablevalue += 0x3000;
                           indexstep += 0x4000;
                           psaedattablekey += 0x4000;

                       }while ( indexstep );
                       if ( (int16_t)encodetype & 0x1000 )
                           getsecvaluefinnal(encryptrecordbuf,saetablefinal);//最后一次加密
                       if ( (int16_t)encodetype & 0x200 )
                       {
                           //sub_EE710(&v1116, &encryptrecordbuf, v1063);
                       }
                       if ( (int16_t)encodetype & 0x400 ) {
                           //sub_EE874(&encryptrecordbuf, &v1116, v1061);
                       }
                       if ( (encodetype & 0x802) == 2050 )
                       {
                           destbuffer = destbufferptr;
                           //sub_EE9EA(destbufferptr, (int)&encryptrecordbuf, *(( *)v1070 + 9), v1060);
                       }
                       else
                       {
                           destbuffer = destbufferptr;
                           if ( (int16_t)encodetype & 0x1000 ) {
                               char *unkownresetdata_2 = encryptrecordbuf;
                               int v694 = 0;
                               char *destbuffer_1 = destbufferptr;
                               do {
                                   int index = 0;
                                   do {
                                       *(char *) (destbuffer_1 + 4 * index) = unkownresetdata_2[index];
                                       ++index;
                                   } while (index != 4);
                                   ++v694;
                                   unkownresetdata_2 += 4;
                                   ++destbuffer_1;
                               } while (v694 != 4);
                           }
                       }
                       if ( uncryptdatalen < 0x11 )
                           break;
                       uncryptdatalen -= 16;
                       xorivbuffer = destbuffer;
                       intdatadstptr_2_offse_1 = bytedatadstptr_2_offse_1_1 + 16;
                       encryptdatabuf += 16;
                       outencryptdatabuf_temp += 16;
                   }
                   //printf("psaeoutbuffer %p %p",psaeoutbuffer,outbufferlen);
                   *outbufferlen = encryptdatalen;
                   memset(psaeoutbuffer,0,encryptdatalen);
                   memcpy(psaeoutbuffer,outencryptdatabuf,encryptdatalen);
                //    if(encryptdatabuf != NULL){
                //       free(encryptdatabuf);
                //    }
                //    if(outencryptdatabuf != NULL){
                //       free(outencryptdatabuf);
                //    }
               }
           }
       } else
       {
           return false;
       }
       return true;

   }

   int makeKeyHash(int key)
   {
   	// TODO: 在此处添加实现代码.
   	int a=0, b=0;
   	int a_result =0, b_result = 0;
	int i;
   	for ( i=0; i<16; i++)
   	{
   		a = 1 << (2 * i);
   		b = 1 << (2 * i + 1);

   		a &= key;
   		b &= key;

   		a = a << (2 * i);
   		b = b << (2 * i + 1);

   		a_result |= a;
   		b_result |= b;
   	}
   	return a_result | b_result;
   }