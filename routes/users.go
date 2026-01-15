package routes

import (
	"net/http"

	"github.com/abhijitpattar/gin-rest-go/models"
	"github.com/abhijitpattar/gin-rest-go/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not CREATE user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": " user created", "user": user.Email})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request data"})
		return
	}

	err = user.ValidateCredentails()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate login token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": " Login successful", "token": token})
}
