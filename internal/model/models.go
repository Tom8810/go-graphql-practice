package gormModel

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Name      string    `json:"name"`
	Books     []*Book   `json:"books" gorm:"foreignkey:AuthorID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Book struct {
	ID        uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Title     string    `json:"title"`
	AuthorID  uint    	`json:"authorId" gorm:"index"`
	Author    *Author   `json:"author" gorm:"foreignkey:AuthorID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateAuthorInput struct {
	Name string `json:"name"`
}

type CreateBookInput struct {
	Title    string `json:"title"`
	AuthorID uint `json:"authorId"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateBookInput struct {
	Title    string `json:"title"`
	AuthorID uint `json:"authorId"`
}
