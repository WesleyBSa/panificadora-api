package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"panificadora-api/utils"

	"github.com/gin-gonic/gin"
)

func SendCode(c *gin.Context) {
	phoneNumber := c.PostForm("phone_number")

	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	err := utils.SendVerificationCode(phoneNumber, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification code"})
		return
	}

	// Simulando o armazenamento do código - Aqui testarei o cache e o armazenamento no banco
	fmt.Printf("Verification code for %s is %s\n", phoneNumber, code)

	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent"})
}
