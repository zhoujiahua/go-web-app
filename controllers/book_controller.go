package controllers

import (
	"strconv"

	"restful/jerry/go-web-app/models"
	"restful/jerry/go-web-app/utils"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Router *gin.RouterGroup
}

func NewBookController(router *gin.Engine) *BookController {
	bookController := &BookController{}
	bookController.Router = router.Group("/api/v1/books")
	return bookController
}

// 获取所有书籍
func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		utils.JSONResponse(c, -1, err.Error(), nil)
		return
	}
	utils.JSONResponse(c, 0, "success", books)
}

// 获取单个书籍
func (bc *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		utils.JSONResponse(c, -2, "Invalid book ID", nil)
		return
	}
	book, err := models.GetBookByID(bookID)
	if err != nil {
		utils.JSONResponse(c, -3, "Book not found", nil)
		return
	}
	utils.JSONResponse(c, 0, "success", book)
}

// 添加书籍
func (bc *BookController) AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.JSONResponse(c, -2, err.Error(), nil)
		return
	}
	err := models.AddBook(&book)
	if err != nil {
		utils.JSONResponse(c, -1, err.Error(), nil)
		return
	}
	utils.JSONResponse(c, 0, "success", book)
}

// 修改书籍
func (bc *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		utils.JSONResponse(c, -2, "Invalid book ID", nil)
		return
	}
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.JSONResponse(c, -2, err.Error(), nil)
		return
	}
	book.ID = bookID
	err = models.UpdateBook(&book)
	if err != nil {
		utils.JSONResponse(c, -1, err.Error(), nil)
		return
	}
	utils.JSONResponse(c, 0, "success", book)
}

// 删除书籍
func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		utils.JSONResponse(c, -2, "Invalid book ID", nil)
		return
	}
	err = models.DeleteBook(bookID)
	if err != nil {
		utils.JSONResponse(c, -1, err.Error(), nil)
		return
	}
	utils.JSONResponse(c, 0, "Book deleted", nil)
}
