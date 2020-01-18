package captcha

import (
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/inwady/easyconfig"
	"wauth/context"
)

// TODO check actx only request

var (
	captchaKey = easyconfig.GetString("captcha.private_key", "")
)

func init() {
	recaptcha.Init(captchaKey)
}

func CheckCaptcha(actx *context.AContext, captcha string) (bool, error) {
	ok, err := recaptcha.Confirm(actx.GetIP(), captcha)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func CheckCaptchaFromGetParam(actx *context.AContext) (bool, error) {
	captcha, ok := actx.GetParam("captcha")
	if !ok {
		return false, nil
	}

	return CheckCaptcha(actx, captcha)
}
