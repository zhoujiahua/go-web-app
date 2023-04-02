package models

import "fmt"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "Book 1", Author: "Author 1"},
	{ID: 2, Title: "Book 2", Author: "Author 2"},
	{ID: 3, Title: "Book 3", Author: "Author 3"},
}

func Init() {
	// 在此初始化
}

func GetAllBooks() []Book {
	return books
}

func GetBookByID(id int) (*Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}

func AddBook(book Book) error {
	book.ID = len(books) + 1
	books = append(books, book)
	return nil
}

func UpdateBook(id int, book Book) error {
	for i, b := range books {
		if b.ID == id {
			books[i] = book
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

func DeleteBook(id int) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}
