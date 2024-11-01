package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"` // Ensure this is a slice of Date structs
	Relations    string   `json:"relations"`    // Assuming relations is a map
	Loca         Location
	Rela   Relation
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Location represents a location with associated dates.
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
	AllDates  Date
}

// Date represents a date associated with an event.
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}


type SearchResult struct {
	Artists []Artist // List of artists matching the search criteria
	Found   bool     // Indicates if any results were found
}