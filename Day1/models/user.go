package models

type Task struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descriptionn string `json:"descriptionn"`
	Statuss      string `json:"statuss"`
	Priority     string `json:"priority"`
	DueDate      string `json:"due_date"`
	CreatedAt    string `json:"created_at"`
}
