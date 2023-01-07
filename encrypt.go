package payuni

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func Aes256GCMEncrypt(plaintext, key, iv string) string {
	bPlaintext := []byte(plaintext)
	bIV := []byte(iv)
	bKey := []byte(key)

	block, err := aes.NewCipher(bKey)

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, bIV, bPlaintext, nil)
	baseCiphertext := base64.StdEncoding.EncodeToString(ciphertext[:len(bPlaintext)])
	baseTag := base64.StdEncoding.EncodeToString(ciphertext[len(bPlaintext):])

	return hex.EncodeToString([]byte(baseCiphertext + ":::" + baseTag))
}

func Aes256GCMDecrypt(ciphertext, key, iv string) string {
	data, _ := hex.DecodeString(ciphertext)
	bCiphertext := []byte(data)
	bIV := []byte(iv)
	bKey := []byte(key)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(bIV, nil, bCiphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}

func SHA256(str string) string {
	sum := sha256.Sum256([]byte(str))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}
