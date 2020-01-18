package api

import (
	"net/http"
	"easysender/model"
	"log"
	"github.com/inwady/easyconfig"
	"easysender/util"
	"fmt"
	"easysender/template"
	"io/ioutil"
)

type TypeSender string

const (
	emailSender TypeSender = "email"
)

var (
	host = easyconfig.GetString("host", "")

	enableAuth = easyconfig.GetInt("smtp_auth.enable", 0)
	smtpSsl = easyconfig.GetInt("smtp_auth.ssl", 1)
	smtpHost = easyconfig.GetString("smtp_auth.host", "smtp.mail.ru")
	smtpEmail = easyconfig.GetString("smtp_auth.email", "e@mail.ru")
	smtpPassword = easyconfig.GetString("smtp_auth.password", "password")

	sender model.Sender
)


func init() {
	var auth *model.EmailSmtpAuth
	if enableAuth != 0 {
		auth = &model.EmailSmtpAuth{
			Host: smtpHost,
			Ssl: smtpSsl != 0,
			Email: smtpEmail,
			Password: smtpPassword,
		}
	}

	var err error
	sender, err = model.InitEmailSender(host, auth)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMessage(w  http.ResponseWriter, r *http.Request) {
	var (
		typeOfSend string
		module string
	)

	var err error
	var ok bool
	if typeOfSend, ok = GetParamFromURI(r, "type"); !ok {
		SendBadParam(w)
		return
	}

	if module, ok = GetParamFromURI(r, "module"); !ok {
		SendBadParam(w)
		return
	}

	switch {
	case typeOfSend == string(emailSender):
		erequest := &EmailRequest{}
		err = ParseJson(r, erequest)
		if err != nil {
			log.Println("parse json", err)
			SendBadParam(w)
			return
		}

		_, server, err := util.EmailParse(erequest.Email)
		if err != nil {
			fmt.Fprintf(w, "bad email")
			return
		}

		email := &model.Email{
			Server: server,
			From: module + "@" + host,
			To: erequest.Email,
			Subject: erequest.Subject,
			MessageHTML: erequest.Data,
		}

		sender.Send(email)
	default:
		SendBadParam(w)
	}
}

func SendMessageTemplate(w  http.ResponseWriter, r *http.Request) {
	var (
		typeOfTemplate string
		emailTo string
		module string
		subject string
	)
	var ok bool
	if typeOfTemplate, ok = GetParamFromURI(r, "type"); !ok {
		SendBadParam(w)
		return
	}

	if module, ok = GetParamFromURI(r, "module"); !ok {
		SendBadParam(w)
		return
	}

	if emailTo, ok = GetParamFromURI(r, "email"); !ok {
		SendBadParam(w)
		return
	}

	if subject, ok = GetParamFromURI(r, "subject"); !ok {
		SendBadParam(w)
		return
	}

	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendServerError(w)
		return
	}

	htmlData, err := template.ExecuteTemplate(typeOfTemplate, jsonBody)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	_, server, err := util.EmailParse(emailTo)
	if err != nil {
		fmt.Fprintf(w, "bad email")
		return
	}

	email := &model.Email{
		Server: server,
		From: module + "@" + host,
		To: emailTo,
		Subject: subject,
		MessageHTML: htmlData,
	}

	sender.Send(email)
}