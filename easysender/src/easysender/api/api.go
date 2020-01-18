package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
)

type EmailRequest struct {
	Email string		`json:"email"`
	Subject string		`json:"subject"`
	Data string			`json:"data"`
}

func GetParamFromURI(r * http.Request, key string) (string, bool) {
	params, ok := r.URL.Query()[key]
	if !ok || len(params) < 1 {
		return "", false
	}

	return params[0], true
}

func ParseJson(r *http.Request, data interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, data); err != nil {
		return err
	}

	return nil
}

func SendBadParam(w http.ResponseWriter) {
	log.Println("bad param")
	fmt.Fprintf(w, "bad param")
}

func SendServerError(w http.ResponseWriter) {
	log.Println("server error")
	fmt.Fprintf(w, "server error !!!")
}