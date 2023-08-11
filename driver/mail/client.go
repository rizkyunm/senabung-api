package mail

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"
)

var (
	mail, mailPassword, hostAddress, hostPort string
	client                                    *smtp.Client
)

func authMail() *smtp.Client {
	mail = os.Getenv("MAIL_SENDER")
	mailPassword = os.Getenv("MAIL_PASSWORD")
	hostAddress = os.Getenv("MAIL_HOST")
	hostPort = os.Getenv("MAIL_PORT")

	auth := smtp.PlainAuth("", mail, mailPassword, hostAddress)

	tlsConfigurations := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         hostAddress,
	}

	conn, err := tls.Dial("tcp", hostAddress+":"+hostPort, tlsConfigurations)
	if err != nil {
		log.Panic(err)
	}

	newClient, err := smtp.NewClient(conn, hostAddress)
	if err != nil {
		log.Panic(err)
	}

	if err = newClient.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = newClient.Mail(mail); err != nil {
		log.Panic(err)
	}

	return newClient
}

func GetMail() *smtp.Client {
	if client != nil {
		return client
	}

	return authMail()
}
