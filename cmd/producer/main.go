package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	EventsProcessed int
}

type CarbonFree struct {
	NumberOfTrees int `json:"numberOfTrees"`
}

type PlantedTreeEvent struct {
	Id int `json:"id"`
}

func plantTree(numberOfTrees int, handler *Handler) {

	// Publish events using Dapr pubsub
	for i := 1; i <= numberOfTrees; i++ {

		event := PlantedTreeEvent{Id: handler.EventsProcessed}
		log.Println("Event Published - Id:", event.Id)
		handler.EventsProcessed++
	}
}

func main() {
	log.Println("Producer App Started ...")

	handler := &Handler{EventsProcessed: 1}
	http.Handle("/plant", handler)
	http.ListenAndServe(":8081", nil)
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	var carbon CarbonFree
	_ = json.NewDecoder(request.Body).Decode(&carbon)
	plantTree(carbon.NumberOfTrees, handler)

}
