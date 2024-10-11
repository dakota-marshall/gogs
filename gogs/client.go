package gogs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Credentials struct {
	ClientID         string
	ClientSecret     string
	Username         string
	Password         string
	AccessToken      string
	RefreshToken     string
	UserID           string
	ChatAuth         string
	UserJwt          string
	NotificationAuth string
}

type Server struct {
	Credentials *Credentials
	IsAuthed    bool
	ApiVersion  string
	BaseUrl     string
	httpClient  *http.Client
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

	// get token if auth provided, otherwise nothing to do
	if server.Credentials != nil {
		data := url.Values{
			"client_id":  {server.Credentials.ClientID},
			"grant_type": {"password"},
			"username":   {server.Credentials.Username},
			"password":   {server.Credentials.Password},
		}

		response, err = server.httpClient.PostForm(server.BaseUrl+"/oauth2/token", data)
		if err != nil {
			server.httpClient = nil
			return err
		}

		// Decode response
		var results APIResult
		if decodeErr := json.NewDecoder(response.Body).Decode(&results); decodeErr != nil {
			return decodeErr
		}
		if results.Code != 200 {
			return fmt.Errorf("Got bad return code from auth endpoint")
		}

		var authData AuthResponse
		if decodeErr := json.NewDecoder(results.Results).Decode(&authData); decodeErr != nil {
			return decodeErr
		}

	}

	defer response.Body.Close()

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
