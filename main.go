package main

import (
	"fmt"
	"net/http"

	"groupie/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	http.HandleFunc("/assets/", handlers.AssetsHandler)
	http.HandleFunc("/artists/", handlers.ArtistDetailHandler)
	fmt.Println("Server is running at http://localhost:3001")

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
