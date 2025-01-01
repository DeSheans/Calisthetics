package gin

import (
	"fmt"
	"gin-web-server/app/handlers"
	"gin-web-server/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	cfg  Config
	hand handlers.Handler
}

// NewGinServer ...
func NewGinServer(cfg Config, handler handlers.Handler) server.Server {
	return ginServer{cfg: cfg, hand: handler}
}

func (g ginServer) Run() error {
	router := gin.Default()

	corsConfiguration(router)

	router.GET("/exercises", g.getAllExerciseCards)
	router.GET("/programs", g.getAllProgramCards)
	router.GET("/programs/:id", g.getProgramByID)
	router.GET("/exercises/:id", g.getExerciseByID)
	router.GET("/exerciseFilter", g.getExerciseFilters)
	router.GET("/programFilter", g.getProgramFilters)

	addr := fmt.Sprintf(":%s", g.cfg.Port)
	if err := router.Run(addr); err != nil {
		return err
	}
	return nil
}

func corsConfiguration(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
