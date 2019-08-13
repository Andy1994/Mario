package main

import (
	"double.com/Database"
	"github.com/gin-gonic/gin"
)

var m = Database.Mario{}

type URI struct {
	ID string `uri:"id"`
}

func main() {

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
	r.Run()

	//router.HandleFunc("/marios", getMarios).Methods("GET")
	//router.HandleFunc("/mario/{id}", getMario).Methods("GET")
	//router.HandleFunc("/mario", createMario).Methods("POST")
	//router.HandleFunc("/mario/{id}", updateMario).Methods("POST")
	//router.HandleFunc("/mario/{id}", deleteMario).Methods("DELETE")
}
