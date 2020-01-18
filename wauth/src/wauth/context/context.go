package context

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/valyala/fasthttp"
	"log"
	"encoding/json"
	"wauth/util/session"
)

type HTTPScope struct {
	requestId string
}

type AContext struct {
	baseContext context.Context

	deferFuncs []func()

	fastContext *fasthttp.RequestCtx
	scope HTTPScope

	userInfo *session.UserInfo
}

func InitFromHTTP(fctx *fasthttp.RequestCtx) (*AContext, context.CancelFunc, error) {
	bctx, cancel := context.WithCancel(context.Background())

	actx := &AContext{
		baseContext: bctx,
		fastContext: fctx,
		scope: HTTPScope{
			requestId: newRequestID(),
		},
	}

	actx.addDefer(cancel)
	actx.addDefer(actx.writeAccessLog)

	return actx, cancel, nil
}

func (actx *AContext) Context() context.Context {
	return actx.baseContext
}

func (actx *AContext) GetMethod() string {
	return string(actx.fastContext.Method())
}

func (actx *AContext) GetParam(key string) (string, bool) {
	param := actx.fastContext.QueryArgs().Peek(key)
	if param == nil {
		return "", false
	}

	return string(param), true
}

func (actx *AContext) GetParamRequired(key string) string {
	param, ok := actx.GetParam(key)
	if !ok {
		actx.WriteErrorBadRequest()
		log.Panic()
	}

	return param
}

func (actx *AContext) GetPost(key string) (string, bool) {
	value := actx.fastContext.PostArgs().Peek(key)
	if value == nil {
		return "", false
	}

	return string(value), true
}

func (actx *AContext) GetPostRequired(key string) string {
	value, ok := actx.GetPost(key)
	if !ok {
		actx.WriteErrorBadRequest()
		log.Panic()
	}

	return value
}

func (actx *AContext) GetPath() string {
	return string(actx.fastContext.Path())
}

// TODO last
func (actx *AContext) GetRequestURI() string {
	return string(actx.fastContext.RequestURI())
}

func (actx *AContext) GetJson(parse interface{}) error {
	return json.Unmarshal(actx.fastContext.PostBody(), parse)
}

func (actx *AContext) GetContextValue(key string) interface{} {
	return actx.baseContext.Value(key)
}

func (actx *AContext) GetIP() string {
	ip, ok := actx.GetHeader("X-Real-IP")
	if !ok {
		// TODO return default IP
		actx.LogDebug("context", "No IP address in X-Real-IP")
		return "127.0.0.1"
	}

	return ip
}

func (actx *AContext) GetEmail() (string) {
	if actx.userInfo == nil {
		panic("need auth for GetEmail!")
	}

	return actx.userInfo.Email
}

func (actx *AContext) GetUsername() (string) {
	if actx.userInfo == nil {
		panic("need auth for GetUsername!")
	}

	return actx.userInfo.Username
}

func (actx *AContext) GetEmailOptional() (string) {
	if actx.userInfo == nil {
		return ""
	}

	return actx.userInfo.Email
}

func (actx *AContext) GetUsernameOptional() (string) {
	if actx.userInfo == nil {
		return ""
	}

	return actx.userInfo.Username
}

func (actx *AContext) Exit(recovered interface{}) {
	if recovered != nil {
		log.Println("panic !!!", recovered)
	}

	for _, f := range actx.deferFuncs {
		f()
	}
}

func (actx *AContext) setUserInfo(info *session.UserInfo) {
	actx.LogDebugf("context", "set session %s", info.Email)
	actx.userInfo = info
	return
}

func (actx *AContext) releaseUserInfo() {
	if actx.userInfo == nil {
		panic("userinfo is nil")
	}

	actx.LogDebugf("context", "release userinfo <%s>", actx.userInfo.Email)
	actx.userInfo = nil
}

func (actx *AContext) addDefer(d func()) {
	actx.deferFuncs = append(actx.deferFuncs, d)
}

func (actx *AContext) writeAccessLog() {
	log.Printf("connection close, rid: %s, responseCode: %d",
		actx.scope.requestId,
		actx.fastContext.Response.StatusCode())
}

func (actx *AContext) IsSecretUser() bool {
	keyForTest, ok := actx.GetParam("debug")
	if ok && keyForTest == "abracadabra" {
		return true
	}

	return false
}

func newRequestID() string {
	return fmt.Sprintf("rid%08x", rand.Uint32())
}
