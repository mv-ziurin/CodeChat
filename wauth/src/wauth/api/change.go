package api

import (
	"net/http"
	"wauth/context"
	"log"
	"wauth/model/user"
)

type ChangeRequest struct {
	NewPassword string	`json:"newpassword"`
	OldPassword string	`json:"oldpassword"`
}

func ChangeDataRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodPost {
		actx.WriteBadMethod()
		return
	}

	cr := &ChangeRequest{}
	err := actx.GetJson(cr)
	if err != nil {
		actx.WriteErrorBadRequestWithError(err)
		return
	}

	ok, err := actx.SetSession()
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !ok {
		actx.WriteErrorCustom(context.ForbiddenStatus, "")
		return
	}

	success, message, err := user.UpdatePasswordUserByOldPassword(actx.GetEmail(), cr.NewPassword, cr.OldPassword)
	if err != nil {
		if err == user.ErrorPasswordAlreadyWas {
			actx.WriteErrorCustom(context.ErrorPasswordAlreadyWas, "")
			return
		}

		actx.WriteError500(err)
		return
	}

	log.Println("success", success)
	if !success {
		actx.WriteErrorCustom(context.ForbiddenStatus, message)
		return
	}

	actx.WriteOk(message)
}
