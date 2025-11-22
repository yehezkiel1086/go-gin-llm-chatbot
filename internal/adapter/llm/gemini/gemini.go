package gemini

import (
	"context"
	"go-gemini-llm/internal/adapter/config"

	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
	conf *config.LLM
}

func InitGemini(ctx context.Context, conf *config.LLM) (*Gemini, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  conf.APIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &Gemini{
		conf: conf,
		client: client,
	}, nil
}

func (g *Gemini) GenerateTextToText(ctx context.Context, prompt string) (string, error) {
	result, err := g.client.Models.GenerateContent(
		ctx,
		g.conf.AIModel,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
