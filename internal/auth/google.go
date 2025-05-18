package auth

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/common"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     common.GetEnv("CLIENT_ID"),
	ClientSecret: common.GetEnv("CLIENT_SECRET"),
	RedirectURL:  common.GetEnv("REDIRECT_URL"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

var OAuthStateString = common.GetEnv("OAUTH_STATE_STRING")

func GetGoogleLoginURL() string {
	return GoogleOAuthConfig.AuthCodeURL(OAuthStateString)
}
