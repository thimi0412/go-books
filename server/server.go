package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thimi0412/go-books/controller"
)

func Init() {
	r := router()

	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	r.Use(CORS())
	u := r.Group("/book")
	{
		ctrl := controller.BookController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
	}
	return r

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
