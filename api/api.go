package api

import (
	"net/http"

	v1 "github.com/arjendevos/instagram-client/api/v1"
	h "github.com/arjendevos/instagram-client/helpers"
	"github.com/rs/zerolog/log"
)

func Start() {
	port := ":10000"
	http.HandleFunc("/api/v1/profile", v1.Profile)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		h.Handle(err)
	}
	log.Info().Str("SERVER RUNNING ON PORT", port).Msg("Server running!")
}
