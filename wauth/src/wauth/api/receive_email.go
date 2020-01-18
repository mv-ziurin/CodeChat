package api

import (
	"wauth/context"
	"net/http"
	"wauth/model/receive_email"
	"github.com/inwady/easyconfig"
	"wauth/model"
	"wauth/validate"
)

var (
	receiveEmailFlag = easyconfig.GetInt("enable.receive_email", 1)

	subjectSubscribe = easyconfig.GetString("subscribe.subject", "Subscribing")
)

func ReceiveEmailRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodPut {
		actx.WriteBadMethod()
		return
	}

	if receiveEmailFlag == 0 {
		actx.WriteErrorCustom(context.ForbiddenStatus, "disable receive email")
		return
	}

	re := &receive_email.ReceiveEmail{}
	err := actx.GetJson(re)
	if err != nil {
		actx.WriteErrorBadRequest()
		return
	}

	if !validate.ValidateEmail(re.Email) {
		actx.WriteErrorBadRequest()
		return
	}

	err = receive_email.NewReceiveEmail(re)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.EasyWriteOk()

	// after response
	model.SendTemplate("subscribe", re.Email, subjectSubscribe, []byte("{}"))
}

/* admin */
func GetAllEmailsRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodGet {
		actx.WriteBadMethod()
		return
	}

	emails, err := receive_email.GetAllEmails()
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.WriteOk(emails)
}
