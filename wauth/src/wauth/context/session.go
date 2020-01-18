package context

import (
	"fmt"
	"math/rand"
	"github.com/valyala/fasthttp"
	"time"
	"github.com/inwady/easyconfig"
	"wauth/model/tokener"
	session_util "wauth/util/session"
	"errors"
)

var (
	sessionRandom *rand.Rand

	cookieKey = easyconfig.GetString("cookie.key", "hauth")
	cookieDomain = easyconfig.GetString("cookie.domain", "*.rhack.ru")
)

func init() {
	source := rand.NewSource(time.Now().Unix())
	sessionRandom = rand.New(source)
}

// generate new session for user
func (actx *AContext) NewSession(userInfo *session_util.UserInfo) error {
	sessionString := newSessionString(userInfo.Email)

	err := tokener.NewSavedDataCustom(sessionString, "session", []interface{}{
		userInfo.Email,
		userInfo.Username,
	})

	if err != nil {
		return err
	}

	setSessionCookie(actx, sessionString)
	actx.setUserInfo(userInfo)
	return nil
}

func (actx *AContext) SetSession() (ok bool, err error) {
	session, ok := getSessionCookie(actx)
	if !ok {
		return
	}

	userinfo, ok, err := session_util.CheckSession(session)
	if !ok || err != nil {
		return
	}

	actx.setUserInfo(userinfo)

	ok = true
	return
}

func (actx *AContext) RemoveSessionOnlyCookie() {
	cookie := &fasthttp.Cookie{}
	cookie.SetKey(cookieKey)
	cookie.SetValue("deleted")
	cookie.SetExpire(time.Unix(0, 0))
	cookie.SetDomain(cookieDomain)
	cookie.SetHTTPOnly(true)

	actx.SetCookie(cookie)
}

func (actx *AContext) RemoveSession() error {
	session, ok := getSessionCookie(actx)
	if !ok {
		return errors.New("RemoveSession: bad cookie")
	}

	_, _, err := tokener.RemoveSavedData("session", session)
	if err != nil {
		return err
	}

	actx.LogDebugf("context", "RemoveSession: user <%s> was removed", actx.GetEmail())

	actx.releaseUserInfo()
	return nil
}

func newSessionString(email string) string {
	return fmt.Sprintf("%08x%08x%08x",
		sessionRandom.Uint32(),
		sessionRandom.Uint32(),
		sessionRandom.Uint32(),
	)
}

func setSessionCookie(actx  *AContext, token string) {
	cookie := &fasthttp.Cookie{}
	cookie.SetKey(cookieKey)
	cookie.SetValue(token)
	cookie.SetDomain(cookieDomain)
	cookie.SetHTTPOnly(true)
	cookie.SetExpire(time.Unix(time.Now().Unix() + 3600 * 24 * 365, 0))
	cookie.SetPath("/")

	// TODO set secure
	// cookie.SetSecure(true)

	actx.SetCookie(cookie)
}

func getSessionCookie(actx *AContext) (sessionId string, ok bool) {
	token, ok := actx.GetCookie(cookieKey)
	if !ok {
		return
	}

	sessionId, ok = parseSessionString(token)
	if !ok {
		return
	}

	ok = true
	return
}

func parseSessionString(session string) (sessionId string, ok bool) {
	sessionId = session

	ok = true
	return
}
