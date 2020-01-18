package context

import (
	"wauth/model/user"
	"wauth/util/session"
)

type RegisterData struct {
	Username string		`json:"username"`
	Email string 		`json:"email"`
	Password string		`json:"password"`
}

func (actx *AContext) MakeRegister(rd *RegisterData) error {
	/* save hash */
	err := user.AddNewUser(rd.Username, rd.Email, rd.Password)
	if err != nil {
		return err
	}

	return nil
}

func (actx *AContext) MakeAuth(email string, username string) error {
	// TODO disable role
	return actx.NewSession(&session.UserInfo{
		Email: email,
		Username: username,
	})
}
