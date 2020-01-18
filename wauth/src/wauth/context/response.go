package context

import (
	"net/http"
	"encoding/json"
	"log"
)

type Status string
var (
	OkStatus Status = "OK"

	ErrorStatus Status = "ERROR"
	BadData Status = "BADDATA"
	NotFound Status = "NOTFOUND"
	ForbiddenStatus Status = "FORBIDDEN"

	ErrorUserExist Status = "ERROR_USER_EXIST"
	ErrorPasswordAlreadyWas Status = "PASSWORD_ALREADY_IN_HISTORY"
)

type ResponseBasic struct {
	Status Status 		`json:"status"`
	Data interface{}	`json:"data"`
}

type ResponseBasicError struct {
	Status Status		`json:"status"`
	Error string		`json:"error"`
}

func (actx *AContext) Redirect(uri string, statusCode int) {
	log.Printf("Redirect user to %s with code %d", uri, statusCode)
	actx.fastContext.Redirect(uri, statusCode)
}

func (actx *AContext) RedirectTemporary(uri string) {
	actx.Redirect(uri, http.StatusFound)
}

func (actx *AContext) RedirectPermanently(uri string) {
	actx.Redirect(uri, http.StatusMovedPermanently)
}

/* work with status code */

func (actx *AContext) SetStatusCode(status int) {
	actx.fastContext.SetStatusCode(status)
}

func (actx *AContext) WriteBadMethod() {
	actx.fastContext.SetStatusCode(http.StatusMethodNotAllowed)
}

/*************************/

func (actx *AContext) WriteErrorCustom(status Status, message string) {
	actx.fastContext.SetStatusCode(http.StatusOK) // TODO need status code?

	b, err := json.Marshal(&ResponseBasicError{
		Status: status,
		Error: message,
	})

	if err != nil {
		log.Panic(err)
	}

	actx.fastContext.Write(b)
}

func (actx *AContext) WriteError(message string) {
	actx.WriteErrorCustom(ErrorStatus, message)
}

func (actx *AContext) WriteErrorStatus(status Status) {
	actx.WriteErrorCustom(status, "")
}

func (actx *AContext) WriteOk(data interface{}) {
	actx.fastContext.SetStatusCode(http.StatusOK) // TODO need status code?

	b, err := json.Marshal(&ResponseBasic{
		Status: OkStatus,
		Data: data,
	})

	if err != nil {
		log.Panic(err)
	}

	actx.fastContext.Write(b)
}

func (actx *AContext) EasyWriteOk() {
	actx.WriteOk(nil)
}

func (actx *AContext) WriteErrorBadRequest() {
	actx.WriteError("bad request")
}

func (actx *AContext) WriteErrorBadRequestWithError(err error) {
	// TODO check error
	log.Println(err)
	actx.WriteErrorBadRequest()
}

func (actx *AContext) WriteErrorNotFound() {
	actx.WriteErrorCustom(NotFound,"not found")
}

func (actx *AContext) WriteErrorProcessed(err error) {
	actx.WriteError(err.Error())
}

func (actx *AContext) WriteError500(err error) {
	// TODO make error
	log.Println(err)
	actx.SetStatusCode(http.StatusInternalServerError)
}

