package main

import (
	"fmt"

	"github.com/jwfriese/lingoose/legacy/prompt"
)

func main() {

	prompt1 := prompt.New("Hello World")
	fmt.Println(prompt1.String())

}
