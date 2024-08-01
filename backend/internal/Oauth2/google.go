package oauth2

import (
	"Astral/internal/jwt/auth"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var google_config *oauth2.Config

type googleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"`
}

func getGoogleOauthURL() (*oauth2.Config, string) {
	google_config = &oauth2.Config{
		RedirectURL:  os.Getenv("CLIENT_CALLBACK_URL_GOOGLE"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	state := GenerateState()
	return google_config, state
}

func GoogleOauthLogin(ctx *gin.Context) {
	config, state := getGoogleOauthURL()
	redirectURL := config.AuthCodeURL(state)

	session := sessions.Default(ctx)
	session.Set("state", state)
	err := session.Save()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, redirectURL)
}

func GoogleCallBack(ctx *gin.Context) {
	session := sessions.Default(ctx)
	state := session.Get("state")
	if state != ctx.Query("state") {
		_ = ctx.AbortWithError(http.StatusUnauthorized, ErrorState)
		return
	}

	// Код авторизации Google (ресурс) AccessToken
	code := ctx.Query("code")
	token, err := google_config.Exchange(ctx, code)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Токен доступа Google
	client := google_config.Client(context.TODO(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer userInfo.Body.Close()

	info, err := ioutil.ReadAll(userInfo.Body)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user googleUser
	err = json.Unmarshal(info, &user)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	userGG := GGUser{
		Name:  user.Name,
		Email: user.Email,
	}
	err = CheckUser(userGG)
	if err != nil {
		ctx.AbortWithError(404, err)
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	signedRefreshToken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error signing new refresh token",
		})
		ctx.Abort()
		return
	}

	// Установите новый токен обновления в файле cookie.
	ctx.SetCookie(
		"refresh_token",
		signedRefreshToken,
		60*60*24*30,
		"/",
		"localhost",
		false,
		true,
	)

	url := os.Getenv("FRONT_END_URL")

	ctx.Redirect(http.StatusTemporaryRedirect, url+"/workspaces")
}
