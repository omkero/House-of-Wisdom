package routes

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	var base_api string = "api/v1"
	// define the structs
	RootRouter := RootRouter{}
	ArticlesRouter := ArticlesRouter{}
	AuthRouter := AuthRouter{}
	TestRouter := TestRouter{}

	go RootRouter.RootGroup(base_api+"/", r)
	go ArticlesRouter.ArticlesGroup(base_api+"/articles", r)
	go AuthRouter.AuthGroup(base_api+"/auth", r)
	go TestRouter.testGroup(base_api+"/test", r)
}
