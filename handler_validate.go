package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidChirp(w http.ResponseWriter, r *http.Request) {
	type incJsonShape struct {
		Body string `json:"body"`
	}

	type rtnJsonShape struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := incJsonShape{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, rtnJsonShape{
		Valid: true,
	})
}

func cleanChirp(string) string {
	// maybe this should be done with the whole json struct?
	// split words
	// lowercase words
	// match words against list
	// replace matched word with ****
	// return string
	return ""
}
