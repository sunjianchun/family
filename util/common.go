package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

//Dealerr 统一处理错误函数
func Dealerr(err error, flag string) {
	if err != nil {
		fmt.Println(err)
		if flag == Return {
			return
		} else if flag == Exit {
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

//CryptoAesEncode加密函数
func CryptoAesEncode(origin string) string {
	originBytes := []byte(origin)
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		Dealerr(err, Return)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	var encodeBytes = make([]byte, len(originBytes))
	cfb.XORKeyStream(encodeBytes, originBytes)
	return string(encodeBytes)
}

//CryptoDecDecode加密函数
func CryptoDecDecode(encodeString string) string {
	encodeStringBytes := []byte(encodeString)
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		Dealerr(err, Return)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	var decodeBytes = make([]byte, len(encodeStringBytes))
	cfbdec.XORKeyStream(decodeBytes, encodeStringBytes)
	return string(decodeBytes)
}

func Md5Encode(origin string) string {
	m := md5.New()
	io.WriteString(m, origin)
	firstMd5Password := fmt.Sprintf("%x", m.Sum(nil))
	io.WriteString(m, salt1)
	io.WriteString(m, firstMd5Password)
	io.WriteString(m, salt2)
	result := fmt.Sprintf("%x", m.Sum(nil))
	return result
}
