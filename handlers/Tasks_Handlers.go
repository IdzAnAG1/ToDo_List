package handlers

import (
	"ToDo_List_v1/Structures"
	"ToDo_List_v1/Variables"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			err := json.NewEncoder(w).Encode(Variables.Tasks)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				_ = fmt.Errorf("Ошибка при получении таски из задачи %v ", err)
			}
		}
	case http.MethodPost:
		{
			var task Structures.Task
			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				_ = fmt.Errorf("Ивалидный инпут %v ", err)
				return
			}
			task.ID = Variables.AllTaskID
			Variables.AllTaskID++
			Variables.Tasks = append(Variables.Tasks, task)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(task)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
