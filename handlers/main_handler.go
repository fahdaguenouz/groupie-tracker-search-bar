package handlers

import (
	"groupie/models"
	"html/template"
	"log"
	"net/http"
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

	data := struct {
		Artists      []models.Artist
		SearchResult struct {
			Artists []models.Artist
		}
	}{
		Artists: artists,
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
