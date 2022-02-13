package logic

import (
	"bytes"
	"email/model"
	"email/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type SendMailLogic interface {
	SendMail(ctx *gin.Context)
}

type sendMailLogic struct {
}

func NewSendMailLogic() SendMailLogic {
	return &sendMailLogic{}
}

func (s *sendMailLogic) SendMail(ctx *gin.Context) {
	var sendRequest requests.EmailRequest
	var request requests.SendMailRequest
	if err := ctx.ShouldBindJSON(sendRequest); err != nil {
		log.Println(err.Error())
		return
	}

	user := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	host := os.Getenv("MAIL_HOST")

	auth := smtp.PlainAuth("", user, password, host)

	templateData := struct {
		Part model.Automobile
	}{
		Part: sendRequest.Automobile,
	}
	if err := parseTemplate(&request, "../template/template.html", templateData); err == nil {
		ok, err := sendEmail(auth, &request)
		if err != nil {
			log.Println(err.Error())
		}
		if ok {
			fmt.Println("Email sent successfully")
		}
	}
}

func sendEmail(auth smtp.Auth, r *requests.SendMailRequest) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.Body)
	addr := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))

	if err := smtp.SendMail(addr, auth, r.From, r.To, msg); err != nil {
		return false, err
	}
	return true, nil
}

func parseTemplate(r *requests.SendMailRequest, templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}
