package mailer

import (
	"fmt"
	"net/smtp"

	"github.com/Vstural/mailer/pkg/config"
	"github.com/Vstural/mailer/pkg/mailutils/outlook"
)

func SendOutlookMail(conf *config.Conf, to []string, title, content string) error {
	auth := outlook.OutLookAuth(conf.UserName, conf.Password)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", conf.Server, conf.ServerPort),
		auth, conf.UserName, to,
		[]byte(outlook.OutLookBody(conf.UserName, to, title, content)))
	if err != nil {
		return err
	}
	return nil
}
