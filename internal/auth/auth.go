package auth

import (
	"github.com/enzof/server-app-bet3.0/pkg/config"
	"github.com/enzof/server-app-bet3.0/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {

	db := config.GetDB()

	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erreur lors du hachage du mot de passe"})
		return
	}
	newUser.Password = hashedPassword

	result := db.Create(&newUser)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})

		return
	}
	c.JSON(200, gin.H{"message": "Inscription réussie!"})
}

func LoginUser(c *gin.Context) {

	db := config.GetDB()
	var loginDetails UserLogin
	var user User

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(400, gin.H{"error": "Invalid login details"})
		return
	}

	result := db.Where("email = ?", loginDetails.Email).First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid login credentials"})
		return
	}

	token, err := util.GenerateToken(user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating JWT token"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful", "token": token})
}
