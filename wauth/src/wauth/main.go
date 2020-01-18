package main

import (
	"wauth/api"
	"github.com/inwady/easyconfig"
	"github.com/valyala/fasthttp"
	"wauth/context"
	"net/http"
	"wauth/database"
	"log"
	"strconv"
	"net/url"
	"fmt"
)

var (
	port = easyconfig.GetInt("bind", 7070)

	accessOrigins = easyconfig.GetArrayString("access_origins", nil)
	mapAccessOrigins = map[string]bool{}
)

func init() {
	/* default map value ~ false */
	for _, origins := range accessOrigins {
		mapAccessOrigins[origins] = true
	}
}

func setSpecialOrigins(actx *context.AContext) {
	actx.SetHeader("Vary", "Origin")

	actx.SetHeader("Access-Control-Allow-Methods", "OPTIONS, GET, HEAD, POST, PUT")
	actx.SetHeader("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, withCredentials")
}

func handleRequest(ctx *fasthttp.RequestCtx) {
	actx, _, err := context.InitFromHTTP(ctx)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	defer func() {
		actx.Exit(recover())
	}()

	actx.LogDebug("MAIN", actx.GetPath())

	origin, ok := actx.GetHeader("Origin")
	if ok {
		originUrl, err := url.Parse(origin)
		if err == nil {
			originHost := fmt.Sprintf("%s://%s", originUrl.Scheme, originUrl.Host)
			log.Println("Origin: ", originHost)
			if mapAccessOrigins[originHost] == true {
				actx.SetHeader("Access-Control-Allow-Origin", origin)
				actx.SetHeader("Access-Control-Allow-Credentials", "true")

				if actx.GetMethod() == http.MethodOptions {
					setSpecialOrigins(actx)
					return
				}
			}
		}
	}

	switch string(ctx.Path()) {
	case "/login":
		api.LoginRequest(actx)
	case "/jslogin":
		api.LoginJSRequest(actx)
	case "/request":
		api.RegisterRequest(actx)
	case "/register":
		api.RegisterFinal(actx)
	case "/logout":
		api.LogoutRequest(actx)
	case "/change":
		api.ChangeDataRequest(actx)
	case "/restore":
		api.MainRestoreRequest(actx)
	case "/jsrestore":
		api.RestoreAfterEmailRequest(actx)
	case "/x":
		api.ReceiveEmailRequest(actx)
	case "/userinfo":
		api.GetConfigRequest(actx)

	// jwt get
	case "/jwt":
		api.GetJWTRequest(actx)

	/* oauth2 exit */
	default:
		actx.WriteErrorNotFound()
	}
}

func main() {
	defer database.DataBaseClose()

	chans := make(chan error, 1)
	go func() {
		chans <- fasthttp.ListenAndServe(":" + strconv.Itoa(port), handleRequest)
	}()

	log.Println("start listen, port", strconv.Itoa(port))

	select {
	case err := <-chans:
		log.Fatal(err)
	}
}
