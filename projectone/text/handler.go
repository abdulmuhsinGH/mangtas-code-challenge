package text

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	//get the passage from the request
	passage := r.FormValue("text")

	//return empty json if no passage is passed
	if len(passage) == 0 {
		log.Println("no text passed")
		http.Error(w, "no text passed", http.StatusBadRequest)
		return
	}

	//call search function
	words := Search(passage)

	//write the response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
