package res

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

// ArtistsArray represents an array of artists
type ArtistsArray []Artist

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

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
