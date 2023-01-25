package main

import (
	"encoding/json"
	"log"
	"main/res"
	"main/server"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", server.MakeHandler(server.HomeHandler))

	http.HandleFunc("/coordinates/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/coordinates/")
		intVar, _ := strconv.Atoi(id)
		jsonCoords, _ := json.Marshal(res.GetCoordinates(intVar))

		// Set the content type to JSON
		w.Header().Set("Content-Type", "application/json")
		// Encode the coordinates as JSON and write them to the response body
		w.Write(jsonCoords)
	})

	log.Print("Starting server at port 8080...\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Print("Please follow this link: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
