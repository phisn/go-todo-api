package domain

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
