package rank

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Status int `json:"status"`
	Data   struct {
		CurrentTier         string `json:"currenttier"`
		CurrentTierPatched  string `json:"currenttierpatched"`
		RankingInTier       int    `json:"ranking_in_tier"`
		MmrChangeToLastGame int    `json:"mmr_change_to_last_game"`
		Elo                 int    `json:"elo"`
	} `json:"data"`
}

type Rank struct {
	Sentence string
}

// GetRank calls the api https://api.henrikdev.xyz/valorant/v1/mmr/eu/devpex/EUW and transforms
// the fields into a human readable sentence.
func GetRank(region string, name string, tagline string) (Rank, error) {
	resp, err := http.Get("https://api.henrikdev.xyz/valorant/v1/mmr/" + region + "/" + name + "/" + tagline)
	if err != nil {
		log.Printf("http.NewRequest: %v", err)
		return Rank{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		return Rank{}, err
	}

	var response Response
	json.Unmarshal(data, &response)

	rank := Rank{
		Sentence: name + " ist momentan " + response.Data.CurrentTierPatched + " mit " + strconv.Itoa(response.Data.RankingInTier) + " punkten. " + name + " hat im letzten Spiel " + strconv.Itoa(response.Data.MmrChangeToLastGame) + " Punkte gemacht.",
	}
	return rank, nil
}
