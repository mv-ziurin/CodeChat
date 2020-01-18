package validate

import (
	"errors"
	"regexp"
	"github.com/inwady/easyconfig"
	"strings"
)

var (
	BadEmailError = errors.New("bad email")
	BadPasswordError = errors.New("bad password")
	BadUsernameError = errors.New("bad username")
)

var (
	usernameLengthMin = easyconfig.GetInt("validate.username_len_min", 4)
	usernameLengthMax = easyconfig.GetInt("validate.username_len_max", 20)

	passwordLengthMin = easyconfig.GetInt("validate.password_len_min", 6)
	passwordLengthMax = easyconfig.GetInt("validate.password_len_max", 30)

	passwordSpecialFlag = easyconfig.GetInt("validate.password_special_symbols_enable", 1)
	passwordSpecial     = easyconfig.GetString("validate.password_special_symbols", " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	passwordNeedSymbols = easyconfig.GetArrayString("validate.password_need_symbols", []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	})

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	usernameRegexp = regexp.MustCompile("^[a-zA-Z0-9_-]*$")
)

func ValidateEmail(email string) bool {
	if len(email) < 3 {
		return false
	}

	return emailRegexp.MatchString(email)
}

func ValidateUsername(username string) bool {
	if (len(username) < usernameLengthMin || len(username) > usernameLengthMax) {
		return false
	}

	return usernameRegexp.MatchString(username)
}

func ValidatePassword(password string) bool {
	if len(password) > passwordLengthMax {
		return false
	}

	return true
}

func ValidatePasswordForRegister(password string) bool {
	if len(password) < passwordLengthMin || len(password) > passwordLengthMax {
		return false
	}

	if passwordSpecialFlag != 0 && !strings.ContainsAny(password, passwordSpecial) {
		return false
	}

	for _, v := range passwordNeedSymbols {
		if !strings.ContainsAny(password, v) {
			return false
		}
	}

	return true
}