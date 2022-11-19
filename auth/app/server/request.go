package server

type RequestClient struct {
	ClientName  string `json:"client_name"`
	RedirectURI string `json:"redirect_uri"`
	Scope       string `json:"scope"`
}

type RequestAuthorize struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}
