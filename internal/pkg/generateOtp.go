package pkg

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
)

func generateOtp() string {
	const digits = 6
	otp := ""

	for i := 0; i < digits; i++ {
		// Generate a random number between 0 and 9
		randomNumber, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			fmt.Println("Error generating otp:", err)
			return ""
		}
		otp += randomNumber.String()
	}
	return otp
}

func GenerateAndSendOtp(email string) (string, error) {
	// Sender's email address and password
	from := "gpurvi.28@gmail.com"
	password := "vbpd oxuh syzi fphi"

	// Recipient's email address
	to := email

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	otp := generateOtp()
	// Message content
	message := []byte("your otp for reseting the password is " + otp)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		return "", err
	}
	return otp, nil
}
