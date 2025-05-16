package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOtpMail() {
	Otp := GetSixDigitRandomString()
	from := mail.NewEmail("ChurchAstro", "askijothi@gmail.com")
	subject := "Login OTP"
	to := mail.NewEmail("Jothiraj", "jothiraj.d@readyassist.in")
	plainTextContent := "Your OTP is" + Otp
	htmlContent := "<div>Welcome to ChurchAstro.<br> Your login OTP is <strong>" + Otp + "</strong></div>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println("otp cannot be sent. reason:", err)
	} else {
		fmt.Println("otp sent", response.StatusCode)
	}
}
