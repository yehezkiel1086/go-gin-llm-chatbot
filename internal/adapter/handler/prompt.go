package handler

import (
	"go-gemini-llm/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromptHandler struct {
	svc port.PromptService
}

func InitPromptHandler(svc port.PromptService) *PromptHandler {
	return &PromptHandler{
		svc: svc,
	}
}

type TextPromptReq struct {
	Prompt string `json:"prompt" binding:"required"`
}

type TextPromptRes struct {
	Prompt string `json:"prompt"`
}

func (h *PromptHandler) TextToTextPrompt(c *gin.Context) {
	var req TextPromptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.TextToTextPrompt(c, req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, TextPromptRes{
		Prompt: res,
	})
}
