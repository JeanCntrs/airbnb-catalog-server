package middlewares

import (
	"net/http"

	"github.com/JeanCntrs/airbnb-catalog-server/db"
)

// CheckDB : Check the status of database and if everything is fine, continue the process, otherwise, stop the execution
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost connection to database", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
