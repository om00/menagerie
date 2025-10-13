package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/om00/menagerie/database"
)

func main() {

	appPort := os.Getenv("APP_PORT_INSIDE")
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatal("error while connectin to the database", err)
	}

	defer db.MainDB.Close()
	database.RunMigrations()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := db.MainDB.Ping(); err != nil {
			http.Error(w, "DB connection failed", 500)
			return
		}
		fmt.Fprintf(w, "✅ Connected to MySQL successfully as ")
	})

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	fmt.Printf("🚀 Server running on :%s\n", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
