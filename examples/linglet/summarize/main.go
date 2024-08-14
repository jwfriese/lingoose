package main

import (
	"context"
	"fmt"

	"github.com/jwfriese/lingoose/linglet/summarize"
	"github.com/jwfriese/lingoose/llm/openai"
	"github.com/jwfriese/lingoose/loader"
	"github.com/jwfriese/lingoose/textsplitter"
	"github.com/jwfriese/lingoose/thread"
)

// download https://raw.githubusercontent.com/hwchase17/chat-your-data/master/state_of_the_union.txt

func main() {

	textLoader := loader.NewTextLoader("state_of_the_union.txt", nil).
		WithTextSplitter(textsplitter.NewRecursiveCharacterTextSplitter(4000, 0))

	summarize := summarize.New(
		openai.New().WithMaxTokens(2000).WithTemperature(0).WithModel(openai.GPT3Dot5Turbo16K0613),
		textLoader,
	).WithCallback(
		func(t *thread.Thread, i, n int) {
			fmt.Printf("Progress : %.0f%%\n", float64(i)/float64(n)*100)
		},
	)

	summary, err := summarize.Run(context.Background())
	if err != nil {
		panic(err)
	}

	println(*summary)
}
