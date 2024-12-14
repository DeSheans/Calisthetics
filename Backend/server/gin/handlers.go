package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g ginServer) getAllExerciseCards(c *gin.Context) {
	log.Println("Gin /exercises endpoint handler starts. gin handler: ", g.hand)
	cards, err := g.hand.GetAllExerciseCards()
	if err == nil {
		c.JSON(http.StatusOK, cards)
	}
}
