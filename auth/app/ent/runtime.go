// Code generated by ent, DO NOT EDIT.

package ent

import (
	"auth/ent/authorizationcode"
	"auth/ent/oauthclient"
	"auth/ent/refreshtoken"
	"auth/ent/schema"
	"auth/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authorizationcodeFields := schema.AuthorizationCode{}.Fields()
	_ = authorizationcodeFields
	// authorizationcodeDescClientID is the schema descriptor for client_id field.
	authorizationcodeDescClientID := authorizationcodeFields[0].Descriptor()
	// authorizationcode.ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	authorizationcode.ClientIDValidator = authorizationcodeDescClientID.Validators[0].(func(string) error)
	// authorizationcodeDescCode is the schema descriptor for code field.
	authorizationcodeDescCode := authorizationcodeFields[1].Descriptor()
	// authorizationcode.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	authorizationcode.CodeValidator = authorizationcodeDescCode.Validators[0].(func(string) error)
	// authorizationcodeDescIssued is the schema descriptor for issued field.
	authorizationcodeDescIssued := authorizationcodeFields[2].Descriptor()
	// authorizationcode.DefaultIssued holds the default value on creation for the issued field.
	authorizationcode.DefaultIssued = authorizationcodeDescIssued.Default.(time.Time)
	oauthclientFields := schema.OAuthClient{}.Fields()
	_ = oauthclientFields
	// oauthclientDescClientID is the schema descriptor for client_id field.
	oauthclientDescClientID := oauthclientFields[0].Descriptor()
	// oauthclient.ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	oauthclient.ClientIDValidator = oauthclientDescClientID.Validators[0].(func(string) error)
	// oauthclientDescClientType is the schema descriptor for client_type field.
	oauthclientDescClientType := oauthclientFields[1].Descriptor()
	// oauthclient.ClientTypeValidator is a validator for the "client_type" field. It is called by the builders before save.
	oauthclient.ClientTypeValidator = oauthclientDescClientType.Validators[0].(func(string) error)
	// oauthclientDescClientName is the schema descriptor for client_name field.
	oauthclientDescClientName := oauthclientFields[2].Descriptor()
	// oauthclient.ClientNameValidator is a validator for the "client_name" field. It is called by the builders before save.
	oauthclient.ClientNameValidator = oauthclientDescClientName.Validators[0].(func(string) error)
	// oauthclientDescRedirectURI is the schema descriptor for redirect_uri field.
	oauthclientDescRedirectURI := oauthclientFields[3].Descriptor()
	// oauthclient.RedirectURIValidator is a validator for the "redirect_uri" field. It is called by the builders before save.
	oauthclient.RedirectURIValidator = oauthclientDescRedirectURI.Validators[0].(func(string) error)
	// oauthclientDescScope is the schema descriptor for scope field.
	oauthclientDescScope := oauthclientFields[4].Descriptor()
	// oauthclient.ScopeValidator is a validator for the "scope" field. It is called by the builders before save.
	oauthclient.ScopeValidator = oauthclientDescScope.Validators[0].(func(string) error)
	refreshtokenFields := schema.RefreshToken{}.Fields()
	_ = refreshtokenFields
	// refreshtokenDescToken is the schema descriptor for token field.
	refreshtokenDescToken := refreshtokenFields[0].Descriptor()
	// refreshtoken.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	refreshtoken.TokenValidator = refreshtokenDescToken.Validators[0].(func(string) error)
	// refreshtokenDescExpired is the schema descriptor for expired field.
	refreshtokenDescExpired := refreshtokenFields[1].Descriptor()
	// refreshtoken.DefaultExpired holds the default value on creation for the expired field.
	refreshtoken.DefaultExpired = refreshtokenDescExpired.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
