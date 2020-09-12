package mailer

import (
	"fmt"
	"mailer/pkg/config"
	"mailer/pkg/mailutils/outlook"

	"net/smtp"
)

func SendMail(conf *config.Conf, to []string, title, content string) error {
	auth := outlook.OutLookAuth(conf.UserName, conf.Password)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", conf.Server, conf.ServerPort),
		auth, conf.UserName, to,
		[]byte(outlook.OutLookBody(conf.UserName, to, title, content)))
	if err != nil {
		return err
	}
	return nil
}
