package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func AesDecrypt(cipherText string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return "", errors.New("密文长度不足")
	}

	iv := data[:blockSize]
	data = data[blockSize:]
	out := make([]byte, len(data))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(out, data)

	return string(out), nil // 注意明文是文本时才安全
}
