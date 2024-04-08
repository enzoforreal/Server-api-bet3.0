package predictions

import (
	"github.com/gin-gonic/gin"
)

func FetchPredictions(c *gin.Context) {
	// Implémentez la logique pour récupérer et afficher les prédictions ici
	c.JSON(200, gin.H{"message": "Prédictions des matchs."})
}
