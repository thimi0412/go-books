package main

import (
	"github.com/thimi0412/go-books/db"
	"github.com/thimi0412/go-books/server"
)

func main() {
	db.Init()
	server.Init()
}
