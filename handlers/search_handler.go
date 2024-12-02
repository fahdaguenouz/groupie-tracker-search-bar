package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupie/models"
)

// SearchHandler processes the search request.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	// Fetch all artists
	artists, err := FetchArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get search query
	query := r.FormValue("q")
	query=strings.ReplaceAll(query,", ","-")
	queryLower := strings.ToLower(query)
	
	// Filter artists based on the query

	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect shared access to filteredArtists
	filteredArtists := make([]models.Artist, 0)

	for _, artist := range artists {
		wg.Add(1) // Increment the WaitGroup counter

		go func(artist models.Artist) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			// Fetch foreign data (including locations) for each artist
			if err := GetForeigenData(&artist); err != nil {
				fmt.Println("Error fetching foreign data:", err)
				return // Skip this artist if there's an error
			}

			mu.Lock() // Lock the mutex before accessing shared data
			defer mu.Unlock()

			// Check if artist name matches the query
			if strings.Contains(strings.ToLower(artist.Name), queryLower) {
				filteredArtists = append(filteredArtists, artist)
				return // Skip to the next artist
			}

			// Check if any member matches the query
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), queryLower) {
					filteredArtists = append(filteredArtists, artist)
					return // No need to check further members
				}
			}

			// Check if FirstAlbum matches the query
			if strings.Contains(strings.ToLower(artist.FirstAlbum), queryLower) {
				filteredArtists = append(filteredArtists, artist)
				return // Skip to the next artist
			}

			// Check if CreationDate matches the query (ensure it's an integer)
			if creationDate, err := strconv.Atoi(query); err == nil && artist.CreationDate == creationDate {
				filteredArtists = append(filteredArtists, artist)
				return // Skip to the next artist
			}

			// Check if any location matches the query after fetching locations
			for _, locat := range artist.Loca.Locations {
				if strings.Contains(strings.ToLower(locat), queryLower) {
					filteredArtists = append(filteredArtists, artist)
					return // No need to check further locations
				}
			}
		}(artist) // Pass a copy of the current artist to the goroutine
	}

	wg.Wait() // Wait for all goroutines to finish

	// var filteredArtists []models.Artist

	// for _, artist := range artists {
	// 	// Fetch foreign data (including locations) for each artist
	// 	if err := GetForeigenData(&artist); err != nil {
	// 		fmt.Println("Error fetching foreign data:", err)
	// 		continue // Skip this artist if there's an error
	// 	}

	// 	// Check if artist name matches the query
	// 	if strings.Contains(strings.ToLower(artist.Name), queryLower) {
	// 		filteredArtists = append(filteredArtists, artist)
	// 		continue // Skip to the next artist
	// 	}

	// 	// Check if any member matches the query
	// 	for _, member := range artist.Members {
	// 		if strings.Contains(strings.ToLower(member), queryLower) {
	// 			filteredArtists = append(filteredArtists, artist)
	// 			break // No need to check further members
	// 		}
	// 	}

	// 	// Check if FirstAlbum matches the query
	// 	if strings.Contains(strings.ToLower(artist.FirstAlbum), queryLower) {
	// 		filteredArtists = append(filteredArtists, artist)
	// 		continue // Skip to the next artist
	// 	}

	// 	// Check if CreationDate matches the query (ensure it's an integer)
	// 	if creationDate, err := strconv.Atoi(query); err == nil && artist.CreationDate == creationDate {
	// 		filteredArtists = append(filteredArtists, artist)
	// 		continue // Skip to the next artist
	// 	}

	// 	// Check if any location matches the query after fetching locations
	// 	for _, locat := range artist.Loca.Locations {
	// 		if strings.Contains(strings.ToLower(locat), queryLower) {
	// 			filteredArtists = append(filteredArtists, artist)
	// 			break // No need to check further locations
	// 		}
	// 	}
	// }

	data := struct {
		Artists []models.Artist
		Query   string
	}{
		Artists: filteredArtists,
		Query:   query,
	}

	// Render the search results template
	template, err := template.ParseFiles("templates/search.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if err := template.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
