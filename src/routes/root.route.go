package routes

import (
	"wisdom/src/controllers"
	"wisdom/src/middleware"

	"github.com/gin-gonic/gin"
)

type RootRouter struct{}

func (rtr *RootRouter) RootGroup(base_path string, r *gin.Engine) {
	var router = r.Group(base_path)
	//router.Use(middleware.VerifyAuthKey())

	{
		_getRoot("/", router)
		_postRoot("/", router)
	}

}

func _getRoot(path string, router *gin.RouterGroup) {

	router.GET(path, func(ctx *gin.Context) {
		controllers.RootControllerGET(ctx)
	})
}

func _postRoot(path string, router *gin.RouterGroup) {
	router.POST(path, middleware.AuthenticateMiddleware(), func(ctx *gin.Context) {
		controllers.RootControllerPOST(ctx)
	})
}
