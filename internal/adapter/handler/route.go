package handler

import (
	"go-gemini-llm/internal/adapter/config"

	"github.com/gin-gonic/gin"
)

type Route struct {
	r *gin.Engine
}

func InitRoute(
	promptHandler *PromptHandler,
) *Route {
	r := gin.New()

	// route groupings
	pb := r.Group("/api/v1")

	// public routes
	pb.POST("/prompt/text-to-text", promptHandler.TextToTextPrompt)

	// protected routes

	return &Route{
		r: r,
	}
}

func (r *Route) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}
