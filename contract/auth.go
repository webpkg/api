package contract

import "github.com/webpkg/api/model"

// AuthRepository interface
type AuthRepository interface {
	// GetAuthByAccessToken get auth by accessToken
	GetAuthByAccessToken(accessToken string) (*model.Auth, error)
}
