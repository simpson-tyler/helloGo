package api

import (
	"log"
	"net/http"
)

// appHandler is a Top Level Exception Handler to return 500 on global exception
func main() {
	http.Handle("/apples/{id}", api.ServeHTTP(HandleApples))
	http.Handle("/trees", api.ServeHTTP(HandleTrees))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
