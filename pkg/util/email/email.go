package email

import (
	conf "gin-mall/conf/sql"
	"gin-mall/consts"
	"gopkg.in/mail.v2"
)

type EmailSender struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpEmailFrom string `json:"smtp_email_from"`
	SmtpPass      string `json:"smtp_pass"`
}

func NewEmailSender() *EmailSender {
	eConfig := conf.Config.Email
	return &EmailSender{
		SmtpHost:      eConfig.SmtpHost,
		SmtpEmailFrom: eConfig.SmtpEmail,
		SmtpPass:      eConfig.SmtpPass,
	}
}

func (s *EmailSender) Send(data, emailTo, subject string) error {
	m := mail.NewMessage()
	m.SetHeader("From", s.SmtpEmailFrom) //SmtpEmailFrom is the email address of the sender
	m.SetHeader("To", emailTo)           // emailTo is the email address of the recipient
	m.SetHeader("Subject", subject)      // subject is the subject of the email
	m.SetBody("text/html", data)         // data is the body of the email
	d := mail.NewDialer(s.SmtpHost, consts.SmtpEmailPort, s.SmtpEmailFrom, s.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil // return nil if the email is sent successfully
}
