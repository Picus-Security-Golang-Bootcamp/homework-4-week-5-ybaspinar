package books

import (
	"gorm.io/gorm"
	"time"
)

import "encoding/json"

type Books []Book

func UnmarshalBooks(data []byte) (Books, error) {
	var r Books
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Book) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Book struct {
	BookID      uint           `json:"bookid",gorm:"primaryKey,unique"`
	Booktitle   string         `json:"booktitle"`
	Pages       int64          `json:"pages"`
	Stockamount int64          `json:"stockamount"`
	Price       float64        `json:"price"`
	Stockid     int64          `json:"stockid"`
	Isbn        int64          `json:"ısbn"`
	AuthorID    uint           `json:"-",gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-",gorm:"index"`
	Author      Author         `gorm:"-"`
}

type Author struct {
	AuthorID   uint           `json:"authorid",gorm:"primaryKey,unique"`
	Authorname string         `json:"authorname"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-",gorm:"index"`
}
