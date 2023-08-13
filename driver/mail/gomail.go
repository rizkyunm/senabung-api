package mail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

var (
	mail, mailPassword, hostAddress, hostPort, senderName string
	dialer                                                *gomail.Dialer
)

func newDialer() *gomail.Dialer {
	mail = os.Getenv("MAIL_SENDER")
	mailPassword = os.Getenv("MAIL_PASSWORD")
	hostAddress = os.Getenv("MAIL_HOST")
	hostPort = os.Getenv("MAIL_PORT")
	senderName = os.Getenv("MAIL_SENDER_NAME")

	port, _ := strconv.Atoi(hostPort)
	dialer = gomail.NewDialer(hostAddress, port, mail, mailPassword)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         hostAddress,
	}

	return dialer
}

func GetDialer() *gomail.Dialer {
	if dialer == nil {
		dialer = newDialer()
	}

	return dialer
}

func GetSenderName() string {
	return senderName
}

func GetHostAddress() string {
	return hostAddress + ":" + hostPort
}

func GetMail() string {
	return mail
}
