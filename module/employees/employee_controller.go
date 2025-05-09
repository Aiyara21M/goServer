package employees

import (
	"encoding/json"
	"fmt"
	"server/secure"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController interface {
	GetEmployees(c *fiber.Ctx) error
	GetEmployeeByID(c *fiber.Ctx) error
	CreateSampleEmployee(c *fiber.Ctx) error
}

type employeeController struct {
	service EmployeeService
}

func (ctrl *employeeController) GetEmployees(c *fiber.Ctx) error {
	employees, err := ctrl.service.GetAllEmployees("db3")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employees)
}

func (ctrl *employeeController) GetEmployeeByID(c *fiber.Ctx) error {

	plaintext, err := secure.DecryptAES(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("Decrypted plaintext:", plaintext)
	type reqModel struct {
		ID int `json:"ID"`
	}
	var data reqModel
	err = json.Unmarshal([]byte(plaintext), &data)
	if err != nil {
		fmt.Print(err)
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})

	}

	id := uint(data.ID)

	employee, err := ctrl.service.GetEmployeeID("db3", id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ✨ มาเริ่มเข้ารหัสกลับก่อนส่ง

	response := map[string]interface{}{
		"employee": employee,
	}

	// 1. marshal เป็น JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to marshal response"})
	}

	// 2. AES Encrypt
	encryptedData, err := secure.AESEncrypt(string(responseJSON))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to encrypt response"})
	}

	// 3. HMAC Sign
	signature := secure.HMACSign(encryptedData)

	c.Set("X-Signature", signature)

	// 4. ส่งกลับ
	return c.JSON(fiber.Map{
		"data": encryptedData,
	})
}
