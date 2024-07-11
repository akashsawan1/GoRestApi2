package main

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
    c.String(200, "testing(Route 1)")
}

func jsonTest(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World! (Route2)",
    })
}

func Route3(c *gin.Context) {
    c.String(200, "This is the 3rd route")
}

func main() {
    router := gin.Default()

    router.GET("/test", Test)
    router.GET("/hello", jsonTest)
    router.GET("/route3" , Route3)
    router.Run(":8080")
}
