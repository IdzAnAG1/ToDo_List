package main

import (
	"ToDo_List_v1/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Дописать Task_Handler

	http.HandleFunc("/tasks", handlers.HandleTasks)
	http.HandleFunc("/tasks/", handlers.HandleTaskByUD)

	fmt.Println("Server was launched at the http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
