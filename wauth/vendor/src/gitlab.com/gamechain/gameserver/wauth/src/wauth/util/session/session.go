package session

import "wauth/model/tokener"

type UserInfo struct {
	Email string
}

func CheckSession(session string) (userinfo *UserInfo, ok bool, err error) {
	data, ok, err := tokener.GetSavedData("session", session)
	if !ok || err != nil {
		return
	}

	if len(data) != 1 {
		panic("len(data) != 1")
	}

	userinfo = &UserInfo{
		Email: data[0].(string),
	}

	return
}
