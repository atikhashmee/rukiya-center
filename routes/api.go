package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Api() {
	userController := controllers.NewUserController()
	authController := controllers.NewAuthController()
	facades.Route().Get("/users/{id}", userController.Show)

	facades.Route().Prefix("api").Group(func(router route.Router) {
		router.Post("/register", authController.Register)
		router.Post("/login", authController.Login)
		router.Middleware(middleware.AuthMiddlware()).Group(func(router route.Router) {
			router.Post("/me", authController.Profile)
		})
	})

}
