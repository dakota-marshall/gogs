package gogs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (creds *Credentials) UpdateRestCreds(src *RestCredentials) {
	if src.AccessToken != "" {
		creds.AccessToken = src.AccessToken
	}
	if src.TokenType != "" {
		creds.TokenType = src.TokenType
	}
	if src.ExpiresIn != 0 {
		creds.ExpiresIn = src.ExpiresIn
	}
	if src.Scope != "" {
		creds.Scope = src.Scope
	}
	if src.RefreshToken != "" {
		creds.RefreshToken = src.RefreshToken
	}
}

func New(clientId string, clientSecret string, username string, password string, apiVersion string, baseUrl string) (*Server, error) {
	creds := []string{clientId, clientSecret, username, password}
	missingCred := false
	for _, cred := range creds {
		if cred == "" {
			missingCred = true
			break
		}
	}

	if missingCred {
		newServer := Server{false, apiVersion, baseUrl, nil, &Credentials{}}
		return &newServer, nil
	} else {
		newCreds := Credentials{clientId, clientSecret, username, password, "", "", "", "", RestCredentials{}}
		newServer := Server{false, apiVersion, baseUrl, nil, &newCreds}
		return &newServer, nil
	}
}

func (server *Server) Connect() error {

	// Setup Client
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	server.httpClient = &http.Client{
		Transport: t,
		Timeout:   time.Second * 60,
	}

	var response *http.Response
	var err error

	// get token if auth provided, otherwise just test connection
	var emptyCreds Credentials
	if *server.Credentials != emptyCreds {
		data := url.Values{
			"client_id":  {server.Credentials.ClientID},
			"grant_type": {"password"},
			"username":   {server.Credentials.Username},
			"password":   {server.Credentials.Password},
		}

		response, err = server.httpClient.PostForm(server.BaseUrl+"/oauth2/token/", data)
		if err != nil {
			server.httpClient = nil
			return err
		}
		if response.StatusCode != 200 {
			server.httpClient = nil
			return fmt.Errorf("Error: Got non 200 status code from auth-endpoint: %d", response.StatusCode)
		}

		// Decode response
		var authData RestCredentials
		if decodeErr := json.NewDecoder(response.Body).Decode(&authData); decodeErr != nil {
			return decodeErr
		}

		// Update fields in server credentials
		server.Credentials.UpdateRestCreds(&authData)
		defer response.Body.Close()
	} else { // Call API root to ensure we can talk to OGS Api
		request, err := http.NewRequest("GET", server.BaseUrl+"/api/"+server.ApiVersion, nil)
		if err != nil {
			server.httpClient = nil
			return err
		}

		request.Header.Set("Accept", "application/json")
		response, err := server.httpClient.Do(request)
		if err != nil || response == nil {
			server.httpClient = nil
			return err
		}
		if response.StatusCode != 200 {
			server.httpClient = nil
			return fmt.Errorf("Non 200 test response: %d", response.StatusCode)
		}
		defer response.Body.Close()
	}

	return nil
}

// NewAPIRequest ...
func (server *Server) NewAPIRequest(method, apiCall string, jsonString []byte) (*ApiResult, error) {
	fullUrl := server.BaseUrl + "/api/" + server.ApiVersion + apiCall

	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	server.httpClient = &http.Client{
		Transport: t,
		Timeout:   time.Second * 60,
	}

	request, requestErr := http.NewRequest(method, fullUrl, bytes.NewBuffer(jsonString))
	if requestErr != nil {
		return nil, requestErr
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+server.Credentials.AccessToken)

	var response *http.Response
	var doErr error
	var result ApiResult

	response, doErr = server.httpClient.Do(request)
	if doErr != nil {
		result.Code = 0
		result.Status = "Error: Request to server failed: " + doErr.Error()
		result.ErrorString = doErr.Error()
		return &result, doErr
	}
	defer response.Body.Close()

	result.Status = response.Status

	// Check Return Codes
	result.Code = response.StatusCode
	switch result.Code {
	case 0:
		result.ErrorString = "Did not get response code"
	default:
		result.ErrorString = result.Status
	}

	// Decode Response Body
	decodeErr := json.NewDecoder(response.Body).Decode(&result.Results)
	if decodeErr != nil {
		result.Code = 0
		result.Status = "Error: failed to decode response body: " + decodeErr.Error()
		result.ErrorString = decodeErr.Error()
		return &result, decodeErr
	}
	return &result, nil

}
