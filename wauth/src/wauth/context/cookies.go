package context

import (
	"github.com/valyala/fasthttp"
)

// TODO set all cookies
func (actx *AContext) SetCookie(cookie *fasthttp.Cookie) {
	actx.fastContext.Response.Header.SetCookie(cookie)
}

func (actx *AContext) SetHeader(key string, value string) {
	actx.fastContext.Response.Header.Set(key, value)
}

func (actx *AContext) GetCookie(key string) (string, bool) {
	bytes := actx.fastContext.Request.Header.Cookie(key)
	if bytes == nil {
		return "", false
	}

	return string(bytes), true
}

func (actx *AContext) GetHeader(key string) (string, bool) {
	bytes := actx.fastContext.Request.Header.Peek(key)
	if bytes == nil {
		return "", false
	}

	return string(bytes), true
}
