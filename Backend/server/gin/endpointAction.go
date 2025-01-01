package gin

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (g ginServer) getAllExerciseCards(c *gin.Context) {
	cards, err := g.hand.GetAllExerciseCards(c.Request.URL.Query())
	if err != nil {
		log.Println("Failed to get all exercise cards due to error: ", err)

	} else {
		c.JSON(http.StatusOK, cards)
	}
}

func (g ginServer) getAllProgramCards(c *gin.Context) {
	programs, err := g.hand.GetAllProgramCards(c.Request.URL.Query())
	if err != nil {
		log.Println("Failed to get all program cards due to error: ", err)
	} else {
		c.JSON(http.StatusOK, programs)
	}
}

func (g ginServer) getProgramByID(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	}
	program, err := g.hand.GetProgramByID(id)
	if err != nil {
		log.Println("Failed to get all program due to error: ", err)
	} else {
		c.JSON(http.StatusOK, program)
	}
}

func (g ginServer) getExerciseByID(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	}
	exercise, err := g.hand.GetExerciseByID(id)
	if err != nil {
		log.Println("Failed to get all exercise due to error: ", err)
	} else {
		c.JSON(http.StatusOK, exercise)
	}
}

func (g ginServer) getExerciseFilters(c *gin.Context) {
	filters, err := g.hand.GetExerciseFilters()
	if err != nil {
		log.Println("Failed to get all exercise filters dut to error: ", err)
	} else {
		c.JSON(http.StatusOK, filters)
	}
}

func (g ginServer) getProgramFilters(c *gin.Context) {
	filters, err := g.hand.GetProgramFilters()
	if err != nil {
		log.Println("Failed to get all program filters dut to error: ", err)
	} else {
		c.JSON(http.StatusOK, filters)
	}
}
