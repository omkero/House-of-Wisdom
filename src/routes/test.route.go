package routes

import (
	"wisdom/src/controllers"

	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (atr *TestRouter) testGroup(base_path string, r *gin.Engine) {
	var router = r.Group(base_path)
	//	router.Use(middleware.AuthenticateMiddleware())     // user must be authenticated to use this route
	//	router.Use(middleware.RateLimitMiddleware(5, 5, 5)) // 5 requests per 5 seconds blocked until 5 seconds pass

	{
		_getTest("/get_test", router)
	}

}

func _getTest(path string, router *gin.RouterGroup) {

	router.POST(path, func(ctx *gin.Context) {
		controllers.Get_Test(ctx)
	})
}
