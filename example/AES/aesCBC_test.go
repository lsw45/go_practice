package AES

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

const (
	sKey        = "dde4b1f8a9e6b814"
	ivParameter = "dde4b1f8a9e6b814"
	content     = "中国最好，中国最棒，ye"
)

func TestBase64(t *testing.T) {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg)) //base64.StdEncoding.EncodeToString():no padding
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}

func TestAESCBC(t *testing.T) {
	key := []byte(sKey)
	iv := []byte(ivParameter)
	text := []byte(content)
	fmt.Printf("key length:%v\n", len(key))
	fmt.Printf("iv length:%v\n", len(iv))
	fmt.Printf("text length:%v\n", len(text))

	pass := []byte(encryptAESCBC(text, key, iv))

	dst := make([]byte, 2*len(pass))
	Encode(dst, pass)

	fmt.Println(string(dst))
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

//AES加密，CBC,Nopadding:AES的NoPadding模式加密的key和data的byte字节数必须为16的倍数
func encryptAESCBC(src, key, iv []byte) string {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	//进行填充
	fmt.Println(blockSize)
	src = ZeroPadding(src, blockSize)
	crypted := make([]byte, len(src))
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, src)
	return string(crypted)
}

const hextable = "0123456789abcdef"

//the hexadecimal(十六进制的) encoding of src.
func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}

	return len(src) * 2
}

func TestAESCBC1(t *testing.T) {
	key := []byte(sKey)
	text := []byte(content)

	pass, _ := AesEncrypt(text, key)
	fmt.Println(string(pass))
}

func AesEncrypt(origData []byte, key []byte) ([]byte, error) { //AES 加密
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = ZeroPadding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
