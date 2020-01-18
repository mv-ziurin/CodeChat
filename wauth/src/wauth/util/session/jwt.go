package session

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/inwady/easyconfig"
	"log"
)

var (
	hmacSecret = []byte(easyconfig.GetString("jwt.key", "secret"))
	expires = easyconfig.GetInt64("jwt.expires", 300)
)

func isExpiredToken(createdTime int64) bool {
	return time.Now().Unix() >= createdTime + expires
}

func AuthJWTToken(projectName string, jwtToken string) (jwt.MapClaims, bool, string) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// TODO check valid
		return hmacSecret, nil
	})

	if err != nil {
		log.Println(err)
		return nil, false, "parse error"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, false, "parse error"
	}


	iat, ok := getIat(claims)
	if !ok {
		return nil, false, "bad jwt format"
	}

	if isExpiredToken(iat) {
		return nil, false, "cookie was expired"
	}

	projectNameJwt, ok := GetStringFromClaims(claims, "project")
	if !ok {
		return nil, false, "bad jwt format"
	}

	if projectName != projectNameJwt {
		return nil, false, "bad project"
	}

	return claims, true, ""
}

func GetStringFromClaims(claims jwt.MapClaims, key string) (res string, ok bool) {
	res, ok = claims[key].(string)
	if !ok {
		return
	}

	ok = true
	return
}

func getIat(claims jwt.MapClaims) (iat int64, ok bool) {
	iatFloat, ok := claims["iat"].(float64)
	if !ok {
		return
	}

	iat = int64(iatFloat)
	ok = true
	return
}
