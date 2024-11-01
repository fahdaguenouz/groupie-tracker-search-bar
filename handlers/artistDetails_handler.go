package handlers

import (
	"groupie/controllers"

	"groupie/models"
	"net/http"
)

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/artists/"):] // Extract artist ID from URL
	var artist models.Artist

	err2 := controllers.FetchData("https://groupietrackers.herokuapp.com/api/artists/"+id, &artist)
	if err2 != nil {
		ErrorHandler(w, r, http.StatusNotFound)
		return 
	}
	// Check if artist was found (ID 0 indicates artist not found)
	if artist.ID == 0 {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	err := GetForeigenData(&artist)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	
	if err := renderTemplate(w, "artist.html", artist); err != nil {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}
