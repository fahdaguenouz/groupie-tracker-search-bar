package handlers

import (
	"groupie/models"
	"html/template"
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
		Artists []models.Artist
	}{
		Artists: artists,
	}
	// Render the main template
	if err := renderTemplate(w, "index.html", data); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		// Log the error for debugging
		// log.Printf("Error parsing template %s: %v", templateName, err)
		return err // Return error instead of writing to response
	}

	if err := tmpl.Execute(w, data); err != nil {
		// Log the error for debugging
		// log.Printf("Error executing template %s: %v", templateName, err)
		return err
	}

	return nil
}
