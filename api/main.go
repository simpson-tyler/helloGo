package api

import (
	"log"
	"net/http"
)

// appHandler is a Top Level Exception Handler to return 500 on global exception
func main() {
	http.Handle("/apples/{id}", appHandler(handleApples))
	http.Handle("/trees", appHandler(handleTrees))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
