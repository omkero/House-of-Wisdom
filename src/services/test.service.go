package services

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type TestService struct{}

func (art *TestService) Get_test(ctx *gin.Context) (gin.H, error) {
	var FormObject FormData
	err := ctx.ShouldBind(&FormObject)

	if err != nil {
		return gin.H{
			"message": "file cannot be bined",
		}, err
	}

	var dirPath string = "public/static/assets/"

	var timeNowUnix int64 = time.Now().Unix()
	var timeStr string = strconv.Itoa(int(timeNowUnix))

	var newFileName string = timeStr + "_" + FormObject.Image.Filename
	var path string = dirPath + newFileName

	uploadErr := ctx.SaveUploadedFile(FormObject.Image, path)
	if uploadErr != nil {
		return gin.H{
			"message": "file cannot be saved",
		}, uploadErr
	}

	return gin.H{
		"message":   "file has been saved",
		"file-name": FormObject.Image.Filename,
	}, nil
}

func CreateFile(text string, fileName string, path string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var string = "hello"

	_, errf := file.Write([]byte(string))
	if errf != nil {
		log.Fatal(errf)
	}
	defer file.Close()

	fmt.Println("file created")
}
