package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wellminozzo/LoginGo-React-teste/db"
	"github.com/wellminozzo/LoginGo-React-teste/models"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := db.DB.Create(user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	res := db.DB.Find(&user, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	res := db.DB.Find(&users)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no users found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateUser models.User
	res := db.DB.Model(&updateUser).Where("id = ?", id).Updates(user)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not updated",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": updateUser,
	})

}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	res := db.DB.Find(&user, id)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	db.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

var secretKey = []byte("tokenSeguro")

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(secretKey)
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if user.Username == "admin" && user.Password == "123456" {
		token, _ := generateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	var dbUser models.User
	res := db.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&dbUser)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Usuário ou senha inválidos",
		})
		return
	}

	token, _ := generateToken(dbUser.Username)
	c.JSON(http.StatusOK, gin.H{
		"user":  dbUser,
		"token": token,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
