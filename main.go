package main

import (
	"double.com/Database"
	"double.com/Model"
	"fmt"
	"github.com/gin-gonic/gin"
)

var m = Database.Mario{}
var user = Model.User{}

type URI struct {
	ID string `uri:"id"`
}

func main() {

	u := Model.Mario{Length:0, Weight: 0, UpdateTime: "2019-08-10 15:04:05"}
	u.UpdateSizeIfNeeded()
	fmt.Println(u.Length)
	fmt.Println(u.Weight)
	fmt.Println(u.UpdateTime)

	r := gin.Default()

	r.GET("/marios", func(c *gin.Context) {
		marios, err := m.GetAll()
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
		} else {
			c.JSON(200, marios)
		}
	})
	r.GET("/mario/:id", func(c *gin.Context) {
		var uri URI
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		mario, err := m.Get(uri.ID)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
		} else if mario == (Database.Mario{}) {
			c.JSON(400, gin.H{"msg": "not found"})
		} else {
			c.JSON(200, mario)
		}
	})
	r.POST("/mario", func(c *gin.Context) {
		var json Database.Mario
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	    id, err := m.Insert(json)
	    if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"msg": id})
		}
	})
	r.Run(":5000")
}
