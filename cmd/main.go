package main

import (
	"context"
	"fmt"
	"go-gemini-llm/internal/adapter/config"
	"go-gemini-llm/internal/adapter/handler"
	"go-gemini-llm/internal/adapter/llm/gemini"
	"go-gemini-llm/internal/core/service"
	"log"
)

func main() {
    // import .env configs
    conf, err := config.InitConfig()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(".env configuration imported successfully")

    // init context
    ctx := context.Background()

    // init gemini llm
    llm, err := gemini.InitGemini(ctx, conf.LLM)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Gemini LLM initialized successfully")

    // dependency injections
    promptSvc := service.InitPromptService(llm)
    promptHandler := handler.InitPromptHandler(promptSvc)

    // init routing
    r := handler.InitRoute(
        promptHandler,
    )

    // serve api
    if err := r.Serve(conf.HTTP); err != nil {
        log.Fatal(err)
    }
}
