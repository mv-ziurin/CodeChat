package api

import (
	"wauth/context"
	"net/http"
	"wauth/model"
	"github.com/inwady/easyconfig"
	"fmt"
	"encoding/json"
	"wauth/model/tokener"
	"wauth/model/notifycli"
	"gitlab.com/gamechain/gameserver/easysender/src/easysender/template/data"
	"wauth/model/event_log"
	"wauth/model/user"
)

// TODO monitor!
/* debug API */
type ResponseCode struct {
	Code string 	`json:"code"`
}

var (
	registerFlag = easyconfig.GetInt("enable.register", 1)

	registerScheme = easyconfig.GetString("scheme", "https")
	registerHost = easyconfig.GetString("host", "auth.rhack.ru")
	redirectSuccess = easyconfig.GetString("register_redirect_success", "https://rhack.ru/dashboard")

	redirectFail = easyconfig.GetString("register_redirect_fail", "https://rhack.ru?error=1")

	emailSubject = easyconfig.GetString("request_email.subject", "Auth from RHack!")

	textIntoNotifyIfSuccess = easyconfig.GetString("notifier.hello_text", "Hello!")
)

func RegisterRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodPost {
		actx.WriteBadMethod()
		return
	}

	if registerFlag == 0 {
		actx.WriteErrorCustom(context.ForbiddenStatus, "disable register")
		return
	}

	ok, err := actx.SetSession()
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if ok {
		actx.WriteError("user has already had cookie")
		return
	}

	register := &context.RegisterData{}
	err = actx.GetJson(register)
	if err != nil {
		actx.WriteErrorBadRequest()
		return
	}

	ok, err = checkCaptchaByAPI(actx)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !actx.IsSecretUser() && !ok {
		actx.WriteErrorCustom(context.ForbiddenStatus, "bad captcha")
		return
	}

	if err = validateRegisterInput(register); err != nil {
		actx.WriteErrorProcessed(err)
		return
	}

	ok, err = user.CheckUserExist(register.Email, register.Username)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if ok {
		actx.WriteErrorStatus(context.ErrorUserExist)
		return
	}

	token, err := tokener.NewSavedData("register", []interface{}{
		register.Email,
		register.Username,
		register.Password,
	})

	if err != nil {
		actx.WriteError500(err)
		return
	}

	if actx.IsSecretUser() {
		actx.WriteOk(&ResponseCode{token})
		return
	}

	link := fmt.Sprintf("%s://%s/register?token=%s", registerScheme, registerHost, token)

	bytes, _ := json.Marshal(&data.SimpleData{
		Title: "Codechat",
		Link: link,
	})

	err = model.SendTemplate("greeting", register.Email, emailSubject, bytes)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	/*** end ***/

	actx.EasyWriteOk()
	event_log.LogEventFromContext(event_log.RequestEvent, actx)
}

func RegisterFinal(actx *context.AContext) {
	token := actx.GetParamRequired("token")

	dataRegister, ok, err := tokener.GetSavedData("register", token)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !ok {
		actx.RedirectTemporary(redirectFail)
		return
	}

	register := &context.RegisterData{
		Email: dataRegister[0].(string),
		Username: dataRegister[1].(string),
		Password: dataRegister[2].(string),
	}

	/* save hash */
	err = actx.MakeRegister(register)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	err = actx.MakeAuth(register.Email, register.Username)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.RedirectTemporary(redirectSuccess)
	event_log.LogEventFromContext(event_log.RegisterEvent, actx)

	// user's been registered
	err = notifycli.NotifySendText(register.Username, textIntoNotifyIfSuccess)
	if err != nil {
		actx.LogError(err)
	}
}
