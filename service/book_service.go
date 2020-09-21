package service

import (
	"log"

	"github.com/thimi0412/go-books/db"
	"github.com/thimi0412/go-books/model"
)

type BookService struct{}

type Book model.Book

type Parameter struct {
	Title    string `json:"title"`
	Category int    `json:"category"`
	Author   string `json:"author"`
}

func (bs *BookService) Search(title string, category int, author string) ([]Book, error) {
	db := db.Connection()
	defer db.Close()

	var book []Book
	p := Parameter{title, category, author}

	tx := db.Find(&book)

	if p.Title != "" {
		tx = tx.Where("title = ?", p.Title).Find(&book)
	}

	if p.Category != 0 {
		tx = tx.Where("category = ?", p.Category).Find(&book)
	}

	if p.Author != "" {
		tx = tx.Where("author = ?", p.Author).Find(&book)
	}
	return book, nil
}

func (bs *BookService) Create(title string, category int, author string) (Book, error) {
	db := db.Connection()
	defer db.Close()

	book := Book{Title: title, Category: category, Author: author}

	if err := db.Create(&book).Error; err != nil {
		log.Println(err)
		return book, err
	}
	return book, nil
}