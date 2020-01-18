package api

import (
	"errors"
	"log"
	"wauth/context"
	"wauth/model"
	"wauth/model/captcha"
	"wauth/model/user/auth"
	"wauth/validate"
)

var (
	sender model.Sender

	BadRoleError = errors.New("bad role")
)

func init() {
	var err error
	sender, err = model.InitSender()
	if err != nil {
		log.Fatal(err)
	}
}

func checkCaptchaByAPI(actx *context.AContext) (bool, error) {
	captchaOk, err := captcha.CheckCaptchaFromGetParam(actx)
	if err != nil {
		return false, err
	}

	if !captchaOk {
		return false, nil
	}

	return true, nil
}

func validateAuthInput(a *auth.Auth) error {
	// TODO validate username

	if !validate.ValidatePassword(a.Password) {
		return validate.BadPasswordError
	}

	return nil
}

func validateRegisterInput(register *context.RegisterData) error {
	if !validate.ValidateEmail(register.Email) {
		return validate.BadEmailError
	}

	if !validate.ValidateUsername(register.Username) {
		return validate.BadUsernameError
	}

	if !validate.ValidatePasswordForRegister(register.Password) {
		return validate.BadPasswordError
	}

	return nil
}
