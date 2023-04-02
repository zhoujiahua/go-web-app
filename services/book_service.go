package services

import (
	"errors"

	"restful/jerry/go-web-app/models"
)

type BookServiceImpl struct {
	books []models.Book
}

func NewBookService() models.BookService {
	return &BookServiceImpl{
		books: []models.Book{
			{ID: 1, Title: "Book 1", Author: "Author 1"},
			{ID: 2, Title: "Book 2", Author: "Author 2"},
			{ID: 3, Title: "Book 3", Author: "Author 3"},
		},
	}
}

func (s *BookServiceImpl) GetAllBooks() []models.Book {
	return s.books
}

func (s *BookServiceImpl) GetBookByID(id int) (models.Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("Book not found")
}

func (s *BookServiceImpl) AddBook(book models.Book) models.Book {
	book.ID = len(s.books) + 1
	s.books = append(s.books, book)
	return book
}

func (s *BookServiceImpl) UpdateBook(id int, book models.Book) (models.Book, error) {
	for i, b := range s.books {
		if b.ID == id {
			book.ID = b.ID
			s.books[i] = book
			return book, nil
		}
	}
	return models.Book{}, errors.New("Book not found")
}

func (s *BookServiceImpl) DeleteBook(id int) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return errors.New("Book not found")
}
