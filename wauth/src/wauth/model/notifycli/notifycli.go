package notifycli

import (
	"encoding/json"
	"github.com/inwady/easyconfig"
	"gitlab.com/gamechain/gameserver/wauth/src/wauth/util"
	"net/url"
	"time"
)

var (
	notifierSchemeAPI = easyconfig.GetString("notifier.scheme", "")
	notifierServerAPI = easyconfig.GetString("notifier.server", "")
	notifierTimeout   = time.Duration(easyconfig.GetInt64("notifier.timeout", 1500))
)

// type ~ text
type TextNotify struct {
	Text string				`json:"text"`
}

type DataNotify struct {
	Type string				`json:"type"`
	Data interface{}		`json:"data"`
}

type Notify struct {
	Type string				`json:"type"`
	Ts int64				`json:"ts"`
	Data interface{}		`json:"data"`
}

func NotifySend(nick string, objectMessage interface{}, saveFlag bool) error {
	text, _ := json.Marshal(Notify{
		Type: "notify",
		Ts: time.Now().Unix(),
		Data: objectMessage,
	})

	var save string
	if saveFlag {
		save = "true"
	} else {
		save = "false"
	}

	return util.SendGetRequest(notifierSchemeAPI, notifierServerAPI, "/push", url.Values{
		"nick": []string{nick},
		"message": []string{string(text)},
		"save": []string{save},
	}, notifierTimeout)
}

func NotifySendText(nick string, message string) error {
	return NotifySend(nick, &DataNotify{
		Type: "text",
		Data: &TextNotify{
			Text: message,
		},
	}, true)
}
