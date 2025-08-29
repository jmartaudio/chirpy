package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidChirp(w http.ResponseWriter, r *http.Request) {
	type incJsonShape struct {
		Body string `json:"body"`
	}

	type rtnJsonShape struct {
		CleanedBody string `json:"cleaned_body"`
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
		CleanedBody: cleanChirp(params.Body),
	})
}

func cleanChirp(chirpBody string) string {
	naughty := make(map[string]struct{})
	naughty["kerfuffle"] = struct{}{}
	naughty["sharbert"] = struct{}{}
	naughty["fornax"] = struct{}{}

	words := strings.Split(chirpBody, " ")

	for i, word := range words {
		if _, ok := naughty[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}
	cleanChirpBody := strings.Join(words, " ")

	return cleanChirpBody
}
