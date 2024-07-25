// models/product.go

package models

import "time"

type Product struct {
	Pno      int       `json:"pno"`
	Cate     string    `json:"cate"`
	Pname    string    `json:"pname"`
	Pcontent string    `json:"pcontent"`
	Img1     string    `json:"img1"`
	Img2     string    `json:"img2"`
	Img3     string    `json:"img3"`
	ResDate  time.Time `json:"resdate"`
	Hits     int       `json:"hits"`
}
