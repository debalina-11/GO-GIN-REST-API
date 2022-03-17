package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello GO",
	})
}
func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameter(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)            // /query?name=avi&age=27
	r.GET("/path/:name/:age", PathParameter) //  /path/avi/100
	r.Run()
}
