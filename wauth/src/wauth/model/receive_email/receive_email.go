package receive_email

import (
	"wauth/database"
)

var (
	sqlInsertEmail = "INSERT INTO received_email (email) VALUES ($1)"
	sqlSelectEmails = "SELECT email FROM received_email"
)

type ReceiveEmail struct {
	Email string `json:"email"`
}

func NewReceiveEmail(re *ReceiveEmail) (error) {
	dbconn := database.GetConnection()

	_, err := dbconn.Exec(sqlInsertEmail, re.Email)
	if err != nil {
		return err
	}

	return nil
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
