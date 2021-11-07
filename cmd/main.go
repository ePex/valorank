package main

import (
	"devpex.com/valorank/pkg/rank"
	"fmt"
	"html"
	"log"
)

func main() {
	valorank, err := rank.GetRank("eu", "mimimaly_", "6969")
	if err != nil {
		log.Printf("Valorank: %v", err)
		return
	}

	fmt.Printf(html.EscapeString(valorank.Sentence))
}
