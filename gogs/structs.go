package gogs

import (
	"net/http"
)

// Struct for handling the data received from the API endpoint
type ApiResult struct {
	Code        int
	Status      string
	Error       string
	ErrorString string
	Results     interface{}
}

// Struct for storing the credentials we receive from the REST API
type RestCredentials struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// Contains all the relevant credentials for both the REST and Socket API's
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

// The primary server struct that contains all state related to the connection to OGS
type Server struct {
	IsAuthed    bool
	ApiVersion  string
	BaseUrl     string
	httpClient  *http.Client
	Credentials *Credentials
}
