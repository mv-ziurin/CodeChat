package model

import (
	"wauth/util"
	"github.com/inwady/easyconfig"
	"errors"
	"log"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/url"
)

var (
	connectSMTPError = errors.New("cannot connect to smtp server")
)

var (
	smtpServer = easyconfig.GetString("smtp.server", "sender.server.ru:80")
	module = easyconfig.GetString("smtp.module", "auth")
	poolSize = easyconfig.GetInt("smtp.pool_size", 4)
)

type Sender interface {
	Send(data interface{})
}

type EmailSender struct {
	pool *util.Pool
}

type Email struct {
	Email string		`json:"email"`
	Subject string		`json:"subject"`
	Data string			`json:"data"`
}

func (es *EmailSender) sendEmailImpl(id int, emailData interface{}) {
	var apiRequest = fmt.Sprintf("http://%s/send?type=email&module=auth", smtpServer)

	email, ok := emailData.(*Email)
	if !ok {
		log.Panic("bad email in pool")
	}

	hc := http.Client{
		Timeout: 2 * time.Second,
	}

	b, err := json.Marshal(email)
	log.Println("Send", string(b))
	if err != nil {
		log.Println("bad email data", err)
		return
	}
	req, _ := http.NewRequest("POST", apiRequest, bytes.NewBuffer(b))

	resp, err := hc.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("bad response from sender email <%s>", email.Email)
	}

	log.Printf("email was sent to <%s>", email.Email)
}

func (es *EmailSender) Send(data interface{}) {
	email, ok := data.(*Email)
	if !ok {
		log.Panic("bad email interface")
	}

	es.pool.ThrowTask(es.sendEmailImpl, email)
}

func InitSender() (Sender, error) {
	es := &EmailSender{}
	es.pool = util.NewPool(poolSize)

	return es, nil
}

// TODO without pool
// TODO import interface from template
func SendTemplate(typeOfTemplate string, emailTo string, subject string, json []byte) error {
	hc := http.Client{
		Timeout: 2 * time.Second,
	}

	serverurl := fmt.Sprintf("http://%s", smtpServer)
	resource := "/send_template"

	data := url.Values{}
	data.Set("type", typeOfTemplate)
	data.Set("email", emailTo)
	data.Set("subject", subject)
	data.Set("module", "auth")

	u, _ := url.ParseRequestURI(serverurl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	log.Println("send data to easysender", urlStr, string(json))
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(json))
	resp, err := hc.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("error", string(data))
		return fmt.Errorf("bad response from sender email <%s>", emailTo)
	}

	log.Printf("email was sent to <%s>", emailTo)
	return nil
}
