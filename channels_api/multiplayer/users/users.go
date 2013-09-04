package users

import (
	"errors"
	"net/http"
)

/*
	takes a refresh token and returns a user token
*/
const (
	USER          = "%s/users/_design/users/_show/user/%s?gameId=%s"
	SECURE_USER   = "%s/users/_design/users/_show/secure_user/%s?gameId=%s&token=%s"
	REFRESH_TOKEN = "%s/users/_design/users/_update/refresh_token/%s?gameId=%s&refresh_token=%s"
	NEW_TOKEN     = "%s/users/_design/users/_update/refresh_token/%s?gameId=%s"
)

func RefreshToken(c *http.Client, baseurl, userId, game, token string) {

}

/*
	generates a new token to use

*/

func NewToken() {

}

func RejectToken() {
	var s SecureUser
}

/*
	userId is well...the userId
	game is the game name like "dragon" or "tcatics"
*/
func FindUser(userId, game string) (*User, error) {
	// find and retrun user..upon return check if request is bogus or not..either way return a 401 if it fails
}

/*
	userId is well...the userId
	game is the game name like "dragon" or "tcatics"
	token is the ACTIVE token not the refresh token

*/
func FindSecuredUser(userId, game, token string) (*User, error) {
	// find and retrun user..upon return check if request is bogus or not..either way return a 401 if it fails
}

type User struct {
	Id      string `json:"id"`
	Game    string `json:"game"`
	Install int64  `json:"install_date"`
	Login   int64  `json:"last_login"`
}

type SecureUser struct {
	User
	Token        string `json:"active_token"`
	Refresh      string `json:"refresh_token"`
	Expires      int64  `json:"active_expires"`
	RefreshAgent string `json:"refresh_ua"`
	ActiveAgent  string `json:"active_ua"`
	Ip           string `json:"active_ip"`
}
