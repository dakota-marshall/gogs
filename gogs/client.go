package gogs

import (
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

// func New(clientId string, clientSecret string, username string, password string, apiVersion string, baseUrl string) (*Server, error) {
// 	return &Server{
// 		&Credentials{
// 			clientId,
// 			clientSecret,
// 			username,
// 			password,
// 		},
// 		false,
// 		apiVersion,
// 		baseUrl,
// 		nil,
// 	}, nil
// }

// func (server *Server) Config(clientId string, clientSecret string, username string, password string, apiVersion string, baseUrl, string) (*Server, error) {
// 	// TODO : Add code to verify parameters
// 	return &Server{
//         &Credentials{
//             clientId,
//             clientSecret,
//             username,
//             password,
//         },
//         false,
//         apiVersion,
//         baseUrl,
//         nil,
//     }, nil
// }

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
	if server.Credentials != nil {
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
// func (server *Server) NewAPIRequest(method, APICall string, jsonString []byte) (*APIResult, error) {
//
// 	fullURL := server.BaseURL + APICall
//
// 	t := &http.Transport{
// 		TLSClientConfig: &tls.Config{
// 			InsecureSkipVerify: server.AllowUnverifiedSSL,
// 		},
// 	}
//
// 	server.httpClient = &http.Client{
// 		Transport: t,
// 		Timeout:   time.Second * 60,
// 	}
//
// 	request, requestErr := http.NewRequest(method, fullURL, bytes.NewBuffer(jsonString))
// 	if requestErr != nil {
// 		return nil, requestErr
// 	}
//
// 	request.SetBasicAuth(server.Username, server.Password)
// 	request.Header.Set("Accept", "application/json")
// 	request.Header.Set("Content-Type", "application/json")
//
// 	var response *http.Response
// 	var doErr error
// 	retries := 0
// 	for {
// 		response, doErr = server.httpClient.Do(request)
//
// 		if !((doErr != nil) || (response == nil || response.StatusCode == 503)) {
// 			break
// 		}
//
// 		if retries >= server.Retries {
// 			break
// 		}
// 		retries++
// 		time.Sleep(server.RetryDelay)
// 	}
//
// 	if doErr != nil {
// 		results := APIResult{
// 			Code:        0,
// 			Status:      "Error : Request to server failed : " + doErr.Error(),
// 			ErrorString: doErr.Error(),
// 			Retries:     retries,
// 		}
// 		return &results, doErr
// 	}
// 	defer response.Body.Close()
//
// 	var results APIResult
// 	if decodeErr := json.NewDecoder(response.Body).Decode(&results); decodeErr != nil {
// 		return nil, decodeErr
// 	}
//
// 	if results.Retries == 0 { // results.Retries have default value so set it.
// 		results.Retries = retries
// 	}
//
// 	if results.Code == 0 { // results.Code has default value so set it.
// 		results.Code = response.StatusCode
// 	}
//
// 	if results.Status == "" { // results.Status has default value, so set it.
// 		results.Status = response.Status
// 	}
//
// 	switch results.Code {
// 	case 0:
// 		results.ErrorString = "Did not get a response code."
// 	case 404:
// 		results.ErrorString = results.Status
// 	case 200:
// 		results.ErrorString = results.Status
// 	default:
// 		results.ErrorString = results.Status
// 		//theError := strings.Replace(results.Results.([]interface{})[0].(map[string]interface{})["errors"].([]interface{})[0].(string), "\n", " ", -1)
// 		//results.ErrorString = strings.Replace(theError, "Error: ", "", -1)
//
// 	}
//
// 	return &results, nil
//
// }
