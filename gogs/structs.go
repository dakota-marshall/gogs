package gogs

import (
	"net/http"
)

type RestCredentials struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type Credentials struct {
	ClientID         string
	ClientSecret     string
	Username         string
	Password         string
	UserID           string
	ChatAuth         string
	UserJwt          string
	NotificationAuth string
	RestCredentials
}

type Server struct {
	IsAuthed    bool
	ApiVersion  string
	BaseUrl     string
	httpClient  *http.Client
	Credentials *Credentials
}
