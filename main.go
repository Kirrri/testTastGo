package main

import (
	"OneGO/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
