// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package repository

import (
	"net/http"
	"sync"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/contract"
	"github.com/webpkg/api/model"
	"github.com/webpkg/web"
)

var (
	_dataRepository     contract.DataRepository
	_onceDataRepository sync.Once
)

// CreateDataRepository return contract.AuthRepository
func CreateDataRepository() contract.DataRepository {

	_onceDataRepository.Do(func() {
		_dataRepository = &DataRepository{
			client: &http.Client{},
		}
	})

	return _dataRepository
}

// DataRepository struct
type DataRepository struct {
	client *http.Client
}

// GetAuthByAccessToken get auth by accessToken
func (r *DataRepository) GetAuthByAccessToken(accessToken string) (*model.Auth, error) {

	var err error

	auth := model.CreateAuth()

	err = web.Get(r.client, config.AuthUrl(), accessToken, auth)

	if err != nil {
		return nil, err
	}

	return auth, nil
}
