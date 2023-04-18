package controllers

import (
	"net/http"
	"sesi_12/database"
	"sesi_12/helpers"
	"sesi_12/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
		"level":     user.Level,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	passwordClient := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err":     "unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	isValid := helpers.ComparePass([]byte(user.Password), []byte(passwordClient))
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err":     "unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email, user.Level)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
