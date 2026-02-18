package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/spf13/cobra"
)

func main() {
	_ = godotenv.Load(".env.local")

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "error: GROQ_API_KEY not set")
		os.Exit(1)
	}

	client := openai.NewClient(
		option.WithBaseURL("https://api.groq.com/openai/v1"),
		option.WithAPIKey(apiKey),
	)

	var rootCmd = &cobra.Command{
		Use:   "fixe <text>",
		Short: "Fix English grammar mistakes",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			text := args[0]

			completion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
				Model: openai.ChatModel("groq/compound-mini"),
				Messages: []openai.ChatCompletionMessageParamUnion{
					openai.SystemMessage("You are a grammar corrector. Fix the grammar mistakes in the user's text. Only output the corrected text, nothing else. No explanations, no quotes, just the fixed text."),
					openai.UserMessage(text),
				},
			})
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}

			fmt.Print(completion.Choices[0].Message.Content)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
