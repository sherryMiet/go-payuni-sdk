package payuni

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

func Aes256GCMEncrypt(plaintext, key, iv string) string {
	bPlaintext := []byte(plaintext)
	bIV := []byte(iv)
	bKey := []byte(key)

	block, err := aes.NewCipher(bKey)

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		logrus.Error(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, bIV, bPlaintext, nil)
	baseCiphertext := base64.StdEncoding.EncodeToString(ciphertext[:len(bPlaintext)])
	baseTag := base64.StdEncoding.EncodeToString(ciphertext[len(bPlaintext):])

	return hex.EncodeToString([]byte(baseCiphertext + ":::" + baseTag))
}

func Aes256GCMDecrypt(ciphertext, key, iv string) string {
	bCiphertext, _ := hex.DecodeString(ciphertext)
	bIV := []byte(iv)
	bKey := []byte(key)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		logrus.Error(err.Error())
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		logrus.Error(err.Error())
	}
	cipherTextStr := string(bCiphertext)
	plaintextStr := strings.Split(cipherTextStr, ":::")
	Encrypt, err := base64.StdEncoding.DecodeString(plaintextStr[0])
	if err != nil {
		return ""
	}
	Tag, err := base64.StdEncoding.DecodeString(plaintextStr[1])
	if err != nil {
		return ""
	}
	Encrypt = append(Encrypt, Tag...)

	plaintext, err := aesgcm.Open(nil, bIV, Encrypt, nil)
	if err != nil {
		logrus.Error(err.Error())
	}

	unescape, err := url.PathUnescape(string(plaintext))
	if err != nil {
		return ""
	}
	return unescape
}

func SHA256(str string) string {
	sum := sha256.Sum256([]byte(str))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}
