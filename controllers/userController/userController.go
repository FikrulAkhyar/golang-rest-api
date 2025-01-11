package userController

import (
	"net/http"

	"github.com/FikrulAkhyar/golang-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Helper untuk menangani error dari database
func handleDBError(c *gin.Context, err error) bool {
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data User Tidak Ditemukan",
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Terjadi Kesalahan Pada Server",
			})
		}
		return true
	}
	return false
}

// Helper untuk validasi user
func validateUser(c *gin.Context, user *models.User) bool {
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data Tidak Valid",
		})
		return false
	}
	if user.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Nama User Tidak Boleh Kosong",
		})
		return false
	}
	if user.Age <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Umur User Harus Lebih Dari 0 dan Tidak Boleh Kosong",
		})
		return false
	}
	return true
}

func Index(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Mendapatkan Data User",
		"data":    users,
	})
}

func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if handleDBError(c, models.DB.First(&user, id).Error) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Mendapatkan Data User",
		"data":    user,
	})
}

func Store(c *gin.Context) {
	var user models.User

	if !validateUser(c, &user) {
		return
	}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Menambahkan Data User",
	})
}

func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if handleDBError(c, models.DB.First(&user, id).Error) {
		return
	}

	if !validateUser(c, &user) {
		return
	}

	models.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Mengubah Data User",
	})
}

func Delete(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if handleDBError(c, models.DB.First(&user, id).Error) {
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Menghapus Data User",
	})
}
