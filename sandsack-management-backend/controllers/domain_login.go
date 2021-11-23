package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"team2/sandsack-management-backend/functions"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// Registration
// @Description Registrates new user in system. Sends verification token to email of the user
// @Tags Authentification
// @Accept json
// @Param input body models.NewPerson true "Registration input"
// @Success 200 {object} map[string]string
// @Router /registration [post]
func (a *App) Login(c *gin.Context){
	var input models.Login

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("Registration error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect request",
		})
		return
	}

	email := strings.ToLower(input.Email)

	exists := service.CheckUserExists(a.DB, email)
	if !exists {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrCode: http.StatusNotFound,
			ErrMessage: "user not found",
		})
		return
	}

	tokens, err := service.GenerateTokens(a.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	user, err := service.GetUserByEmail(a.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	// check password is correct
	ok := functions.CheckPasswordHash(input.Password, user.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "wrong password",
		})
		return
	}

	// if email is not verified, user cannot be logged in
	if !user.IsEmailVerified {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "should verify email",
		})
		return
	}

	// if user is not activated, not possible to log in
	if !user.IsActivated {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "user is deactivated",
		})
		return
	}

	var refreshToken string
	// if we do not have token in database, then put it in
	if len(user.Token) == 0 {
		if err := service.UpdateUserToken(a.DB, email, tokens["refresh_token"]); err != nil {
			log.Println("UpdateUserToken error", err.Error())
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				ErrCode: http.StatusInternalServerError,
				ErrMessage: "something went wrong",
			})
			return
		}
		refreshToken = tokens["refresh_token"]
	} else {
		refreshToken = user.Token
	}

	tokens["refresh_token"] = refreshToken

	// return access and refresh tokens
	c.JSON(http.StatusOK, tokens)
	return
}




func (a *App) Hello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
	return
}