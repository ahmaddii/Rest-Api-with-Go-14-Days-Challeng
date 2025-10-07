package controllers

import (
	"Day1/database"
	"Day1/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// create a post or create function which post a task

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		http.Error(w, "Invalid Data from Body", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO tasks(title,descriptionn,statuss,priority,due_date,created_at) VALUES (?,?,?,?,?,?)"

	_, err = database.DB.Exec(query, task.Title, task.Descriptionn, task.Statuss, task.Priority, task.DueDate, task.CreatedAt)

	if err != nil {

		http.Error(w, "Data Insertion Failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) //created and posted
	json.NewEncoder(w).Encode(map[string]string{"message": "Task Created Successfully"})

}

// now get all tasks functions

func GetTasks(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id,title,descriptionn,statuss,priority,due_date,created_at FROM tasks")

	if err != nil {

		http.Error(w, "Error while getting data ", http.StatusInternalServerError)
		return
	}

	defer rows.Close() // close the row to prevent from data leak

	var tasks []models.Task // store the tasks in the slice of tasks

	for rows.Next() {

		var task models.Task

		err := rows.Scan(&task.Id, &task.Title, &task.Descriptionn, &task.Statuss, &task.Priority, &task.DueDate, &task.CreatedAt)

		if err != nil {

			http.Error(w, "Error while Scanning the data", http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)

	}

	w.Header().Set("Content-Type", "application/json") // data what the client shows
	json.NewEncoder(w).Encode(tasks)

}

// Now get Tasks by ID
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var task models.Task

	query := "SELECT id, title, descriptionn, statuss, priority, due_date, created_at FROM tasks WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(
		&task.Id,
		&task.Title,
		&task.Descriptionn,
		&task.Statuss,
		&task.Priority,
		&task.DueDate,
		&task.CreatedAt,
	)

	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// now for Updating specifc task and edit it

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		http.Error(w, "Invalid Data from Body", http.StatusBadRequest)

		return

	}

	query := `UPDATE tasks SET title = ?, descriptionn = ? , statuss = ? , priority = ? , due_date = ? , created_at = ? WHERE id = ?`

	_, err = database.DB.Exec(query, task.Title, task.Descriptionn, task.Statuss, task.Priority, task.DueDate, task.CreatedAt, id)

	if err != nil {

		http.Error(w, "Error While Updating Data", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task Updated Succesfully !"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	query := "DELETE FROM tasks WHERE id = ?"

	result, err := database.DB.Exec(query, id)

	if err != nil {

		http.Error(w, "Error while deleting a task ", http.StatusInternalServerError)
		return

	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {

		http.Error(w, "Error while checking deleted", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {

		http.Error(w, "Task not found", http.StatusNotFound)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task Deleted Successfully !"})

}
