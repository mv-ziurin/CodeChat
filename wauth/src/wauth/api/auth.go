package api

import (
	"github.com/inwady/easyconfig"
	"net/http"
	"wauth/context"
	"wauth/model/event_log"
	"wauth/model/user/auth"
)

var (
	authRedirectSuccess = easyconfig.GetString("auth_redirect_success", "https://rhack.ru/dashboard")
	authRedirectFail = easyconfig.GetString("auth_redirect_fail", "https://rhack.ru?error=1")

	logoutRedirect = easyconfig.GetString("logout_redirect", "https://site.ru")
)

func LoginJSRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodPost {
		actx.WriteBadMethod()
		return
	}

	a := &auth.Auth{}
	err := actx.GetJson(a)
	if err != nil {
		actx.WriteErrorBadRequestWithError(err)
		return
	}

	if err := validateAuthInput(a); err != nil {
		actx.WriteErrorProcessed(err)
		return
	}

	_, message, err := auth.Login(a)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if message != "" {
		actx.WriteErrorCustom(context.ForbiddenStatus, message)
		return
	}

	actx.EasyWriteOk()
}

func LoginRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodPost {
		actx.WriteBadMethod()
		return
	}

	a := &auth.Auth{
		EmailOrUsername: actx.GetPostRequired("email"),
		Password: actx.GetPostRequired("password"),
	}

	if err := validateAuthInput(a); err != nil {
		actx.WriteErrorProcessed(err)
		return
	}

	ui, message, err := auth.Login(a)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if message != "" {
		actx.RedirectTemporary(authRedirectFail)
		return
	}

	err = actx.MakeAuth(ui.Email, ui.Username)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.RedirectTemporary(authRedirectSuccess)
	event_log.LogEventFromContext(event_log.LoginEvent, actx)
}

func LogoutRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodPost {
		if actx.GetMethod() == http.MethodGet {
			// TODO redirect user fine
			actx.RedirectTemporary(logoutRedirect)
			return
		}

		actx.WriteBadMethod()
		return
	}

	ok, err := actx.SetSession()
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.RemoveSessionOnlyCookie()

	if !ok {
		actx.RedirectTemporary(logoutRedirect)
		return
	}

	// log
	event_log.LogEventFromContext(event_log.LogoutEvent, actx)

	err = actx.RemoveSession()
	if err != nil {
		actx.WriteError500(err)
		return
	}

	// goodbye, user!
	actx.RedirectTemporary(logoutRedirect)
}
