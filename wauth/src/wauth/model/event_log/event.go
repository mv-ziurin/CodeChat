package event_log

import (
	"wauth/context"
	"wauth/database"
	"time"
	"log"
	"fmt"
)

var (
	sqlInsertEvent = "INSERT INTO event_log (event, username, email, ip, user_agent, event_time) VALUES ($1, $2, $3, $4, $5, $6);"
)

type EventType string
var (
	LoginEvent EventType = "login"
	LogoutEvent EventType = "logout"
	RegisterEvent EventType = "register"
	RequestEvent EventType = "request"
)

type Event struct {
	Event EventType
	Username string
	Email string
	Ip string
	UserAgent string
	EventTime time.Time
}

func LogEvent(e *Event) {
	dbconn := database.GetConnection()
	_, err := dbconn.Exec(sqlInsertEvent,
		e.Event,
		returnPtr(&e.Username),
		returnPtr(&e.Email),
		e.Ip,
		e.UserAgent,
		e.EventTime)

	logEvent(e)
	if err != nil {
		log.Println("bad event log:", err)
		return
	}
}

func LogEventFromContext(event EventType, actx *context.AContext) {
	ua, _ := actx.GetHeader("User-Agent")
	e := &Event{
		Event: event,
		Email: actx.GetEmailOptional(),
		Ip: actx.GetIP(),
		UserAgent: ua,
		EventTime: time.Now(),
	}

	LogEvent(e)
}

func logEvent(e *Event) {
	var email string
	if e.Email != "" {
		email = e.Email
	} else {
		email = "{UNKNOWN EMAIL}"
	}

	log.Println("event log",
		fmt.Sprintf("{%s}", e.Event),
		email,
		e.Ip,
		e.UserAgent,
		e.EventTime)
}

func returnPtr(data *string) *string {
	if data != nil && *data == "" {
		return nil
	}

	return data
}
