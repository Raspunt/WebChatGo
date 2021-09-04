package routers

import (
	db "BackendSide/Database"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func ServerStart() {

	r := gin.Default()
	r.Use(location.Default())

	// set path to html and css

	r.LoadHTMLGlob("frontend/templates/*.html")
	r.Static("/static", "./frontend/static")

	r.GET("/", func(c *gin.Context) {
		mapMessages := db.GetAllMessages()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"list": mapMessages,
		})
	})

	r.POST("/", func(c *gin.Context) {
		textFromUser := c.PostForm("textInput")
		// url := location.Get(c)
		db.CreateMessage(0, c.Request.RemoteAddr, textFromUser)

		mapMessages := db.GetAllMessages()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"list": mapMessages,
		})
	})

	// r.GET("/login", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.html", gin.H{
	// 		"list": "lol",
	// 	})
	// })

	// r.GET("/reg", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "register.html", gin.H{
	// 		"list": "lol",
	// 	})
	// })

	r.Run(":5000")
}
