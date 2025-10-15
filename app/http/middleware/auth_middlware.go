package middleware

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AuthMiddlware() http.Middleware {
	return func(ctx http.Context) {
		authorization := ctx.Request().Header("Authorization")
		bearerToken := strings.Split(authorization, " ")
		bearer, token := bearerToken[0], bearerToken[1]
		if bearer != "Bearer" || token == "" {
			ctx.Response().String(http.StatusNotFound, "Token is not provided").Abort()
		}
		user, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			ctx.Response().String(http.StatusNotFound, "Invalid token").Abort()
		}
		ctx.WithValue("user", user)
		ctx.Request().Next()
	}
}
