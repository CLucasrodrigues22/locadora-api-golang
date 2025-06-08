package auth

import (
	"context"
	"encoding/json"
	"github.com/CLucasrodrigues22/api-locadora/internal/auth"
	"github.com/CLucasrodrigues22/api-locadora/internal/auth/services"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GoogleLoginHandler(ctx *gin.Context) {
	url := services.GetGoogleLoginURL()
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallbackHandler(ctx *gin.Context) {
	state := ctx.Query("state")
	if state != services.OAuthStateString {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth state"})
		return
	}

	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	token, err := services.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	client := services.GoogleOAuthConfig.Client(context.Background(), token)

	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	defer res.Body.Close()

	var user schemas.GoogleUser
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	tokenJWT, err := auth.GenerateJWT(user.Email, user.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   tokenJWT,
		"name":    user.Name,
		"email":   user.Email,
		"picture": user.Picture,
	})
}
