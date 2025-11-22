package handler

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-llm-chatbot/internal/core/port"
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

type ImagePromptReq struct {
	Prompt string `form:"prompt" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
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

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func (h *PromptHandler) ImageToTextPrompt(c *gin.Context) {
	var req ImagePromptReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(req.File, "./assets/" + req.File.Filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// image to text prompt service logic
	res, err := h.svc.ImageToTextPrompt(c, req.Prompt, req.File)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
