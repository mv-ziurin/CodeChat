package context

import (
	"fmt"
	"log"
)

func (actx *AContext) LogError(err error) {
	log.Printf("[email:%s] error: %v", actx._logGetEmail(), err)

	// TODO parse function which called this
}

func (actx *AContext) LogErrorText(errorMessage string) {
	log.Printf("[email:%s] error: %v", actx._logGetEmail(), errorMessage)
}

func (actx *AContext) LogDebug(module string, message string) {
	log.Printf("[%s] [email:%s] %s", module, actx._logGetEmail(), message)
}

func (actx *AContext) LogDebugf(module string, format string, v ...interface{}) {
	resultString := fmt.Sprintf(format, v)
	log.Printf("[%s] [email:%s] %s", module, actx._logGetEmail(), resultString)
}

func (actx *AContext) _logGetEmail() string {
	email := actx.GetEmailOptional()
	if email == "" {
		return "<unknown>"
	}

	return email
}
// without context
func PrintDebugf(module string, format string, v ...interface{}) {
	resultString := fmt.Sprintf(format, v)
	log.Printf("[%s] %s", module, resultString)
}
