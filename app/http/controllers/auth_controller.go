package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/debug"
)

func dd(v ...any) {
	debug.DD(v...)
}

type AuthController struct {
	// Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Inject services
	}
}

func (r *AuthController) Register(ctx http.Context) http.Response {
	var user models.User
	ctx.Request().Bind(&user)
	user.Type = "admin"
	user.Password, _ = facades.Hash().Make(user.Password)
	user.DateOfBirth = "2009-01-01"
	err := facades.Orm().Query().Create(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"status":  1,
			"message": err.Error(),
		})
	}
	return ctx.Response().Json(http.StatusOK, map[string]any{
		"status":  1,
		"message": "success",
	})

}

func (r *AuthController) Index(ctx http.Context) http.Response {
	var user models.User
	ctx.Request().Bind(&user)

	var password string = user.Password
	err := facades.Orm().Query().Where("email", user.Email).First(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"status":  1,
			"message": "User is not found",
		})
	}
	// Verify password
	if !facades.Hash().Check(password, user.Password) {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"status":  1,
			"message": "Password is not matched",
		})
	}

	token, err := facades.Auth(ctx).Login(&user)
	dd(token)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, map[string]any{
			"status":  1,
			"message": "Authentication failed",
		})
	}

	return ctx.Response().Json(http.StatusOK, map[string]any{
		"status":  0,
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}
