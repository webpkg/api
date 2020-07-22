package proxy

import (
	"github.com/webpkg/api/model"
	"github.com/webpkg/api/repository"
)

// GetAuthByAccessToken get auth by accessToken
func GetAuthByAccessToken(accessToken string) (*model.Auth, error) {
	repo := repository.CreateAuthRepository()
	return repo.GetAuthByAccessToken(accessToken)
}
