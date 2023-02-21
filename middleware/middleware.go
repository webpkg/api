// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package middleware

import (
	"github.com/gostartkit/api/proxy"
	"github.com/gostartkit/api/rbac"
	"github.com/webpkg/web"
)

// Chain call something before next callback
func Chain(next web.Callback, keys ...string) web.Callback {

	return func(ctx *web.Context) (web.Data, error) {

		before(ctx)

		err := bearerAuth(ctx, keys...)

		if err != nil {
			return nil, err
		}

		val, err := next(ctx)

		after(ctx)

		return val, err
	}
}

// Direct call something before next callback
func Direct(next web.Callback, keys ...string) web.Callback {

	return func(ctx *web.Context) (web.Data, error) {

		before(ctx)

		val, err := next(ctx)

		after(ctx)

		return val, err
	}
}

// before call before controller action
func before(ctx *web.Context) {
	ctx.AcceptContentType()
}

// after call after controller action
func after(ctx *web.Context) {

}

// bearerAuth bearer authorization
func bearerAuth(ctx *web.Context, keys ...string) error {

	accessToken := rbac.ParseBearerToken(ctx.Get("Authorization"))

	if accessToken == "" {
		return web.ErrUnauthorized
	}

	cat, err := proxy.GetAuthByAccessToken(accessToken)

	if err != nil {
		return err
	}

	ctx.Init(cat.UserID, cat.UserRight)

	if !rbac.Check(cat.UserRight, keys...) {
		return web.ErrForbidden
	}

	return nil
}
