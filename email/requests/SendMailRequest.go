package requests

import "email/model"

type SendMailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type EmailRequest struct {
	Automobile model.Automobile `json:"automobile"`
}
