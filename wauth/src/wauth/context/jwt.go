package context

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/inwady/easyconfig"
	"time"
)

var (
	hmacSecret = []byte(easyconfig.GetString("jwt.key", "secret"))
	expires = easyconfig.GetInt64("jwt.expires", 300)
)

func (actx *AContext) GenerateProjectToken(project string, claims jwt.MapClaims) (string, error) {
	return actx.generateSigningToken(project, claims)
}

func (actx *AContext) GenerateDefaultClaims(project string) jwt.MapClaims {
	return jwt.MapClaims{
		"project": project,
		"email": actx.GetEmail(),
		"username": actx.GetUsername(),
		"iat": time.Now().Unix(),
	}
}

func (actx *AContext) generateSigningToken(project string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
