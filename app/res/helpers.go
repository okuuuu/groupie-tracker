package res

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type GeocodeResponse struct {
	Results []struct {
		Locations []struct {
			LatLng struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"latLng"`
		} `json:"locations"`
	} `json:"results"`
}

func capitalize(s string) string {
	rns := []rune(s)
	for i, r := range rns {
		if i == 0 || s[i-1] == ' ' {
			rns[i] = r - 32
		}
	}
	res := string(rns)
	res = strings.ReplaceAll(res, "Usa", "USA")
	res = strings.ReplaceAll(res, "Uk", "UK")
	return res
}

func addressToGeoCode(updmap map[string][]string) []Coordinate {
	coordinates := []Coordinate{}
	//fmt.Print(updmap)
	keys := make([]string, 0, len(updmap))
	for k := range updmap {
		keys = append(keys, k)
	}
	//fmt.Print(keys)
	for _, key := range keys {
		apiKey := "cAWwy1hW2pUfGWvBK1Z18XRSPCAKbfEA"
		location := key
		baseURL := "http://www.mapquestapi.com/geocoding/v1/address"

		// Build the query parameters for the API request
		queryParams := url.Values{}
		queryParams.Add("key", apiKey)
		queryParams.Add("location", location)

		// Construct the full API URL with the base URL and query parameters
		apiURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

		// Make a GET request to the API
		resp, err := http.Get(apiURL)
		if err != nil {
			// Handle error
			fmt.Println(err)

		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// Handle error
			fmt.Println(err)

		}

		// Unmarshal the JSON response into a GeocodeResponse struct
		var geocodeResponse GeocodeResponse
		err = json.Unmarshal(body, &geocodeResponse)
		if err != nil {
			// Handle error
			fmt.Println(err)

		}

		// Extract the latitude and longitude from the response
		lat := geocodeResponse.Results[0].Locations[0].LatLng.Lat
		lng := geocodeResponse.Results[0].Locations[0].LatLng.Lng

		coordinates = append(coordinates, Coordinate{lat, lng, location})
	}
	return coordinates
}

func formatConcerts(relation **Relation) {
	(*relation).DatesLocations = createUpdMap(relation)
}

func createUpdMap(relation **Relation) map[string][]string {
	updmap := make(map[string][]string)
	for key, val := range (*relation).DatesLocations {
		key = strings.ReplaceAll(key, "_", " ")
		key = strings.ReplaceAll(key, "-", ", ")
		key = capitalize(key)
		updmap[key] = val
	}
	return updmap
}
