package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	associatedData := "mall_transaction"

	nonce := "ozOhsln77kRB"                   // 普通字符串，16进制：37b8e8a308c354048d245f6d
	key := "91320402MA1UUEJ25Etnwy0288878888" // The AES key, either 16 or 32 bytes to select AES-128 or AES-256.
	plainText := "172.10.99.88"

	//nonce = "ozOhsln77kRB"
	//key = "91320402MA1UUEJ25Etnwy0288878888"
	//cipherText := "Oq1f8IDGqmwh65TYOcnbOa+YM+zoIm/PZojsLLJKPOyXveuQtsS54Nh/RzU/FkKruR/C9qG8NUurry5HOctaf2p+KzBZNqGfdaT3FHDeIoyfDXCkaleZMApKLZ3qmHtNuoDtNUYhJXOIjVat8rXHG0++XwF5BVfPph4ClWujJheg62HSkiipknmt9Q3Mt2+ZwJx95LsJfIapnkhWo0qjj7YF6OS7vcWjilrrngpUKMifnGKv0QC4/z/McMo6Z6qAo7FtRa+YuR7who1YleWjuVGsDeotoP+9MmOfvPEDG2iS/RhH2Sn+RXDs1k0gOvp62BeX3vLEEYK/Ck/UMVL5fGDQso7/viN5cLG8Un5Uct9lx3beI/6Hqwv0nk4jb5nJ1XkI"

	//nonce = "nr8xWB87eiJm"
	//cipherText = "2g8f42jmvPyY+um+rPJEmt/RWWAOmgXSjJLY1j5pxpxc/G8Ce5OZL9mpM46mwBNqnhTdUglVlu9pgUePXl/I9f37krN3YEgMYjC0vihNatSO0+3vgBC40e26onxbHdAKHUeQ/J1yX8gxNB0BEDhkUS0s4uDMOIsCBdyZnoBIbqorSvTQ+wjPxxZehzxVVFjKQNleCndBFfKNSyB2Yiz6kmYo3S3qyOhU6K/OYHmZQ4W+kZ49blXziDnWPwJNLZRL5OfDt/kqhkTWj3RYCZzRS0Pqm/PiiSvJU3gmhkD3cmxCfqkG0d5mnPlr4YkZytyWNRj1nfacgIn0rapBMJgcIcHDmPt/KsJ8eE6jFFML68rH7dnftZEVxuvCykpInYmJNZiK"
	cipherText := exampleNewGCM_encrypt(plainText, key, nonce, associatedData)
	newPlain := exampleNewGCM_decrypt(cipherText, key, nonce, associatedData)

	fmt.Println("plain:", plainText)
	fmt.Println("cipher:", cipherText)
	fmt.Println("new plain:", newPlain)
}

func exampleNewGCM_encrypt(src, k, n, a string) string {
	key := []byte(k)
	plaintext := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//nonce, _ := hex.DecodeString(n) n是16进制
	nonce := []byte(n)
	additionalData := []byte(a)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, additionalData)

	return base64.StdEncoding.EncodeToString(ciphertext) //string(ciphertext)是乱码，为了得到可见字符串，对[]byte进行base64加密
}

func exampleNewGCM_decrypt(src, k, n, a string) string {
	key := []byte(k)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext, err := base64.StdEncoding.DecodeString(src) //base64解密
	//nonce,_:= hex.DecodeString(n) n是16进制
	nonce := []byte(n)
	additionalData := []byte(a)

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, additionalData)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
