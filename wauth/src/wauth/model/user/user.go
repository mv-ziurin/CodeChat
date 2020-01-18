package user

import (
	"wauth/database"
	"log"
	"wauth/util"
	"errors"
	"database/sql"
	"wauth/model/password_history"
)

var (
	sqlUserInsert = "INSERT INTO users (email, username, hash) VALUES ($1, $2, $3)"

	checkUserSelect = "SELECT DISTINCT COUNT(*) FROM users WHERE email = $1 OR username = $2"
	checkUserSelectByEmail = "SELECT DISTINCT COUNT(*) FROM users WHERE email = $1"
	sqlSelectUserByEmail = "SELECT email, username, steam_id FROM  users WHERE email = $1"

	sqlUpdateSteamIdByEmail = "UPDATE users SET steam_id = $1 WHERE email = $2"

	// checkValidUsername = "SELECT 1 FROM users WHERE email = $1 AND hash = $2"
	sqlUpdatePassword                  = "UPDATE users SET hash = $1 WHERE email = $2"
	sqlUpdatePasswordIfOldPasswordTrue = "UPDATE users SET hash = $1 WHERE email = $2 AND hash = $3"

	sqlSelectEmails = "SELECT email FROM users"
)

var (
	ErrorPasswordAlreadyWas = errors.New("password already in history")
)

type User struct {
	Email string 		`json:"email"`
	Username string		`json:"username"`
	SteamId string		`json:"steam_id"`
}

func GetUserByEmail(email string) (*User, bool, error) {
	dbconn := database.GetConnection()

	var steamId sql.NullString

	u := &User{}
	err := dbconn.QueryRow(sqlSelectUserByEmail, email).Scan(&u.Email, &u.Username, &steamId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	if steamId.Valid {
		u.SteamId = steamId.String
	}

	return u, true, nil
}

func CheckUserExistByEmail(email string) (bool, error) {
	// TODO error
	dbconn := database.GetConnection()

	var count int64
	err := dbconn.QueryRow(checkUserSelectByEmail, email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func CheckUserExist(email string, username string) (bool, error) {
	// TODO error
	dbconn := database.GetConnection()

	var count int64
	err := dbconn.QueryRow(checkUserSelect, email, username).Scan(&count)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func AddNewUser(username string, email string, password string) (err error) {
	hexHash := util.MainHashHexify(password)

	// TODO in one transaction !!!
	dconn := database.GetConnection()
	_, err = dconn.Exec(sqlUserInsert, email, username, hexHash)
	if err != nil {
		return
	}

	err = password_history.StoreOldPassword(email, hexHash)
	if err != nil {
		return
	}

	return
}

func GetAllEmails() ([]string, error) {
	const defaultEmailSize = 10

	dbconn := database.GetConnection()

	rows, err := dbconn.Query(sqlSelectEmails)
	if err != nil {
		return nil, err
	}

	var (
		email string
		emails = make([]string, 0, defaultEmailSize)
	)

	for rows.Next() {
		err = rows.Scan(&email)
		if err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	return emails, nil
}
