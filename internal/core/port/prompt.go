package port

import (
	"context"
	"mime/multipart"
)

type PromptService interface {
	TextToTextPrompt(ctx context.Context, prompt string) (string, error)
	ImageToTextPrompt(ctx context.Context, prompt string, file *multipart.FileHeader) (string, error)
}
