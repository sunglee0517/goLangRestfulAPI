// models/qna.go

package models

import (
	"database/sql"
	"time"
)

type Qna struct {
	Qno     int           `json:"qno"`
	Lev     int           `json:"lev"`
	Parno   sql.NullInt64 `json:"parno"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Author  string        `json:"author"`
	Resdate time.Time     `json:"resdate"`
	Hits    int           `json:"hits"`
}
