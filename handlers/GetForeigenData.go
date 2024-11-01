package handlers

import (
	"groupie/controllers"
	"groupie/models"
)

func GetForeigenData(data *models.Artist) error {

	var locations models.Location
	var date models.Date
	var relation models.Relation // Assuming Location is a struct for parsed data
	// Assuming Location is a struct for parsed data
	err := controllers.FetchData(data.Locations, &locations) // Fetching data for locations
	if err != nil {
		return err
	}
	err2 := controllers.FetchData(locations.Dates, &date) // Fetching data for locations
	if err2 != nil {
		return err2
	}
	err3 := controllers.FetchData(data.Relations, &relation) // Fetching data for locations
	if err3 != nil {
		return err3
	}

	data.Loca = locations
	data.Loca.AllDates = date
	data.Rela = relation
	return nil
}
