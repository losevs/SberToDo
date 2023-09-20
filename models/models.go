package models

import "time"

type ToDo struct {
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Date        time.Time `json:"date" gorm:"type:timestamp"`
	Flag        bool      `json:"flag"`
}

type ToDoRequest struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Date        string `json:"date"`
	Flag        bool   `json:"flag"`
}
