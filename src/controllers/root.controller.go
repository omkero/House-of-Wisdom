package controllers

import (
	"github.com/gin-gonic/gin"
)

func RootControllerGET(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "works get root",
	})
}

func RootControllerPOST(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data": "works post root",
	})
}
