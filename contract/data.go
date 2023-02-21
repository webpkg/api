// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package contract

import "github.com/gostartkit/api/model"

// DataRepository interface
type DataRepository interface {
	// GetAuthByAccessToken get auth by accessToken
	GetAuthByAccessToken(accessToken string) (*model.Auth, error)
}
