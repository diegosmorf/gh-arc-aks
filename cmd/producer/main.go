
package main

import (
	"context"
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
	// Create a new client for Dapr using the SDK
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalln("Error to create instace of Dapr Client: ", err)
	}
	defer client.Close()

	handler := &Handler{EventsProcessed: 1}
	http.Handle("/plant", handler)
	http.ListenAndServe(":8081", nil)
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	var carbon CarbonFree
	_ = json.NewDecoder(request.Body).Decode(&carbon)
	plantTree(carbon.NumberOfTrees, handler)

}
