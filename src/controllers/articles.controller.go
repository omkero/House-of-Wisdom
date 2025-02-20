package controllers

import (
	"github.com/gin-gonic/gin"
)

func Get_all_articles_controller(ctx *gin.Context) {
	response, err := articleObject.Get_all_articles_service(ctx)

	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, response)

}
func Get_article_controller(ctx *gin.Context) {
	response, err := articleObject.Get_article_service(ctx)

	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, response)
}

func Create_new_article_controller(ctx *gin.Context) {
	response, err := articleObject.Create_new_article_service(ctx)

	if err != nil {
		ctx.JSON(404, response)
		return
	}

	ctx.JSON(200, response)
}

func Delete_article_controller(ctx *gin.Context) {
	_, status_code, err := articleObject.Delete_user_article(ctx)

	if err != nil {
		ctx.JSON(status_code, gin.H{
			"message":     err.Error(),
			"status_code": status_code,
		})
		return
	}

	ctx.JSON(status_code, gin.H{
		"message":     "Article has been deleted",
		"status_code": status_code,
	})
}

func Edit_article_controller(ctx *gin.Context) {
	_, status_code, err := articleObject.Update_user_article(ctx)

	if err != nil {
		ctx.JSON(status_code, gin.H{
			"message":     err.Error(),
			"status_code": status_code,
		})
		return
	}

	ctx.JSON(status_code, gin.H{
		"message":     "Article has been updated",
		"status_code": status_code,
	})
}
