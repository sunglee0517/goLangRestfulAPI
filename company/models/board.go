// models/board.go

package models

import "time"

type Board struct {
	No      int       `json:"no"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	ResDate time.Time `json:"resdate"`
	Hits    int       `json:"hits"`
}
