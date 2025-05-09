package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	aesKey []byte
	aesIV  []byte
	once   sync.Once
)

func initAES() {
	key := os.Getenv("AES_KEY")
	iv := os.Getenv("AES_IV")

	if len(key) != 16 || len(iv) != 16 {

		key = "1234567890123456"
		iv = "6543210987654321"
	}

	aesKey = []byte(key)
	aesIV = []byte(iv)
}

func DecryptAES(c *fiber.Ctx) (string, error) {
	var req struct {
		Data string `json:"data"`
	}
	if err := c.BodyParser(&req); err != nil {
		return "", fmt.Errorf("invalid encrypted body")
	}

	plaintext, err := AESDecrypt(req.Data)
	if err != nil {
		return "", fmt.Errorf("AES decrypt failed")
	}
	// fmt.Println("Decrypted plaintext:", plaintext)
	return plaintext, nil
}

func AESDecrypt(base64Cipher string) (string, error) {
	once.Do(initAES)

	cipherText, err := base64.StdEncoding.DecodeString(base64Cipher)
	if err != nil {
		return "", err
	}
	fmt.Println("Ciphertext length:", len(cipherText))

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	if len(cipherText)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	fmt.Println("Ciphertext length:", len(cipherText))
	fmt.Println("cipherText :", (cipherText))
	mode := cipher.NewCBCDecrypter(block, aesIV)
	plain := make([]byte, len(cipherText))
	mode.CryptBlocks(plain, cipherText)

	fmt.Printf("Raw bytes: %x\n", plain)

	padLen := int(plain[len(plain)-1])
	fmt.Println("padLen:", padLen)
	if padLen > 0 && padLen <= aes.BlockSize {
		fmt.Println("len(plain)-padLen:", len(plain), "-", padLen)
		fmt.Printf("bytes: %x\n", plain[:16])

		plain = plain[:len(plain)-padLen]
	} else {
		return "", errors.New("invalid padding")
	}

	fmt.Printf("Clean bytes: %x\n", plain)

	return string(plain), nil
}

func AESEncrypt(plainText string) (string, error) {
	once.Do(initAES)

	padLen := aes.BlockSize - (len(plainText) % aes.BlockSize)
	padding := make([]byte, padLen)
	for i := range padding {
		padding[i] = byte(padLen)
	}
	paddedText := append([]byte(plainText), padding...)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, aesIV)
	mode.CryptBlocks(ciphertext, paddedText)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
