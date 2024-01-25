int nativewcswbaes(char* library,char *srcbuffer, unsigned int srcbuffersize,
                     char *iv, unsigned int ivsize,
                     char *saedattablekey, unsigned int saedattablekeylen,
                     char *saedattablevalue, unsigned int saedattablevaluelen,
                     char *saetablefinal, unsigned int saetablefinallen,
                     unsigned int encodetype,
                     char *outbuffer, unsigned int *outbufferlen);

                     int makeKeyHash(int key);
