package res

import "fmt"

// Endpoints represents the API endpoints
type Endpoints struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Coordinate struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Location string  `json:"location"`
}

// ArtistsArray represents an array of artists
type ArtistsArray []Artist

// Relation represents a relation between dates and locations
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// GetArtists returns a list of artists
func GetArtists() (*ArtistsArray, error) {
	// Set the endpoint to the artists endpoint
	endpoint := endpointArtists

	// Set the options map with the endpoint
	options := map[string]interface{}{
		"endpoint": endpoint,
	}

	// Create a new ArtistsArray
	artists := new(ArtistsArray)

	// Make the API request using the makePetition function
	err := makePetition(options, &artists)
	if err != nil {
		// Return an empty ArtistsArray and the error if the request fails
		return &ArtistsArray{}, err
	}

	// Return the artists array
	return artists, nil
}

func GetRelation(number int) (*Relation, error) {
	endpoint := endpointRelation

	options := map[string]interface{}{
		"endpoint": endpoint,
		"params": map[string]int{
			"integer": number,
		},
	}

	relation := new(Relation)

	err := makePetition(options, &relation)
	if err != nil {
		return &Relation{}, err
	}

	formatConcerts(&relation)

	return relation, nil
}

func GetCoordinates(artistNumber int) []Coordinate {
	endpoint := endpointRelation

	options := map[string]interface{}{
		"endpoint": endpoint,
		"params": map[string]int{
			"integer": artistNumber,
		},
	}
	relation := new(Relation)

	makePetition(options, &relation)

	updMap := createUpdMap(&relation)

	fmt.Print(updMap)
	return addressToGeoCode(updMap)
}
