package session

import "gitlab.com/gamechain/gameserver/wauth/src/wauth/model/tokener"

type UserInfo struct {
	Email string
	Username string
}

func CheckSession(session string) (userinfo *UserInfo, ok bool, err error) {
	data, ok, err := tokener.GetSavedData("session", session)
	if !ok || err != nil {
		return
	}

	if len(data) != 2 {
		panic("len(data) != 2")
	}

	userinfo = &UserInfo{
		Email: data[0].(string),
		Username: data[1].(string),
	}

	return
}
