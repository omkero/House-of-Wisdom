package utils

import (
	"strconv"
	"time"
	"wisdom/src/models"

	"github.com/gin-gonic/gin"
)

func SaveImageToDir(ArticleBody *models.CreateArticle, ctx *gin.Context) (string, string, string, error) {

	var dirPath string = "public/static/assets/"

	// get current time in unix
	var timeNowUnix int64 = time.Now().Unix()
	var timeStr string = strconv.Itoa(int(timeNowUnix)) // convert unix int64 to string

	var newFileName string = timeStr + "_" + ArticleBody.File.Filename // inject unix time to the file name (the file must be unique)
	var path string = dirPath + newFileName                            // the full path with newFileName
	var publicPath string = "assets/" + newFileName                    // public path

	uploadErr := ctx.SaveUploadedFile(ArticleBody.File, path) // execute

	// if something wrong
	if uploadErr != nil {
		return "file cannot be saved", "", "", uploadErr
	}

	return publicPath, dirPath, newFileName, nil
}

func SaveImageToDirEdit(ArticleBody *models.EditArticleForm, ctx *gin.Context) (string, string, string, error) {

	var dirPath string = "public/static/assets/"

	// get current time in unix
	var timeNowUnix int64 = time.Now().Unix()
	var timeStr string = strconv.Itoa(int(timeNowUnix)) // convert unix int64 to string

	var newFileName string = timeStr + "_" + ArticleBody.File.Filename // inject unix time to the file name (the file must be unique)
	var path string = dirPath + newFileName                            // the full path with newFileName
	var publicPath string = "assets/" + newFileName                    // public path

	uploadErr := ctx.SaveUploadedFile(ArticleBody.File, path) // execute

	// if something wrong
	if uploadErr != nil {
		return "file cannot be saved", "", "", uploadErr
	}

	return publicPath, dirPath, newFileName, nil
}
