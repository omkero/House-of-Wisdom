package services

import (
	"fmt"
	"os"
	"wisdom/src/models"
	"wisdom/src/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (aus *AuthService) LoginUserServices(ctx *gin.Context) (gin.H, gin.H) {

	//var UserData models.AuthUserData
	var jwt utils.Jwt
	var UserAuth models.AuthLoginModel
	err := ctx.ShouldBindJSON(&UserAuth)

	errorResponse := gin.H{
		"message":       "Error cannot get the data",
		"error_message": err,
	}

	if UserAuth.Email == "" || UserAuth.Password == "" {
		errorResponse = gin.H{
			"message":     "Email or Password not provided",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}
	}

	if err != nil {
		//fmt.Printf("Error while trying to bind UserAuth")
		return nil, errorResponse
	}

	response, err := auth_repository_object.Select_user_data(UserAuth)

	if err != nil {
		errorResponse := gin.H{
			"message":     "User not found",
			"status_code": 404,
			"error_type":  "USER_NOT_FOUND",
		}
		return nil, errorResponse
	}

	isMatch := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(UserAuth.Password))

	// verify password
	if isMatch != nil {
		errorResponse = gin.H{
			"message":     "Error this password incorrect",
			"status_code": 401,
			"error_type":  "AUTHENTICATION_ERROR",
		}
		return nil, errorResponse
	}

	var singedToken = jwt.SignAuthToken(string(response.Username), response.User_id, []byte(os.Getenv("LOGIN_PRIVATE_KEY")))
	var responseData = gin.H{
		"email_address":  response.Email_address,
		"user_id":        response.User_id,
		"username":       response.Username,
		"articles_count": response.Articles_count,
		"created_at":     response.Created_at,
		"token":          singedToken,
		"status_code":    200,
	}
	return responseData, nil
}

func (aus *AuthService) SignupUserService(ctx *gin.Context) (gin.H, int, gin.H) {

	var UserAuth models.AuthSignupModel
	err := ctx.ShouldBindJSON(&UserAuth)

	errorResponse := gin.H{
		"message":       "Error cannot signup",
		"error_message": err,
	}
	successResponse := gin.H{"message": ""}

	if UserAuth.Email == "" {
		errorResponse = gin.H{
			"message":     "Email is not provided",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}

		return nil, 400, errorResponse
	}

	if UserAuth.Username == "" {
		errorResponse = gin.H{
			"message":     "Username is not provided",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}
		return nil, 400, errorResponse
	}

	if UserAuth.Password == "" {
		errorResponse = gin.H{
			"message":     "Password is not provided",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}
		return nil, 400, errorResponse
	}

	isExist, err := auth_repository_object.Verify_user_exist_by_email(UserAuth)

	fmt.Print(err)

	if err != nil {
		errorResponse = gin.H{
			"message":     "Ops an error has occured",
			"status_code": 500,
			"error_type":  "SERVER_ERROR",
		}
		return nil, 500, errorResponse
	}

	if isExist {
		errorResponse = gin.H{
			"message":     "This user already exist.",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}
		return nil, 401, errorResponse
	}

	isInserted := auth_repository_object.Insert_new_user(UserAuth)

	if !isInserted {
		errorResponse = gin.H{
			"message":     "User not been created",
			"status_code": 400,
			"error_type":  "VALIDATION_ERROR",
		}
		return nil, 401, errorResponse
	}

	if isInserted {
		successResponse = gin.H{
			"message":     "You have been seccessfully created",
			"status_code": 201,
		}
	}

	return successResponse, 201, nil
}
