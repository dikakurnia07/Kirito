package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Create a new Gin router
    r := gin.Default()

    // Define the root route
    r.GET("/", func(c *gin.Context) {
        // Return a JSON response with the message
        c.JSON(200, gin.H{
            "msg": "Hello World",
        })
    })

    // Run the server on port 8080
    r.Run(":8080")
}
