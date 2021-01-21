package middleware

import (
	"github.com/webpkg/api/proxy"
	"github.com/webpkg/api/rbac"
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
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetHeader("access-control-allow-origin", "*")
}

// after call after controller action
func after(ctx *web.Context) {

}

// bearerAuth bearer authorization
func bearerAuth(ctx *web.Context, keys ...string) error {

	auth := ctx.GetHeader("Authorization")

	accessToken, err := rbac.TryParseBearerToken(auth)

	if err != nil {
		return err
	}

	cat, err := proxy.GetAuthByAccessToken(accessToken)

	if err != nil {
		return err
	}

	ctx.UserID = cat.UserID

	if !rbac.Check(cat.Right, keys...) {
		return web.ErrForbidden
	}

	return nil
}
