// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package proxy

import (
	"github.com/webpkg/api/model"
	"github.com/webpkg/api/repository"
)

var (
	dataRepository = repository.CreateDataRepository()
)

// GetAuthByAccessToken get auth by accessToken
func GetAuthByAccessToken(accessToken string) (*model.Auth, error) {
	return dataRepository.GetAuthByAccessToken(accessToken)
}
