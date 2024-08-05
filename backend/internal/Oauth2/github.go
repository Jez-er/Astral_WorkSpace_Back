package oauth2

import (
	"Astral/internal/jwt/auth"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubConfig *oauth2.Config

type githubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUserName   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type githubEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

func getPrimaryEmail(client *http.Client) (string, error) {
	emailInfo, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return "", err
	}
	defer emailInfo.Body.Close()

	var emails []githubEmail
	err = json.NewDecoder(emailInfo.Body).Decode(&emails)
	if err != nil {
		return "", err
	}

	for _, email := range emails {
		if email.Primary {
			return email.Email, nil
		}
	}
	return "", fmt.Errorf("no primary email found")
}

func getGithubOauthURL() (*oauth2.Config, string) {
	githubConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("CLIENT_CALLBACK_URL_GITHUB"),
		ClientID:     os.Getenv("CLIENT_ID_GIT"),
		ClientSecret: os.Getenv("CLIENT_SECRET_GIT"),
		Scopes: []string{
			"user",
			"repo",
			"user:email",
		},
		Endpoint: github.Endpoint,
	}

	state := GenerateState()
	return githubConfig, state
}

func GithubOauthLogin(ctx *gin.Context) {
	config, state := getGithubOauthURL()
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

func GithubCallBack(ctx *gin.Context) {
	session := sessions.Default(ctx)
	state := session.Get("state")
	if state != ctx.Query("state") {
		_ = ctx.AbortWithError(http.StatusUnauthorized, ErrorState)
		return
	}

	code := ctx.Query("code")
	token, err := githubConfig.Exchange(ctx, code)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := githubConfig.Client(context.TODO(), token)
	userInfo, err := client.Get("https://api.github.com/user")
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

	var user githubUser
	err = json.Unmarshal(info, &user)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	primaryEmail, err := getPrimaryEmail(client)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Email = primaryEmail

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	// userGG := GGUser{
	// 	Name:  user.Login,
	// 	Email: user.Email,
	// 	DisplayName: user.Name,
	// }
	// err = CheckUser(userGG)
	// if err != nil {
	// 	ctx.AbortWithError(404, err)
	// }
	
	signedRefreshToken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error signing new refresh token",
		})
		ctx.Abort()
		return
	}

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
