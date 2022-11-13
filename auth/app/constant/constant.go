package constant

type ResponseTypes int

const (
	Code ResponseTypes = iota
)

func (r ResponseTypes) String() string {
	switch r {
	case Code:
		return "code"
	}
	return ""
}

type GrantTypes int

const (
	AuthCode GrantTypes = iota
	RefreshToken
)

func (g GrantTypes) String() string {
	switch g {
	case RefreshToken:
		return "refresh_token"
	case AuthCode:
		return "authorization_code"
	}
	return ""
}

type ClientTypes int

const (
	Public ClientTypes = iota
)

func (c ClientTypes) String() string {
	switch c {
	case Public:
		return "public"
	}
	return ""
}

type QueryParamsAuthorization int

const (
	ResponseType QueryParamsAuthorization = iota
	ClientID
	Scope
	RedirectURI
	State
)

func (q QueryParamsAuthorization) String() string {
	switch q {
	case ResponseType:
		return "response_type"
	case ClientID:
		return "client_id"
	case Scope:
		return "scope"
	case RedirectURI:
		return "redirect_uri"
	case State:
		return "state"
	}
	return ""
}
