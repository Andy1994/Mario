package main

import (
	"double.com/Database"
	"double.com/Model"
	"github.com/gin-gonic/gin"
)

var m = Database.Mario{}
var users map[string]*Model.User

type URI struct {
	ID string `uri:"id"`
}

type Header struct {
	UUID string `header:"uuid"`
}

func main() {

	users = make(map[string]*Model.User)

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
	r.GET("/new", func(c *gin.Context) {
		h := Header{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}
		if h.UUID != "" {
			u := users[h.UUID]
			if u == nil {
				u := Model.NewUser(h.UUID)
				users[h.UUID] = u
				c.JSON(201, u)
			} else {
				u.UpdateIfNeeded()
				c.JSON(200, u)
			}
		} else {
			c.JSON(400, "UUID not found!")
		}
	})
	r.Run(":5000")
}
