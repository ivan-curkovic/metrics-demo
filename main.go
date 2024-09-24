package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	r.Use(p.Instrument())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
