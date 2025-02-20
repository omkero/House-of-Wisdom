package controllers

import (
	"github.com/gin-gonic/gin"
)

func Get_Test(ctx *gin.Context) {
	var response, err = testObject.Get_test(ctx)

	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, response)
}
