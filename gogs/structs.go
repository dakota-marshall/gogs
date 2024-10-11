package gogs

type APIResult struct {
	Error       float64 `json:"error"`
	ErrorString string
	Status      string      `json:"Status"`
	Code        int         `json:"Code"`
	Results     interface{} `json:"results"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}
