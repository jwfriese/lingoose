package main

import (
	"context"
	"fmt"

	"github.com/jwfriese/lingoose/legacy/chat"
	"github.com/jwfriese/lingoose/legacy/prompt"
	"github.com/jwfriese/lingoose/llm/openai"
)

func main() {

	chat := chat.New(
		chat.PromptMessage{
			Type:   chat.MessageTypeSystem,
			Prompt: prompt.New("You are a professional joke writer"),
		},
		chat.PromptMessage{
			Type:   chat.MessageTypeUser,
			Prompt: prompt.New("Write a joke about geese"),
		},
	)

	llm := openai.NewChat()

	err := llm.ChatStream(context.Background(), func(output string) {
		fmt.Printf("%s", output)
	}, chat)
	if err != nil {
		panic(err)
	}

	fmt.Println()

}
