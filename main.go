package main

import (
	"log"
	"net/http"
	"os"

	"MyTravel/api"
)

func main() {
	// Koneksi DB (Local)
	api.ConnectDB()

	// Gunakan NewRouter (bukan Handler lagi)
	http.Handle("/", api.NewRouter())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("MyTravel backend running on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}