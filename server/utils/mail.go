package utils

import (
	"errors"
	"net/smtp"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SendMail(notificationText string) (err error) {
	smtpHost := viper.GetString("mail.host")
	if smtpHost == "" {
		log.Error("SMTP Host not configured")
		return errors.New("SMTP Host not configured")
	}

	smtpHostPort := smtpHost + ":" + strconv.Itoa(viper.GetInt("mail.port"))
	smtpUsername := viper.GetString("mail.username")
	smtpPassword := viper.GetString("mail.password")
	smtpSender := viper.GetString("mail.sender")
	if smtpSender == "" {
		smtpSender = smtpUsername
	}
	smtpRecipients := viper.GetStringSlice("mail.recipients")

	msg := notificationText

	if smtpUsername != "" {
		smtpAuth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
		err = smtp.SendMail(smtpHostPort, smtpAuth, smtpSender, smtpRecipients, []byte(msg))
	} else {
		err = smtp.SendMail(smtpHostPort, nil, smtpSender, smtpRecipients, []byte(msg))
	}
	if err != nil {
		log.WithError(err).Error("Failed to send notification mail")
		return
	}

	return
}
