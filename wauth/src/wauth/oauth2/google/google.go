package google

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"github.com/inwady/easyconfig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"wauth/context"
	"wauth/model/user"
)

type Oauth2GoogleData struct {
	Email string		`json:"email"`
	Name string			`json:"given_name"`
}

var (
	authRedirectSuccess = easyconfig.GetString("auth_redirect_success", "https://rhack.ru/dashboard")

	scheme = easyconfig.GetString("scheme", "")
	host = easyconfig.GetString("host", "")
)

var (
	clientId = "718385126001-dgaagqt9vubv8jt215mrgr1rs3i0mfu2.apps.googleusercontent.com"
	clientSecret = "h1vsda1FslzWo9sbDoocUVAM"

	scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
	}

	conf = &oauth2.Config{
		ClientID: clientId,
		ClientSecret: clientSecret,
		RedirectURL:  fmt.Sprintf("%s://%s/oauth2/google_callback", scheme, host),
		Scopes: scopes,
		Endpoint: google.Endpoint,
	}
)

func Oauth2Google(actx *context.AContext) {
	if actx.GetMethod() != http.MethodGet {
		actx.WriteBadMethod()
		return
	}

	authUrl := conf.AuthCodeURL("state")
	actx.RedirectTemporary(authUrl)
}

func Oauth2GoogleCallback(actx *context.AContext) {
	log.Println("SEE !!!", actx.GetMethod())
	log.Println(actx.GetRequestURI())

	state := actx.GetParamRequired("state")

	// check state
	_ = state

	code := actx.GetParamRequired("code")
	token, err := conf.Exchange(context2.Background(), code)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	client := conf.Client(context2.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		actx.WriteError500(err)
		return
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	log.Println("Resp body: ", string(data))

	oauth2Data := &Oauth2GoogleData{}
	err = json.Unmarshal(data, oauth2Data)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	u, ok, err := user.GetUserByEmail(oauth2Data.Email)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	if !ok {
		// no that user
		err := actx.MakeRegister(&context.RegisterData{
			Username: oauth2Data.Name,
			Email: oauth2Data.Email,
		})

		if err != nil {
			actx.WriteError500(err)
			return
		}
	}

	err = actx.MakeAuth(u.Email, u.Username)
	if err != nil {
		actx.WriteError500(err)
		return
	}

	actx.RedirectTemporary(authRedirectSuccess)
}
