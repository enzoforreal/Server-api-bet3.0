package main

import (
	"github.com/enzof/server-app-bet3.0/internal/auth"
	"github.com/enzof/server-app-bet3.0/internal/auth/predictions"
	"github.com/enzof/server-app-bet3.0/internal/middleware"
	"github.com/enzof/server-app-bet3.0/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&auth.User{})
	if err != nil {
		panic("Échec lors de la migration automatique.")
	}
}

func main() {
	config.InitConfig()
	db, err := config.InitDB()
	if err != nil {
		panic("Échec de l'initialisation de la base de données.")
	}
	migrate(db)

	realUserDB := &auth.RealUserDB{DB: db}
	router := gin.Default()

	httpRouter := gin.New()
	httpRouter.GET("/*any", func(c *gin.Context) {
		c.Redirect(301, "https://"+c.Request.Host+c.Request.RequestURI)
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bienvenue sur l'Application de Conseils de Paris Sportifs!"})
	})

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", func(c *gin.Context) { auth.RegisterUser(c, realUserDB) })
		authRoutes.POST("/login", func(c *gin.Context) { auth.LoginUser(c, realUserDB) })
	}

	protectedRoutes := router.Group("/").Use(middleware.JWTAuthMiddleware())
	{
		protectedRoutes.GET("/predictions", predictions.FetchPredictions)
	}

	go func() {
		router.RunTLS(":443", "/mnt/c/Users/enzof/Project/Server-app-bet3.0/cert.pem", "/mnt/c/Users/enzof/Project/Server-app-bet3.0/decrypted_key.pem")
	}()

	httpRouter.Run(":8080")
}
