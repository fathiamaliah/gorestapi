package productcontroller

import (
	"encoding/json"
	"gorestapi/database"
	"gorestapi/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := database.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data isnt found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return

		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "There is nothing to be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully updated"})

}

func Delete(c *gin.Context) {
	var product models.Product
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if database.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product can't be deleted"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted"})

}
