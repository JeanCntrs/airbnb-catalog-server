package main

import (
	"log"

	"github.com/JeanCntrs/airbnb-catalog-server/db"
	"github.com/JeanCntrs/airbnb-catalog-server/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Without connection to database")
		return
	}

	handlers.RouteHandlers()
}
