// models/dataroom.go

package models

import "time"

type Dataroom struct {
	Dno      int       `json:"dno"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Datafile string    `json:"datafile"`
	ResDate  time.Time `json:"resdate"`
	Hits     int       `json:"hits"`
}
