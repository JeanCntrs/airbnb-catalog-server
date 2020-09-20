package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/airbnb-catalog-server/db"
)

// ItemList : Get a list with the items
func ItemList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	tempPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Must send the page parameter as an integer", http.StatusBadRequest)
		return
	}

	pag := int64(tempPage)

	result, status := db.ReadAllItems(pag, search)
	if status == false {
		http.Error(w, "An error occurred while trying to read items", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
