package routes

import (
	"net/http"

	"github.com/extimsu/JsonRestApi/models"
	"github.com/extimsu/JsonRestApi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind JSON" + err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, users)
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind JSON" + err.Error()})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create JWT token " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
