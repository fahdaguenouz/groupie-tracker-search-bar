// handlers/artists.go
package handlers

import (
	"groupie/models"
	"groupie/controllers"
)

// FetchArtists fetches the list of all artists from the API.
func FetchArtists() ([]models.Artist, error) {
	var artists []models.Artist
	url := "https://groupietrackers.herokuapp.com/api/artists"

	if err := controllers.FetchData(url, &artists); err != nil {
		return nil, err
	}

	return artists, nil
}
