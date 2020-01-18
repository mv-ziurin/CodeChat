package auth

import (
	"database/sql"
	"wauth/database"
	"wauth/util"
	"wauth/util/session"
)

// TODO lower ??

const (
	sqlUserSelect = "SELECT email, username FROM users WHERE (email = $1 OR username = $1) AND hash = $2"
)

type Auth struct {
	EmailOrUsername string		`json:"email"`
	Password string				`json:"password"`
}

// ok, backMessage, error
func Login(auth *Auth) (*session.UserInfo, string, error) {
	hexHash := util.MainHashHexify(auth.Password)

	dconn := database.GetConnection()

	ui := &session.UserInfo{}
	err := dconn.QueryRow(sqlUserSelect, auth.EmailOrUsername, hexHash).Scan(&ui.Email, &ui.Username)
	if err == sql.ErrNoRows {
		return nil, "bad login/password", nil
	}

	if err != nil {
		return nil, "", err
	}

	return ui, "", nil
}
