package routes

import (
	"wisdom/src/controllers"
	"wisdom/src/middleware"

	"github.com/gin-gonic/gin"
)

type ArticlesRouter struct{}

func (atr *ArticlesRouter) ArticlesGroup(base_path string, r *gin.Engine) {
	var router = r.Group(base_path)
	//	router.Use(middleware.AuthenticateMiddleware()) // user must be authenticated to use this route
	//router.Use(middleware.RateLimitMiddleware(5, 1, 5)) // 5 requests per 5 seconds blocked until 5 seconds pass

	{
		_getAllArticle("/get_all_articles", router)
		_getArticleByTitle("/get_article", router)
		_createNewArticle("/create_new_article", router)
		_deleteArticle("/delete_article", router)
		_editArticle("/edit_article", router)
	}

}

func _getAllArticle(path string, router *gin.RouterGroup) {

	router.POST(path, func(ctx *gin.Context) {
		controllers.Get_all_articles_controller(ctx)
	})
}

func _getArticleByTitle(path string, router *gin.RouterGroup) {

	router.POST(path, func(ctx *gin.Context) {
		controllers.Get_article_controller(ctx)
	})
}

func _createNewArticle(path string, router *gin.RouterGroup) {
	router.POST(path, middleware.AuthenticateMiddleware(), func(ctx *gin.Context) {
		controllers.Create_new_article_controller(ctx)
	})
}

func _deleteArticle(path string, router *gin.RouterGroup) {
	router.DELETE(path, middleware.AuthenticateMiddleware(), func(ctx *gin.Context) {
		controllers.Delete_article_controller(ctx)
	})
}

func _editArticle(path string, router *gin.RouterGroup) {
	router.POST(path, middleware.AuthenticateMiddleware(), func(ctx *gin.Context) {
		controllers.Edit_article_controller(ctx)
	})
}
