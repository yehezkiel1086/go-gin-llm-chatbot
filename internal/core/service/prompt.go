package service

import (
	"context"
	"go-gemini-llm/internal/adapter/llm/gemini"
)

type PromptService struct {
	llm *gemini.Gemini
}

func InitPromptService(llm *gemini.Gemini) *PromptService {
	return &PromptService{
		llm: llm,
	}
}

func (s *PromptService) TextToTextPrompt(ctx context.Context, prompt string) (string, error) {
	res, err := s.llm.GenerateTextToText(ctx, prompt)
	if err != nil {
		return "", err
	}
	return res, nil
}
