package email

import (
	"fmt"
	"github.com/ribaraka/go-srv-example/config"
	"gopkg.in/gomail.v2"
	"log"
)

func SendVerifyMassage(c config.Config, email, token string) error {

	emailBody := fmt.Sprintf(`Thank you for creating an account!
	Please confirm your email address
	by clicking on the link :
	<a href="http://%s//verify?email=%s&token=%s">http://localhost:8081/verify?email=%s&token=%s </a>`,
	c.ServerHost, email, token, email, token)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", c.Mail.Sender)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Email confirm")
	mailer.SetBody("text/html", emailBody)

	dialer := gomail.NewDialer(
		c.Mail.Host,
		c.Mail.Port,
		c.Mail.AuthEmail,
		c.Mail.AuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		err := fmt.Errorf("no user token found %s", err)
		log.Println(err)
		return err
	}
	fmt.Printf( "Mail sent %s!", email)
	return nil
}