package customer

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"net/smtp"

	"coffee-culture.uk/internal/config"
	qrcode "github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SendCustomerQR(customerEmail string, customerID primitive.ObjectID, templatePath string, customerUsername string) error {
	//generate QR code
	qrCodeData := customerID.Hex()
    qrCode, err := qrcode.Encode(string(qrCodeData), qrcode.Medium, 256)
    if err != nil {
        return err
    }
	qrCodeBase64 := base64.StdEncoding.EncodeToString(qrCode)

	//load html template file
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
        return err
    }

	//data to populate the template
	data := struct {
		CustomerUsername string
		qrCodeBase64 string
	}{
		CustomerUsername: customerUsername,
		qrCodeBase64: qrCodeBase64,
	}

	//Execute the template with the data
	err = t.Execute(&body, data)
	if err != nil {
		return err
	}

	SendMailSimpleHTML("hi", body.String(), []string{customerEmail})
	
	return nil


}

func SendMailSimpleHTML(subject string, html string, to []string) {
	auth := smtp.PlainAuth(
		"",
		config.AppConfig().Email.CCEmail,
		config.AppConfig().Email.AppPassword,
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg:= "Subject: " + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		config.AppConfig().Email.CCEmail,
		to,
		[]byte(msg),
	)
	if err != nil {
		panic(err)
	}
}