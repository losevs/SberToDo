package models

type ToDo struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Date        string `json:"date"`
	Flag        bool   `json:"flag"`
}
