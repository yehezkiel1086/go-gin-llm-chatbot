package port

import "context"

type PromptService interface {
	TextToTextPrompt(ctx context.Context, prompt string) (string, error)
}
