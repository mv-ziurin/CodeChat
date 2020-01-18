package model

import (
	"log"
	"bytes"
	"github.com/inwady/easyconfig"
	"easysender/util"
	"net/smtp"
	"strings"
	"net"
	"io"
	"crypto/tls"
	"time"
)

var (
	poolSize = easyconfig.GetInt("email.pool_size", 4)
)

type EmailSmtpAuth struct {
	Ssl bool
	Host string
	Email string
	Password string
}

type EmailSender struct {
	pool *util.Pool

	localHost string
	authInfo *EmailSmtpAuth
}

type Email struct {
	Server string
	From string

	To string
	Subject string
	MessageHTML string
}

func addSmtpPort(server string, ssl bool) string {
	if strings.Contains(server, ":") {
		return server
	}

	var result string
	if ssl {
		result = server + ":465"
	} else {
		result = server + ":25"
	}

	return result
}

func connectSmtpServer(host string, ssl bool) (*smtp.Client, error) {
	server := strings.TrimSuffix(host, ".")
	server = addSmtpPort(server, ssl)

	var (
		err error
		cn *smtp.Client
	)

	if ssl {
		tlsconfig := &tls.Config {
			InsecureSkipVerify: true,
			ServerName:         host,
		}

		tmp, err := tls.Dial("tcp", server, tlsconfig)
		if err != nil {
			return nil, err
		}

		cn, err = smtp.NewClient(tmp, host)
	} else {
		cn, err = smtp.Dial(server)
	}

	if err != nil {
		return nil, err
	}

	return cn, nil
}

func (es *EmailSender) sendHeaders(writer io.Writer, data map[string]string) error {
	var err error
	for header, value := range data {
		if _, err = writer.Write([]byte(header + ": " + value + "\r\n")); err != nil {
			return err
		}

		log.Println(header + ": " + value)
	}

	_, err = writer.Write([]byte("\r\n"))
	if err != nil {
		return err
	}

	return nil
}

func (es *EmailSender) sendEmail(cn *smtp.Client, email *Email) error {
	var (
		auth = es.authInfo != nil
		err error
	)

	if auth {
		err = cn.Mail(es.authInfo.Email)
	} else {
		err = cn.Mail(email.From)
	}

	if err != nil {
		return err
	}

	if err = cn.Rcpt(email.To); err != nil {
		return err
	}

	writer, err := cn.Data()
	if err != nil {
		return err
	}

	defer writer.Close()

	headerMap := map[string]string{
		"Subject": email.Subject,
		"Date": time.Now().Format(time.RFC1123Z),
		"To": email.To,
		"Content-Type": "Text/HTML; charset=UTF-8",
	}

	if auth {
		headerMap["From"] = es.authInfo.Email
	} else {
		headerMap["From"] = email.From
	}

	err = es.sendHeaders(writer, headerMap)
	if err != nil {
		return err
	}

	buf := bytes.NewBufferString(email.MessageHTML)
	if _, err = buf.WriteTo(writer); err != nil {
		return err
	}

	return nil
}

func (es *EmailSender) sendEmailImpl(id int, emailData interface{}) {
	email, ok := emailData.(*Email)
	if !ok {
		log.Panic("bad email in pool")
	}

	var (
		auth = es.authInfo != nil
		client *smtp.Client
		err error
	)

	if auth {
		client, err = connectSmtpServer(es.authInfo.Host, es.authInfo.Ssl)
	} else {
		addrs, err := net.LookupMX(es.localHost)
		if err != nil {
			log.Println(err)
			return
		}

		// connect to all addresses
		for _, addr := range addrs {
			client, err = connectSmtpServer(addr.Host, false) // TODO change this
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		log.Println("cannot connect, err: ", err)
		return
	}

	defer client.Close()

	// need auth
	if auth {
		auth := smtp.PlainAuth("",
			es.authInfo.Email,
			es.authInfo.Password,
			es.authInfo.Host)

		err := client.Auth(auth)
		if err != nil {
			log.Println("cannot auth, err:", err)
			return
		}
	} else {
		if err = client.Hello(es.localHost); err != nil {
			log.Println("cannot localHost, err:", err)
			return
		}
	}

	if err = es.sendEmail(client, email); err != nil {
		log.Println("cannot send email, err:", err)
		return
	}

	if err = client.Quit(); err != nil {
		log.Println("cannot quit, err:", err)
		return
	}

	log.Println("success email!")
}

func (es *EmailSender) Send(data interface{}) bool {
	email, ok := data.(*Email)
	if !ok {
		return false
	}

	es.pool.ThrowTask(es.sendEmailImpl, email)
	return true
}

func InitEmailSender(localHost string, auth *EmailSmtpAuth) (Sender, error) {
	es := &EmailSender{}
	es.pool = util.NewPool(poolSize)

	es.localHost = localHost
	es.authInfo = auth

	return es, nil
}
