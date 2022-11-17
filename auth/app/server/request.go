package server

type RequestClient struct {
	ClientName  string `json:"client_name"`
	RedirectURI string `json:"redirect_uri"`
	Scope       string `json:"scope"`
}
