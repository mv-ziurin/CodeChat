package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/inwady/easyconfig"
	"net/http"
	"wauth/context"
)

var (
	projectsArray = easyconfig.GetArrayString("jwt.projects", []string{})
	projectsMap map[string]bool
)

func init() {
	// TODO in database
	projectsMap = make(map[string]bool)
	for _, project := range projectsArray {
		projectsMap[project] = true
	}
}

func GetJWTRequest(actx *context.AContext) {
	if actx.GetMethod() != http.MethodGet {
		actx.WriteBadMethod()
		return
	}

	ok, err := actx.SetSession()
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !ok {
		actx.WriteErrorCustom(context.ForbiddenStatus, "")
		return
	}

	// TODO check enable (true / false)
	project := actx.GetParamRequired("project")
	_, ok = projectsMap[project]
	if !ok {
		actx.WriteError("cannot find project")
		return
	}

	defaultClaims := actx.GenerateDefaultClaims(project)
	err = _foolBusinessLogic(actx, project, defaultClaims)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	jwtToken, err := actx.GenerateProjectToken(project, defaultClaims)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.WriteOk(jwtToken)
}

func _foolBusinessLogic(actx *context.AContext, project string, claims jwt.MapClaims) error {
	return nil
}