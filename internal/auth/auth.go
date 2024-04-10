package auth

import (
	"github.com/enzof/server-app-bet3.0/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context, db UserDB) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := util.ValidateStruct(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erreur lors du hachage du mot de passe"})
		return
	}
	newUser.Password = hashedPassword

	if err := db.CreateUser(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Inscription réussie!"})
}

func LoginUser(c *gin.Context, db UserDB) {
	var loginDetails UserLogin
	var user *User // Change to pointer type

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(400, gin.H{"error": "Invalid login details"})
		return
	}

	user, err := db.GetUserByEmail(loginDetails.Email) // This now correctly matches the expected type
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
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
