package utils

import "fmt"

func SendVerificationCode(phoneNumber string, code string) error {
	// Aqui integrarei algum serviço real como Twilio, Vonage, etc. (Vou pesquisar mais sobre)
	fmt.Printf("Sending code %s to phone number %s\n", code, phoneNumber)
	return nil
}
