package gin

import (
	"fmt"
	"gin-web-server/app/handlers"
	"gin-web-server/server"

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

	router.GET("/exercises", g.getAllExerciseCards)

	// localhost:8080
	if err := router.Run(fmt.Sprintf("%s:%s", g.cfg.Host, g.cfg.Port)); err != nil {
		return err
	}

	return nil
}
