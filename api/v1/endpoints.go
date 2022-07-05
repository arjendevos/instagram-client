package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/arjendevos/instagram-client/models/responses"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("profile.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload responses.Profile
	err = json.Unmarshal(data, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
