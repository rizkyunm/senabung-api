package mail

import (
	"bytes"
	"github.com/rizkyunm/senabung-api/driver/mail"
	"github.com/rizkyunm/senabung-api/helper"
	"github.com/rizkyunm/senabung-api/transaction"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"strings"
)

func SendThankYouEmail(transaction transaction.Transaction, transactionNotification transaction.TransactionNotificationInput) (err error) {
	amount := helper.FormatCommas(strings.Split(transactionNotification.GrossAmount, ".")[0])
	templateData := struct {
		Name          string
		TrxNo         string
		TrxDate       string
		Email         string
		PhoneNumber   string
		CampaignName  string
		Amount        string
		URLDonasi     string
		CampaignImage string
	}{
		Name:          transaction.User.Name,
		TrxNo:         transactionNotification.OrderID,
		TrxDate:       transactionNotification.TransactionTime,
		Email:         transaction.User.Email,
		PhoneNumber:   transaction.User.PhoneNumber,
		CampaignName:  transaction.Campaign.Name,
		Amount:        "Rp " + amount,
		URLDonasi:     "https://senabung.me/donasi/" + transaction.Campaign.Slug,
		CampaignImage: transaction.Campaign.CampaignImage,
	}

	r := newRequest([]string{templateData.Email}, "Terimakasih telah berpartisipasi "+templateData.Name, "")
	if err = r.ParseTemplate("/mail/template/thankyou.html", templateData); err != nil {
		return err
	}
	ok, err := r.SendEmail()
	if err != nil || !ok {
		return err
	}

	return nil
}

func SendWelcomeEmail(name, email string) (err error) {
	templateData := struct {
		Nama string
	}{Nama: name}

	r := newRequest([]string{email}, "Selamat datang di Senabung "+name, "")
	if err = r.ParseTemplate("/mail/template/welcome.html", templateData); err != nil {
		return err
	}
	ok, err := r.SendEmail()
	if err != nil || !ok {
		return err
	}

	return nil
}

// Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func newRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.GetMail())
	m.SetHeader("To", r.to...)
	m.SetHeader("Subject", r.subject)
	m.SetBody("text/html", r.body)

	if err := mail.GetDialer().DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	dir, _ := os.Getwd()
	t, err := template.ParseFiles(dir + templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	r.body = buf.String()
	return nil
}
