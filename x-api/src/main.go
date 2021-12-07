package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func init() {
	// db := config.Init()
	// migration.Migrate(db)
}


func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}