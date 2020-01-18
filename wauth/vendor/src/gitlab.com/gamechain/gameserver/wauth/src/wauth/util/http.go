package util

import (
	"net/url"
	"time"
	"net/http"
	"errors"
	"log"
)

var (
	StatusCodeError = errors.New("bad status code")
)

func IsHttpError(status int) bool {
	return status >= 500 && status < 600
}

func SendGetRequest(scheme string,
		server string,
		api string,
		values url.Values,
		timeout time.Duration) error {
	client := http.Client{
		Timeout: timeout * time.Millisecond,
	}

	log.Println("Test", timeout)
	request := UrlEncodeURI(scheme, server, api, values)

	resp, err := client.Get(request)
	if err != nil {
		return err
	}

	if IsHttpError(resp.StatusCode) {
		return StatusCodeError
	}

	return nil
}
