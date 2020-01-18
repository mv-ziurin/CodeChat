package api

import (
	"wauth/context"
	"net/http"
	"wauth/model/user"
)

func EmailsAdminRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodGet {
		actx.WriteBadMethod()
		return
	}

	emails, err := user.GetAllEmails()
	if err != nil {
		actx.WriteErrorBadRequestWithError(err)
		return
	}

	actx.WriteOk(emails)
}

