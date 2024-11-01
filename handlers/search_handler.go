package handlers

import (
	"groupie/models"
	"html/template"
	"net/http"
	"strings"
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
	queryLower := strings.ToLower(query)

	// Filter artists based on the query
	var filteredArtists []models.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), queryLower) {
			filteredArtists = append(filteredArtists, artist)
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), queryLower) {
				filteredArtists = append(filteredArtists, artist)
				break
			}
		}
		if strings.Contains(strings.ToLower(artist.Locations), queryLower) {
			filteredArtists = append(filteredArtists, artist)
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), queryLower) {
			filteredArtists = append(filteredArtists, artist)
		}
		if strings.Contains(strings.ToLower(string(artist.CreationDate)), queryLower) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	// Create a template data structure
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
