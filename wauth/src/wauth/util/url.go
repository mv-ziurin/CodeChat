package util

import (
	"fmt"
	"net/url"
)

func UrlEncodeURI(scheme string, host string, resource string, uriValues url.Values) string {
	serverUrl := fmt.Sprintf("%s://%s", scheme, host)

	u, _ := url.ParseRequestURI(serverUrl)
	u.Path = resource
	u.RawQuery = uriValues.Encode()
	return fmt.Sprintf("%v", u)
}
