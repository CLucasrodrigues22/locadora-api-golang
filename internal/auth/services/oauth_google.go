package services

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	logger *logs.Logger
)

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     utils.GetEnv("CLIENT_ID", logger),
	ClientSecret: utils.GetEnv("CLIENT_SECRET", logger),
	RedirectURL:  utils.GetEnv("REDIRECT_URL", logger),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

var OAuthStateString = utils.GetEnv("OAUTH_STATE_STRING", logger)

func GetGoogleLoginURL() string {
	return GoogleOAuthConfig.AuthCodeURL(OAuthStateString)
}
