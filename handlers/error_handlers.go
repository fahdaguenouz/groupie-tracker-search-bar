package handlers

import (
	"groupie/models"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	// Check if the header has already been sent
	if w.Header().Get("Content-Type") != "" {
		return // Avoid writing the header again
	}

	// Set the response header
	w.WriteHeader(statusCode)

	template, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Could not load error template", http.StatusInternalServerError)
		return
	}

	customError := models.Error{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
	}

	if err := template.Execute(w, customError); err != nil {
		http.Error(w, "Could not render error template", http.StatusInternalServerError)
	}

}
