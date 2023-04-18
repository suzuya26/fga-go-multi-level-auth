package controllers

import (
	"net/http"
	"sesi_12/database"
	"sesi_12/helpers"
	"sesi_12/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, product)

}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id = ?", productId).Updates(models.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	product := models.Product{}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Preload("User").First(&product, productId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, product)
}

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	product := []models.Product{}

	err := db.Preload("User").Find(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	product := models.Product{}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Delete(&product, productId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusNoContent, nil)
}
