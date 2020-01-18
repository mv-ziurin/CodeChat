package tokener

import (
	"github.com/inwady/easyconfig"
	"time"
	"github.com/tarantool/go-tarantool"
	"log"
	"fmt"
	"gitlab.com/gamechain/gameserver/wauth/src/wauth/util"
)

var (
	tntServer = easyconfig.GetString("tarantool_tokener.server", "notifier_tarantool")
	tntUsername = easyconfig.GetString("tarantool_tokener.username", "tester")
	tntPassword = easyconfig.GetString("tarantool_tokener.password", "password")

	tntTimeout = time.Duration(easyconfig.GetInt("tarantool_tokener.timeout", 2000))

	conn *tarantool.Connection
)

func init() {
	opts := tarantool.Opts{
		User: tntUsername,
		Pass: tntPassword,
		Timeout: tntTimeout * time.Millisecond,

		Reconnect: 1 * time.Second,
		MaxReconnects: 3,
	}

	// TODO without timeout
	log.Printf("start sleep for tarantool tokener")
	time.Sleep(6 * time.Second)

	var err error
	conn, err = tarantool.Connect(tntServer, opts)
	if err != nil {
		log.Fatalf("tokener, tarantool connection refused: %v", err)
	}
}

func callFunction(fname string, args []interface{}) (resp *tarantool.Response, err error) {
	resp, err = conn.Call(fname, args)
	if err != nil {
		return
	}

	if resp.Code != tarantool.OkCode {
		return nil, fmt.Errorf("return bad code %d", resp.Code)
	}

	// TODO check tuple
	return
}

func NewSavedDataCustom(token string, module string, data []interface{}) (err error) {
	data = append([]interface{}{token}, data...)

	callName := fmt.Sprintf("storage.%s:add", module)
	_, err = callFunction(callName, data)

	return
}

func NewSavedData(module string, data []interface{}) (token string, err error) {
	token = util.GenerateDigitToken32()
	err = NewSavedDataCustom(token, module, data)
	return
}

func RemoveSavedData(module string, token string) ([]interface{}, bool, error) {
	callName := fmt.Sprintf("storage.%s:remove", module)
	return getterCallSavedData(token, callName)
}

func GetSavedData(module string, token string) ([]interface{}, bool, error) {
	callName := fmt.Sprintf("storage.%s:select", module)
	return getterCallSavedData(token, callName)
}

func getterCallSavedData(token string, method string) ([]interface{}, bool, error) {
	resp, err := callFunction(method, []interface{}{token})
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

	// token, time, ...
	if len(tuple) <= 2 {
		return nil, false, fmt.Errorf("tuple error, len(tuple) <= 2")
	}

	return tuple[2:], true, nil
}