package models

type ToDo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	DueDate   string `json:"due_date"`
	UserID    string `json:"user_id"`
}
