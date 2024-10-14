package gogs

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Create server globals from env file
func CreateServerObject() Server {
	var envErr = godotenv.Load("../.env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	var ogsCreds = Credentials{
		os.Getenv("OGS_CLIENT_ID"),
		os.Getenv("OGS_CLIENT_SECRET"),
		os.Getenv("OGS_USERNAME"),
		os.Getenv("OGS_PASSWORD"),
		"", "", "", "",
		RestCredentials{},
	}
	var server = Server{false, "v1", "https://online-go.com", nil, &ogsCreds}
	return server
}

var ogsServer = CreateServerObject()

func TestConnectNoAuth(t *testing.T) {
	var ogsServer = Server{false, "v1", "https://online-go.com", nil, &Credentials{}}
	err := ogsServer.Connect()
	assert.Nil(t, err, "Err should be nil")

	assert.NotNil(t, ogsServer.httpClient, "httpClient shouldnt be nil")
}
func TestConnectNoAuthBadApiVer(t *testing.T) {
	var ogsServer = Server{false, "v2", "https://online-go.com", nil, &Credentials{}}
	err := ogsServer.Connect()
	assert.NotNil(t, err, "Err shouldnt be nil")

	assert.Nil(t, ogsServer.httpClient, "httpClient should be nil")
}
func TestConnectNoAuthBadUrl(t *testing.T) {
	var ogsServer = Server{false, "v1", "https://onlin-go.com", nil, &Credentials{}}
	err := ogsServer.Connect()
	assert.NotNil(t, err, "Err shouldnt be nil")

	assert.Nil(t, ogsServer.httpClient, "httpClient should be nil")
}

func TestConnectAuth(t *testing.T) {
	err := ogsServer.Connect()
	assert.Nil(t, err, "Err should be nil")

	assert.NotNil(t, ogsServer.httpClient, "httpClient shouldnt be nil")
	assert.NotEmpty(t, ogsServer.Credentials.AccessToken, "AccessToken shouldnt be empty")
	assert.NotEmpty(t, ogsServer.Credentials.RefreshToken, "RefreshToken shouldnt be empty")
}
func TestConnectAuthBadCreds(t *testing.T) {
	var ogsCredsBad = Credentials{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", RestCredentials{}}
	var ogsServerBad = Server{false, "v1", "https://online-go.com", nil, &ogsCredsBad}

	err := ogsServerBad.Connect()
	assert.NotNil(t, err, "Err shouldnt be nil")
	assert.Nil(t, ogsServerBad.httpClient, "httpClient should be nil")
	assert.Empty(t, ogsServerBad.Credentials.AccessToken, "AccessToken should be empty")
	assert.Empty(t, ogsServerBad.Credentials.RefreshToken, "RefreshToken should be empty")

}

func TestNewServerConstructorNoAuth(t *testing.T) {
	var ogsNewServer, err = New("", "", "", "", "v1", "https://online-go.com")
	assert.Nil(t, err, "Err should be nil")

	err = ogsNewServer.Connect()
	assert.Nil(t, err, "Err should be nil")
	// Test we are unauthenticated
	assert.Empty(t, ogsNewServer.Credentials.AccessToken, "AccessToken should be empty")
}
func TestNewServerConstructorBadUrl(t *testing.T) {
	var ogsNewServer, err = New("", "", "", "", "v1", "https://onlin-go.com")
	assert.Nil(t, err, "Err should be nil")

	err = ogsNewServer.Connect()
	assert.NotNil(t, err, "Err shouldnt be nil")
	// Test we are unauthenticated
	assert.Empty(t, ogsNewServer.Credentials.AccessToken, "AccessToken should be empty")
}
func TestNewServerConstructorBadApiVer(t *testing.T) {
	var ogsNewServer, err = New("", "", "", "", "v4", "https://onlin-go.com")
	assert.Nil(t, err, "Err should be nil")

	err = ogsNewServer.Connect()
	assert.NotNil(t, err, "Err shouldnt be nil")
	// Test we are unauthenticated
	assert.Empty(t, ogsNewServer.Credentials.AccessToken, "AccessToken should be empty")
}
func TestNewServerConstructorAuth(t *testing.T) {
	clientId := os.Getenv("OGS_CLIENT_ID")
	clientSecret := os.Getenv("OGS_CLIENT_SECRET")
	username := os.Getenv("OGS_USERNAME")
	password := os.Getenv("OGS_PASSWORD")

	var ogsNewServer, err = New(clientId, clientSecret, username, password, "v1", "https://online-go.com")
	assert.Nil(t, err, "Err should be nil")

	err = ogsNewServer.Connect()
	assert.Nil(t, err, "Err should be nil")

	assert.NotEmpty(t, ogsNewServer.Credentials.AccessToken, "AccessToken shouldnt be empty")
}

func TestApiRequest(t *testing.T) {
	// Call /me/games
	result, err := ogsServer.NewAPIRequest("GET", "/me/games", nil)
	assert.Nil(t, err, "Err should be nil")
	assert.NotEmpty(t, result.Results, "Results interface shouldnt be empty")
	assert.Equal(t, result.Code, 200, "Should receive a 200")

}
func TestBadRequest(t *testing.T) {
	// Call /me/games
	result, err := ogsServer.NewAPIRequest("GET", "/bad/request", nil)
	assert.NotNil(t, err, "Err shouldnt be nil")
	assert.Empty(t, result.Results, "Results interface should be empty")
	assert.NotEqual(t, result.Code, 200, "Shouldnt receive a 200")

}
