package bookcontroller

import (
	"encoding/json"
	"net/http"

	"RestApi/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func Show(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := models.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func Create(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func Update(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	if models.DB.Model(&book).Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat diupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate!" })
}

func Delete(c *gin.Context) {
	var book models.Book
	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
