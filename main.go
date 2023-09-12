package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		var time_start = time.Now()
		var list []string
		for i := 0; i < 8000; i++ {
			list = append(list, "Hello World")
		}
		var time_end = time.Now()
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"start":    time_start,
			"end":      time_end,
			"time_sec": time_end.Sub(time_start).Seconds(),
		})
	})

	r.Run(
		":8080", // listen and serve on
	) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
