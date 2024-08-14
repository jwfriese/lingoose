package main

import (
	"context"
	"fmt"

	"github.com/jwfriese/lingoose/llm/antropic"
	"github.com/jwfriese/lingoose/thread"
)

func main() {
	antropicllm := antropic.New().WithModel("claude-3-opus-20240229").WithStream(
		func(response string) {
			if response != antropic.EOS {
				fmt.Print(response)
			} else {
				fmt.Println()
			}
		},
	)

	t := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("How are you?"),
		),
	)

	err := antropicllm.Generate(context.Background(), t)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
}
