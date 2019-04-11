package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/home/raj/Documents/code/go/src/templates/*")
	router.GET("/index", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html",gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")

}
