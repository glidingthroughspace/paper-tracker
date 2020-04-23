package utils

import (
	"errors"
	"fmt"
	"paper-tracker/config"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

// SendMail sends out an email to the configured email recipients
func SendMail(notificationTitle, notificationText string) (err error) {
	smtpHost := config.GetString(config.KeyMailHost)
	if smtpHost == "" {
		log.Error("SMTP Host not configured")
		return errors.New("SMTP Host not configured")
	}

	smtpPort := config.GetInt(config.KeyMailPort)
	smtpUsername := config.GetString(config.KeyMailUsername)
	smtpPassword := config.GetString(config.KeyMailPassword)
	smtpSender := config.GetString(config.KeyMailSender)
	if smtpSender == "" {
		smtpSender = smtpUsername
	}
	smtpRecipients := config.GetStringSlice(config.KeyMailRecipients)

	msg := fmt.Sprintf("<h1>Paper-Tracker Notification</h1><p>%s</p>", notificationText)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpSender)
	m.SetHeader("To", smtpRecipients...)
	m.SetHeader("Subject", notificationTitle)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	err = d.DialAndSend(m)
	if err != nil {
		log.WithError(err).Error("Failed to send notification mail")
		return
	}

	return
}
