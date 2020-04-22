package utils

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// SendMail sends out an email to the configured email recipients
func SendMail(notificationTitle, notificationText string) (err error) {
	smtpHost := viper.GetString("mail.host")
	if smtpHost == "" {
		log.Error("SMTP Host not configured")
		return errors.New("SMTP Host not configured")
	}

	smtpPort := viper.GetInt("mail.port")
	smtpUsername := viper.GetString("mail.username")
	smtpPassword := viper.GetString("mail.password")
	smtpSender := viper.GetString("mail.sender")
	if smtpSender == "" {
		smtpSender = smtpUsername
	}
	smtpRecipients := viper.GetStringSlice("mail.recipients")

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
