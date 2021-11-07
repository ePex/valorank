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
	region := r.URL.Query().Get("region")
	name := r.URL.Query().Get("name")
	tagline := r.URL.Query().Get("tagline")
	valorank, err := rank.GetRank(region, name, tagline)
	if err != nil {
		log.Printf("Valorank: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, html.EscapeString(valorank.Sentence))
}
