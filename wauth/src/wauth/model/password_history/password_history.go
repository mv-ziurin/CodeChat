package password_history

import (
	"wauth/database"
	"time"
)

var (
	sqlSelectOldPassword = "SELECT DISTINCT COUNT(*) FROM password_history WHERE email = $1 AND old_password = $2"
	sqlInsertOldPassword = "INSERT INTO password_history (email, old_password, old_date) VALUES ($1, $2, $3)"
)

func StoreOldPassword(email string, oldPassword string) (err error) {
	dbconn := database.GetConnection()
	_, err = dbconn.Exec(sqlInsertOldPassword, email, oldPassword, time.Now())
	if err != nil {
		return
	}
	return
}

func IsPasswordInOldPassword(email string, passwordWhichWillBeChecked string) (ok bool, err error) {
	dbconn := database.GetConnection()

	var count int64
	err = dbconn.QueryRow(sqlSelectOldPassword, email, passwordWhichWillBeChecked).Scan(&count)
	if err != nil {
		return
	}

	if count > 0 {
		ok = true
		return
	}

	return
}
