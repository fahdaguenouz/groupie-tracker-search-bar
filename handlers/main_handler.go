package handlers

import (
	"html/template"
	"log"
	"net/http"

	"groupie/models"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	artists, err := FetchArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	// Prepare to collect unique locations
	locationSet := make(map[string]struct{}) // To store unique locations

	for _, artist := range artists {
		if err := GetForeigenData(&artist); err == nil { // Fetch foreign data including locations
			for _, loc := range artist.Loca.Locations {
				locationSet[loc] = struct{}{} // Collect unique locations
			}
		}
	}

	// Convert locationSet to a slice for passing to the template
	var locations []string
	for loc := range locationSet {
		locations = append(locations, loc)
	}

	data := struct {
		Artists   []models.Artist
		Locations []string // Add Locations field here
	}{
		Artists:   artists,
		Locations: locations,
	}
	// Render the main template
	if err := renderTemplate(w, "index.html", data); err != nil {
		log.Printf("Error rendering template: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		// Return the error to be handled by the calling function
		return err
	}

	if err := tmpl.Execute(w, data); err != nil {
		// Return the error to be handled by the calling function
		return err
	}

	return nil
}
