package controllers

import (
	"github.com/gin-gonic/gin"
)

func LoginUserController(ctx *gin.Context) {
	response, err := authObject.LoginUserServices(ctx)

	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, response)
}

func SignupUserController(ctx *gin.Context) {
	response, status_code, error_message := authObject.SignupUserService(ctx)

	if error_message != nil {
		ctx.JSON(status_code, error_message)
		return
	}

	ctx.JSON(status_code, response)
}
