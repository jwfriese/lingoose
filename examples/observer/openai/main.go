package main

import (
	"context"

	"github.com/jwfriese/lingoose/llm/openai"
	"github.com/jwfriese/lingoose/observer"
	"github.com/jwfriese/lingoose/observer/langfuse"
	"github.com/jwfriese/lingoose/thread"
)

func main() {
	ctx := context.Background()

	o := langfuse.New(ctx)
	trace, err := o.Trace(&observer.Trace{Name: "Who are you"})
	if err != nil {
		panic(err)
	}

	ctx = observer.ContextWithObserverInstance(ctx, o)
	ctx = observer.ContextWithTraceID(ctx, trace.ID)

	span, err := o.Span(
		&observer.Span{
			TraceID: trace.ID,
			Name:    "SPAN",
		},
	)
	if err != nil {
		panic(err)
	}

	ctx = observer.ContextWithParentID(ctx, span.ID)

	openaillm := openai.New()

	t := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("Hello, who are you?"),
		),
	)

	err = openaillm.Generate(ctx, t)
	if err != nil {
		panic(err)
	}

	o.Flush(ctx)
}
