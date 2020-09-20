package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JeanCntrs/airbnb-catalog-server/db"
)

// ItemDetail : Allows to extract the values of the item by id
func ItemDetail(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send id parameter", http.StatusBadRequest)
		return
	}

	item, err := db.ReadItem(ID)
	if err != nil {
		http.Error(w, "An error ocurred while trying to find the item "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
