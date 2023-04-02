package models

import (
	"errors"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookService interface {
	GetAllBooks() []Book
	GetBookByID(id int) (Book, error)
	AddBook(book Book) Book
	UpdateBook(id int, book Book) (Book, error)
	DeleteBook(id int) error
}

var books = []Book{
	{ID: 1, Title: "Book 1", Author: "Author 1"},
	{ID: 2, Title: "Book 2", Author: "Author 2"},
	{ID: 3, Title: "Book 3", Author: "Author 3"},
}

func GetAllBooks() ([]Book, error) {
	return books, nil
}

func GetBookByID(id int) (Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, errors.New("Book not found")
}

func AddBook(book *Book) error {
	book.ID = len(books) + 1
	books = append(books, *book)
	return nil
}

func UpdateBook(book *Book) error {
	for i, b := range books {
		if b.ID == book.ID {
			books[i] = *book
			return nil
		}
	}
	return errors.New("Book not found")
}

func DeleteBook(id int) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("Book not found")
}
