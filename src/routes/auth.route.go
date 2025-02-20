package routes

import (
	"wisdom/src/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (art *AuthRouter) AuthGroup(base_path string, r *gin.Engine) {
	var router = r.Group(base_path)

	{
		_auth_login("/login", router)
		_auth_signup("/signup", router)
	}
}

func _auth_login(path string, router *gin.RouterGroup) {
	router.POST(path, func(ctx *gin.Context) {
		controllers.LoginUserController(ctx)
	})
}

func _auth_signup(path string, router *gin.RouterGroup) {
	router.POST(path, func(ctx *gin.Context) {
		controllers.SignupUserController(ctx)
	})
}
