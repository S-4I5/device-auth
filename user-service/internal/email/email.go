package email

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"user-service/internal/config"
	"user-service/internal/model/entity"
)

type Sender interface {
	SendEmailVerification(user entity.User, code uuid.UUID)
}

type senderService struct {
	dialer *gomail.Dialer
}

func NewSender(cfg config.MailConfig) *senderService {
	fmt.Println(cfg)
	return &senderService{dialer: gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Password)}
}

func (s *senderService) SendEmailVerification(user entity.User, code uuid.UUID) {

	go func(user entity.User, code uuid.UUID, dialer *gomail.Dialer) {
		m := gomail.NewMessage()
		m.SetHeader("From", "rt.auth@mail.ru")
		m.SetHeader("To", user.Email)
		m.SetHeader("Subject", "Hello!")
		m.SetBody("text/html", code.String())

		err := dialer.DialAndSend(m)
		if err != nil {
			fmt.Println("cannot send email: ", err.Error())
		}

	}(user, code, s.dialer)

}
