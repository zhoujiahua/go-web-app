package models

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
