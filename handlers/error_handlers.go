package handlers

import (
	"groupie/models"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	headerWritten := false

	// If we haven't already written the header, do so
	if w.Header().Get("Content-Type") == "" {
		w.WriteHeader(statusCode)
		headerWritten = true
	}
	template, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// If we failed to load the template, we should write an error response
		if !headerWritten {
			http.Error(w, "Could not load error template", http.StatusInternalServerError)
		}
		return
	}

	customError := models.Error{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
	}

	if err := template.Execute(w, customError); err != nil {
		if !headerWritten {
			http.Error(w, "Could not render error template", http.StatusInternalServerError)
		}
	}

}
