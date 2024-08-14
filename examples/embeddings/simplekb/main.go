package main

import (
	"context"

	openaiembedder "github.com/jwfriese/lingoose/embedder/openai"
	"github.com/jwfriese/lingoose/index"
	"github.com/jwfriese/lingoose/index/option"
	"github.com/jwfriese/lingoose/index/vectordb/jsondb"
	qapipeline "github.com/jwfriese/lingoose/legacy/pipeline/qa"
	"github.com/jwfriese/lingoose/llm/openai"
	"github.com/jwfriese/lingoose/loader"
	"github.com/jwfriese/lingoose/textsplitter"
)

func main() {
	docs, _ := loader.NewPDFToTextLoader("./kb").WithTextSplitter(textsplitter.NewRecursiveCharacterTextSplitter(2000, 200)).Load(context.Background())
	index := index.New(jsondb.New(), openaiembedder.New(openaiembedder.AdaEmbeddingV2)).WithIncludeContents(true)
	index.LoadFromDocuments(context.Background(), docs)
	qapipeline.New(openai.NewChat().WithVerbose(true)).WithIndex(index).Query(context.Background(), "What is the NATO purpose?", option.WithTopK(1))
}
