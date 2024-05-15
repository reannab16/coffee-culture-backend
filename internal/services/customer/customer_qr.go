package customer

import (
	"bytes"
	"encoding/base64"
	"html/template"

	"coffee-culture.uk/internal/config"
	qrcode "github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
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

	SendGoMail();

	return nil

}


func SendGoMail() error {
	m := gomail.NewMessage()
m.SetHeader("From", config.AppConfig().Email.CCEmail)
m.SetHeader("To", config.AppConfig().Email.CCEmail)
// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
m.SetHeader("Subject", "Hello!")
m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
// m.Attach("/home/Alex/lolcat.jpg")

d := gomail.NewDialer("smtp.example.com", 587, config.AppConfig().Email.CCEmail, config.AppConfig().Email.AppPassword)

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
	return err
}

return nil
}