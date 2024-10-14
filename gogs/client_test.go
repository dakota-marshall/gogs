package gogs

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
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
	var ogsServer = Server{false, "v1", "https://online-go.com", nil, nil}
	err := ogsServer.Connect()
	if err != nil {
		t.Error(err)
	}
	if ogsServer.httpClient == nil {
		t.Errorf("Failed to successfully connect to OGS")
	}
}
func TestConnectNoAuthBadApiVer(t *testing.T) {
	var ogsServer = Server{false, "v2", "https://online-go.com", nil, nil}
	err := ogsServer.Connect()
	if err == nil {
		t.Errorf("No error received on bad API type")
	}
	if ogsServer.httpClient != nil {
		t.Errorf("server.httpClient not nil")
	}
}
func TestConnectNoAuthBadUrl(t *testing.T) {
	var ogsServer = Server{false, "v1", "https://onlin-go.com", nil, nil}
	err := ogsServer.Connect()
	if err == nil {
		t.Errorf("No error received on bad API type")
	}
	if ogsServer.httpClient != nil {
		t.Errorf("server.httpClient not nil")
	}
}

func TestConnectAuth(t *testing.T) {
	err := ogsServer.Connect()
	if err != nil {
		t.Error(err)
	}
	if ogsServer.httpClient == nil {
		t.Errorf("Failed to successfully connect to OGS")
	}
	if ogsServer.Credentials.AccessToken == "" {
		t.Errorf("Failed to get AccessToken from OGS")
	}
	if ogsServer.Credentials.RefreshToken == "" {
		t.Errorf("Failed to get RefreshToken from OGS")
	}
}
func TestConnectAuthBadCreds(t *testing.T) {
	var ogsCredsBad = Credentials{"", "", "", "", "", "", "", "", RestCredentials{}}
	var ogsServerBad = Server{false, "v1", "https://online-go.com", nil, &ogsCredsBad}

	err := ogsServerBad.Connect()
	if err == nil {
		t.Errorf("Failed auth error is nil")
	}
	if ogsServerBad.httpClient != nil {
		t.Errorf("server.httpClient is not nil")
	}
	if ogsServerBad.Credentials.AccessToken != "" {
		t.Errorf("AccessToken is not empty: " + ogsServerBad.Credentials.AccessToken)
	}
	if ogsServerBad.Credentials.RefreshToken != "" {
		t.Errorf("RefreshToken is not empty: " + ogsServerBad.Credentials.RefreshToken)
	}
}
