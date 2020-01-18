package api

import (
	"wauth/context"
	"net/http"
	"wauth/model/user"
)

func GetConfigRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodGet {
		actx.WriteBadMethod()
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

	user, ok, err := user.GetUserByEmail(actx.GetEmail())
	if err != nil || !ok {
		actx.WriteError500(err)
		return
	}

	actx.WriteOk(user)
}
