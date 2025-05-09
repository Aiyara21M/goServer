package secure

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const secretKey = "mysecret"

func VerifyHMACMiddleware(c *fiber.Ctx) error {
	rawBody := c.Body()
	signature := c.Get("X-Signature")

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(rawBody)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	fmt.Println("Raw body:", string(rawBody))
	fmt.Println("Expected signature:", expectedMAC)
	fmt.Println("Received signature:", signature)

	if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid HMAC signature",
		})
	}

	return c.Next()
}

func HMACSign(message string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
