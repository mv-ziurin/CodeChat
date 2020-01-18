package user

import (
	"wauth/model/password_history"
	"errors"
	"wauth/database"
	"wauth/util"
)

func UpdatePasswordUserByOldPassword(email string, newPassword string, oldPassword string) (success bool, message string, err error) {
	if newPassword == oldPassword {
		return false, "passwords are equal", nil
	}

	newPasswordHash := util.MainHashHexify(newPassword)
	oldPasswordHash := util.MainHashHexify(oldPassword)

	ok, err := password_history.IsPasswordInOldPassword(email, newPasswordHash)
	if err != nil {
		return
	}

	if ok {
		err = ErrorPasswordAlreadyWas
		return
	}

	dbconn := database.GetConnection()
	resp, err := dbconn.Exec(sqlUpdatePasswordIfOldPasswordTrue, newPasswordHash, email, oldPasswordHash)
	if err != nil {
		return
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil { return }
	if rowsAffected != 1 {
		return false, "bad old password", nil
	}

	err = password_history.StoreOldPassword(email, newPasswordHash)
	if err != nil {
		return
	}

	return true, "success", nil
}

func UpdatePasswordUser(email string, newPassword string) (err error) {
	newPasswordHash := util.MainHashHexify(newPassword)
	ok, err := password_history.IsPasswordInOldPassword(email, newPasswordHash)
	if err != nil {
		return
	}

	if ok {
		return ErrorPasswordAlreadyWas
	}

	dbconn := database.GetConnection()
	resp, err := dbconn.Exec(sqlUpdatePassword, newPasswordHash, email)
	if err != nil {
		return
	}

	rowAffected, err := resp.RowsAffected()
	if err != nil { return }
	if rowAffected != 1 {
		return errors.New("UpdatePasswordUser, bad email")
	}

	err = password_history.StoreOldPassword(email, newPasswordHash)
	if err != nil {
		return
	}

	return
}
