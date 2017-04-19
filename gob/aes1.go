package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// 数据操作的偏移量
var IV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	// 要加密的内容
	content := []byte("my name is nljb")
	// KEY, 必须是16,24,32位的[]byte
	// 分别对应AES_128,AES_192,AES_256
	key := "B8XKCA7IVW6WB7GX76V771RN8LJCY2H0"
	// 通过密钥生成一个新的密码块
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// 加密模式 (ECB、CBC、CFB、OFB)
	// IV是initialization vector的意思
	// 就是加密的初始话矢量，初始化加密函数的变量
	// 也就是加密动作中的 数据操作的偏移量
	cfb := cipher.NewCFBEncrypter(c, IV)
	// 存储密码, 必须与块体的长度相同
	fmt.Println(len(content))
	ciphertext := make([]byte, len(content))
	// 流化, 必须与块体的长度相同
	cfb.XORKeyStream(ciphertext, content)
	// 输出
	fmt.Println(string(content), ciphertext)
	// ------------------------------------- //
	// 解密模式 (ECB、CBC、CFB、OFB)
	// 也就是说，解密的时候也需要加密时的密钥与偏移量
	cfbdec := cipher.NewCFBDecrypter(c, IV)
	// 存储数据, 必须与块体的长度相同
	plaintextCopy := make([]byte, len(content))
	// 流化, 必须与块体的长度相同
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	// 输出
	fmt.Println(ciphertext, string(plaintextCopy))
}
