package server

type Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseClient struct {
	ClientID   string `json:"client_id"`
	ClientName string `json:"client_name"`
}

type ResponseError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
