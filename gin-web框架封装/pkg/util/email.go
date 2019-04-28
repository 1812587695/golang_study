package util

import (
	"net/smtp"
	"html/template"
	"bytes"
	"fmt"
	"strings"
)

const (
	username = "@163.com"
	password = ""
	host = "smtp.163.com:25"
)

func SendEmail(to string, name string, message string, subject string) error {
	auth := smtp.PlainAuth("", username, password, strings.Split(host, ":")[0])

	t, _ := template.ParseFiles("pkg/util/email-template.html")

	var body bytes.Buffer

	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: " + subject + "\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Name string
		Email string
		Message string
	}{
		Name: name,
		Email: to,
		Message: message,
	})

	return smtp.SendMail(host, auth, username, []string{to}, body.Bytes())
}