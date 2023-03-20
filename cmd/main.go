package main

import (
	"helloGo/internal/api"
	"log"
	"net/http"
)

// appHandler is a Top Level Exception Handler to return 500 on global exception
func main() {
	http.HandleFunc("/apples/{id}", api.AppHandler(api.HandleApples).CustomServeHttp)
	http.HandleFunc("/trees", api.AppHandler(api.HandleTrees).CustomServeHttp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
