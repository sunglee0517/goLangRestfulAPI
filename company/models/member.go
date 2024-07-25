// models/member.go

package models

import "time"

type Member struct {
	ID       string    `json:"id"`
	PW       string    `json:"pw"`
	Name     string    `json:"name"`
	Birth    time.Time `json:"birth"`
	Email    string    `json:"email"`
	Tel      string    `json:"tel"`
	Addr1    string    `json:"addr1"`
	Addr2    string    `json:"addr2"`
	Postcode string    `json:"postcode"`
	RegDate  time.Time `json:"regdate"`
}
