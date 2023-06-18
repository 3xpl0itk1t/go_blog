package models

type PostModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
