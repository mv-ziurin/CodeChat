package main

import (
	"net/http"
	"log"
	"github.com/inwady/easyconfig"
	"easysender/api"
	"strconv"
)

var (
	port = easyconfig.GetInt("bind", 5555)
)

func main() {
	http.HandleFunc("/send", api.SendMessage)
	http.HandleFunc("/send_template", api.SendMessageTemplate)

	errors := make(chan error, 1)
	go func() {
		errors <- http.ListenAndServe(":" + strconv.Itoa(port), nil)
	}()

	log.Println("start listen, port", strconv.Itoa(port))

	for err := range errors {
		log.Fatalln(err)
	}
}
