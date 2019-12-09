package models

import "time"

type RequestTodoModel struct {
	Title     string    `json:"title"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type ResponseTodoModel struct {
	ID        int64     `json:"id",omitempty`
	Title     string    `json:"title",omitempty`
	Status    string    `json:"status",omitempty`
	CreatedBy string    `json:"createdBy",omitempty`
	CreatedAt time.Time `json:"createdAt",omitempty`
}

type UpdateTodoModel struct {
	NewTitle *string `json:"newTitle",omitempty`
	Status   *string `json:"status",omitempty`
}
