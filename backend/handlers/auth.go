package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var cred Credentials
var conf *oauth2.Config

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func init() {
	var c Credentials
	file, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		fmt.Printf("Error reading creds.json: %v", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &c)
	conf = &oauth2.Config{
		ClientID:     c.Cid,
		ClientSecret: c.Csecret,
		RedirectURL:  "http://localhost:8080/auth",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "openid"},
		Endpoint:     google.Endpoint,
	}
}

func loginHandler(c *gin.Context) {
	state := RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
}

func authHandler(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(401, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}
	code := c.Query("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.AbortWithError(500, fmt.Errorf("Code exchange failed: %s", err.Error()))
		return
	}
	client := conf.Client(oauth2.NoContext, tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo?alt=json")
	if err != nil {
		c.AbortWithError(500, fmt.Errorf("Failed getting user info: %s", err.Error()))
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(500, fmt.Errorf("Failed reading response body: %s", err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   string(data),
	})
}
