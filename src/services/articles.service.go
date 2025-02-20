package services

import (
	"errors"
	"fmt"
	"os"
	"wisdom/src/models"
	"wisdom/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ArticleService struct{}

func (art *ArticleService) Get_all_articles_service(ctx *gin.Context) ([]models.Article, gin.H) {

	var PageNumber models.GetArticlesLIMIT

	bindingErr := ctx.ShouldBindJSON(&PageNumber)
	if bindingErr != nil {
		errMsg := gin.H{
			"message": "Missing required inputs",
		}
		return nil, errMsg
	}

	var limit int = 10
	var offset int = (PageNumber.Page - 1) * limit

	data, err := repositoryObject.Select_all_Articles_repo(limit, offset)

	if err != nil {
		msg := gin.H{
			"data": "cannot find data",
		}

		return nil, msg
	}

	return data, nil
}

func (art *ArticleService) Get_article_service(ctx *gin.Context) (models.Article, gin.H) {

	var Article models.GetArticle

	bindingErr := ctx.ShouldBindJSON(&Article)
	if bindingErr != nil {
		errMsg := gin.H{
			"message": "Missing required inputs",
		}
		return models.Article{}, errMsg
	}

	data, err := repositoryObject.Select_Article_repo(Article.Article_title)

	if err != nil {
		msg := gin.H{
			"data": "cannot find data",
		}

		return models.Article{}, msg
	}

	return data, nil
}

func (art *ArticleService) Create_new_article_service(ctx *gin.Context) (gin.H, error) {
	J := utils.Jwt{}
	var ArticleBody models.CreateArticle

	jwtToken, err := J.DecodeJwtToken(ctx)
	bindingErr := ctx.ShouldBind(&ArticleBody)
	fmt.Println(bindingErr)

	if err != nil {
		errMsg := gin.H{
			"message":     "Cannot decode",
			"status_code": 500,
			"error_type":  "SERVER_ERROR",
		}

		return errMsg, nil
	}
	if bindingErr != nil {
		errMsg := gin.H{
			"message":     "Missing required inputs",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}

		return errMsg, nil
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		errMsg := gin.H{
			"message": "error cannot verify this token",
		}
		return errMsg, errors.New("error cannot verify this token")
	}

	// sub meaning the username from the token
	sub, ok := claims["sub"].(string)
	if !ok {
		errMsg := gin.H{
			"message":     "error cannot find sub",
			"status_code": 500,
			"error_type":  "VALIDATION_ERROR",
		}
		return errMsg, errors.New("error cannot find sub")
	}

	/*
		The issue with the bis claim lies in how you're attempting to extract it from the claims object.
		Specifically, the bis value in the claims appears to be a number (23 in this example).
		When attempting to assert it as an int (claims["bis"].(int)),
		it fails because JSON unmarshalling usually interprets numbers as float64 in Go.
	*/

	bisFloat, ok := claims["bis"].(float64)

	if !ok {
		errMsg := gin.H{
			"message":     "error cannot find bis",
			"status_code": 500,
			"error_type":  "VALIDATION_ERROR",
		}
		return errMsg, errors.New("error cannot find bis")
	}

	bisInt := int(bisFloat)

	path, dirPath, fileName, errB := utils.SaveImageToDir(&ArticleBody, ctx)
	if errB != nil {
		return nil, errB
	}

	resp, err := repositoryObject.Insert_new_article(ArticleBody, sub, bisInt, path)
	if err != nil {

		// remove the file from the disk
		errR := os.Remove(dirPath + fileName)
		if errR != nil {
			fmt.Println(errR)
		}
		errMsg := gin.H{
			"error":       err,
			"message":     resp,
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}
		return errMsg, err
	}

	secMsg := gin.H{
		"data": resp,
	}
	return secMsg, nil
}

func (art *ArticleService) Delete_user_article(ctx *gin.Context) (gin.H, int, error) {
	J := utils.Jwt{}

	articleID := models.ArticleID{}

	jwtToken, err := J.DecodeJwtToken(ctx)
	bindingErr := ctx.ShouldBindJSON(&articleID)

	if err != nil {
		return nil, 500, err
	}

	if bindingErr != nil {
		return nil, 500, bindingErr
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, 401, errors.New("error cannot verify this token")
	}
	// Get the 'sub' claim
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, 401, errors.New("error cannot find sub")
	}

	response, status_code, err := repositoryObject.Delete_article(articleID, sub)

	if err != nil {
		return nil, status_code, err
	}

	responseSuccess := gin.H{
		"message": response,
	}

	return responseSuccess, 201, nil
}

func (art *ArticleService) Update_user_article(ctx *gin.Context) (gin.H, int, error) {
	J := utils.Jwt{}

	articleID := models.EditArticleForm{}

	jwtToken, err := J.DecodeJwtToken(ctx)
	bindingErr := ctx.ShouldBind(&articleID)

	if err != nil {
		return nil, 500, err
	}

	if bindingErr != nil {
		return nil, 500, errors.New("error missing required fields")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, 401, errors.New("error cannot verify this token")
	}

	// Get the 'bis' claim
	bis, ok := claims["bis"].(float64)
	if !ok {
		return nil, 401, errors.New("error cannot find sub")
	}

	response, status_code, err := repositoryObject.Update_article(articleID, int(bis), ctx)

	if err != nil {

		return nil, status_code, err
	}

	responseSuccess := gin.H{
		"message": response,
	}

	return responseSuccess, 201, nil
}
