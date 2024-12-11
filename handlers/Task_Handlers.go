package handlers

import (
	"ToDo_List_v1/Structures"
	"ToDo_List_v1/Variables"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HandleTaskByUD(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr[0:])
	if err != nil {
		fmt.Println(idStr)
		http.Error(w, "Инвалидный ID ", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		{
			for _, task := range Variables.Tasks {
				if task.ID == id {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(task)
					return
				}
			}
			http.Error(w, "Таска не найдена", http.StatusNotFound)
		}
	case http.MethodPut:
		var updatedTask Structures.Task
		err := json.NewDecoder(r.Body).Decode(&updatedTask)
		if err != nil {
			http.Error(w, "Инвалидный инпут", http.StatusBadRequest)
			return
		}
		for i, task := range Variables.Tasks {
			if task.ID == id {
				Variables.Tasks[i].Name = updatedTask.Name
				Variables.Tasks[i].Done = updatedTask.Done
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(Variables.Tasks[i])
				return
			}
		}
		http.Error(w, "Таска не найдена", http.StatusNotFound)
	case http.MethodDelete:
		{
			for i, task := range Variables.Tasks {
				if task.ID == id {
					Variables.Tasks = append(Variables.Tasks[:i], Variables.Tasks[i+1:]...)
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			http.Error(w, "таска не найдена", http.StatusNoContent)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
