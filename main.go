package main

import (
	routers "OneGO/urls"
	"log"
	"net/http"
)

func main() {
	router := routers.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
