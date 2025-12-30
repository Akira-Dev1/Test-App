package handlers
import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

// GenerateCode создает случайный 6-значный код
func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// sendEmail отправляет письмо с кодом через SMTP
func sendEmail(to, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "max80021151@gmail.com") 
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Ваш код подтверждения")
	m.SetBody("text/plain", fmt.Sprintf("Ваш код: %s", code))

	d := gomail.NewDialer("smtp.gmail.com", 587, "max80021151@gmail.com", "qnxk role tjvn zvmn") 

	return d.DialAndSend(m)
}
