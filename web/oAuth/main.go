package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/pat"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://127.0.0.1:3000/auth/google/callback",
	ClientID:     loadEnvVal("GOOGLE_CLIENT_ID"),
	ClientSecret: loadEnvVal("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, cookie) //http에 cookie 등록

	return state
}

func googleAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	oauthstate, err := r.Cookie("oauthstate")
	if err != nil {
		log.Printf("oauthstate is not Founded%s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	if r.FormValue("state") != oauthstate.Value {
		log.Printf("invalid google oauth state cookie: %s state:%s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Printf("Failed to load code value %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	fmt.Fprint(w, string(data))
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}
	res, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get User Info %s\n", err.Error())
	}

	return ioutil.ReadAll(res.Body)
}
func loadEnvVal(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func main() {
	mux := pat.New()

	mux.HandleFunc("/auth/google/login", googleLoginHandler)
	mux.HandleFunc("/auth/google/callback", googleAuthCallbackHandler)
	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
