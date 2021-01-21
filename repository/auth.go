package repository

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/webpkg/api/contract"
	"github.com/webpkg/api/helper"
	"github.com/webpkg/api/model"
)

var (
	_authRepository     contract.AuthRepository
	_onceAuthRepository sync.Once
)

// CreateAuthRepository return contract.AuthRepository
func CreateAuthRepository() contract.AuthRepository {

	_onceAuthRepository.Do(func() {
		_authRepository = &AuthRepository{
			client: &http.Client{},
		}
	})

	return _authRepository
}

// AuthRepository struct
type AuthRepository struct {
	client *http.Client
}

// GetAuthByAccessToken get auth by accessToken
func (r *AuthRepository) GetAuthByAccessToken(accessToken string) (*model.Auth, error) {
	var err error

	val := model.CreateAuth()

	ac := WebConfig().Auth

	for _, addr := range ac.Addr {

		uri := fmt.Sprintf("%s/client/auth/%s", addr, ac.AppKey)

		err = helper.HTTPGet(r.client, uri, accessToken, val)

		if err != nil {
			if err == helper.ErrUnExpectedError {
				continue
			}

			return nil, err
		}

		return val, nil
	}

	return nil, err
}
