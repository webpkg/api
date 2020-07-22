package middleware

import (
	"github.com/webpkg/api/proxy"
	"github.com/webpkg/api/rbac"
	"github.com/webpkg/web"
)

// Chain call something before next callback
func Chain(next web.Callback, keys ...string) web.Callback {

	return func(ctx *web.Context) {

		before(ctx)

		err := bearerAuth(ctx, keys...)

		if err != nil {

			if err == rbac.ErrPermissionDenied {
				ctx.WriteHeader(403)
			} else {
				ctx.WriteHeader(401)
			}

			ctx.WriteJSON(err.Error())

			return
		}

		next(ctx)

		after(ctx)
	}
}

// Direct call something before next callback
func Direct(next web.Callback) web.Callback {

	return func(ctx *web.Context) {

		before(ctx)

		next(ctx)

		after(ctx)
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
		return rbac.ErrPermissionDenied
	}

	return nil
}
