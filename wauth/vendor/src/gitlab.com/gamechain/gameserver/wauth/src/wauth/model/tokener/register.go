package tokener

import (
	"fmt"
	"log"
)

type RegisterData struct {
	Email string 		`json:"email"`
	Password string		`json:"password"`
}

func NewRegisterData(token string, r *RegisterData) (err error) {
	_, err = callFunction("storage.register:add", []interface{}{
		token,
		r.Email,
		r.Password,
	})

	return
}

func GetRegisterData(token string) (*RegisterData, bool, error) {
	resp, err := callFunction("storage.register:remove", []interface{}{token})
	if err != nil {
		return nil, false, err
	}

	var (
		tuple []interface{}
		ok bool
	)

	log.Println(resp.Data)
	if len(resp.Data) != 1 {
		return nil, false, nil
	}

	if tuple, ok = resp.Data[0].([]interface{}); !ok {
		return nil, false, fmt.Errorf("tuple error, resp.Data[0] -> []interface{}")
	}

	if len(tuple) == 0 {
		return nil, false, nil
	}

	// token, time, email, password
	if len(tuple) != 4 {
		return nil, false, fmt.Errorf("tuple error, len(tuple) != 4")
	}

	return &RegisterData{
		Email: tuple[2].(string),
		Password: tuple[3].(string),
	}, true, nil
}
