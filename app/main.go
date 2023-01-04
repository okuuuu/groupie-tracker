package main

import (
	"log"
	"main/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.MakeHandler(server.HomeHandler))

	log.Print("Starting server at port 8080...\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Print("Please follow this link: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
