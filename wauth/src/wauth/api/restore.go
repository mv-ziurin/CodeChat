package api

import (
	"wauth/context"
	"wauth/model"
	"github.com/inwady/easyconfig"
	"fmt"
	"wauth/validate"
	"wauth/model/tokener"
	"encoding/json"
	"gitlab.com/gamechain/gameserver/easysender/src/easysender/template/data"
	"wauth/model/user"
)

type RestoreApi struct {
	Email string	`json:"email"`
}

type RestorePasswordApi struct {
	Token string		`json:"token"`
	NewPassword string	`json:"newpassword"`
}

var (
	restoreSubject = easyconfig.GetString("restore_email.subject", "")


	restoreScheme = easyconfig.GetString("scheme", "https")
	restoreBasicHost = easyconfig.GetString("base_host", "codechat.ru")
)

func MainRestoreRequest(actx *context.AContext) {
	ra := &RestoreApi{}
	err := actx.GetJson(ra)
	if err != nil {
		actx.WriteErrorBadRequestWithError(err)
		return
	}

	ok, err := user.CheckUserExistByEmail(ra.Email)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !ok {
		actx.WriteError("bad user")
		return
	}

	token, err := tokener.NewSavedData("change_password", []interface{}{
		ra.Email,
	})

	if err != nil {
		actx.WriteError500(err)
		return
	}

	link := fmt.Sprintf("%s://%s?token=%s", restoreScheme, restoreBasicHost, token)
	bytes, _ := json.Marshal(&data.SimpleData{
		Title: "Codechat",
		Link: link,
	})

	err = model.SendTemplate("passrestore", ra.Email, restoreSubject, bytes)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.EasyWriteOk()
}

func RestoreAfterEmailRequest(actx *context.AContext) {
	rpa := &RestorePasswordApi{}
	err := actx.GetJson(rpa)
	if err != nil {
		actx.WriteErrorBadRequestWithError(err)
		return
	}

	// RestoreApi
	data, ok, err := tokener.RemoveSavedData("change_password", rpa.Token)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !ok {
		actx.WriteErrorCustom(context.BadData, "cannot find token")
		return
	}

	if len(data) != 1 {
		panic("len(data) != 1")
	}

	ra := &RestoreApi{
		Email: data[0].(string),
	}

	if !validate.ValidatePasswordForRegister(rpa.NewPassword) {
		actx.WriteErrorProcessed(validate.BadPasswordError)
		return
	}

	err = user.UpdatePasswordUser(ra.Email, rpa.NewPassword)
	if err != nil {
		if err == user.ErrorPasswordAlreadyWas {
			actx.WriteErrorCustom(context.ErrorPasswordAlreadyWas, "")
			return
		}

		actx.WriteError500(err)
		return
	}

	actx.EasyWriteOk()
}
