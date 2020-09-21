package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thimi0412/go-books/service"
)

type BookController struct{}

func (bc BookController) Index(c *gin.Context) {
	title := c.Query("title")
	category, _ := strconv.Atoi(c.Query("category"))
	author := c.Query("author")

	var bs service.BookService
	p, err := bs.Search(title, category, author)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

func (bc BookController) Create(c *gin.Context) {
	var json service.Parameter
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		title := json.Title
		category := json.Category
		author := json.Author

		var bs service.BookService
		if p, err := bs.Create(title, category, author); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, p)
		}
	}
}
