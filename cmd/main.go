package main

import (
	"fmt"
	"html"
	"log"
	"valorank/pkg/rank"
)

func main() {
	valorank, err := rank.GetRank()
	if err != nil {
		log.Printf("Valorank: %v", err)
		return
	}

	fmt.Printf(html.EscapeString(valorank.Sentence))
}
