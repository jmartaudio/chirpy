package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmartaudio/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}
	recivedAPI, err := auth.GetAPIKey(r.Header)
	if err != nil || recivedAPI != cfg.polka_key {
		respondWithError(w, http.StatusUnauthorized, "Invalid API Key", err)
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	user, err := cfg.db.GetUserByID(r.Context(), params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find user", err)
		return
	}

	err = cfg.db.IsChirpyRed(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not set red status", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
