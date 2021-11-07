package valorank

import (
	"devpex.com/valorank/pkg/rank"
	"fmt"
	"html"
	"log"
	"net/http"
)

// Valorank gets the rank sentence and writes it to the body
func Valorank(w http.ResponseWriter, r *http.Request) {
	valorank, err := rank.GetRank()
	if err != nil {
		log.Printf("Valorank: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, html.EscapeString(valorank.Sentence))
}
