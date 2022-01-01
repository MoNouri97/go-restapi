package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" binding:"required"`
}

var db = make([]Person, 0)

func UsersController(r *gin.Engine) {

	users := r.Group("/users")
	{
		users.GET("/:id", func(c *gin.Context) {
			id, e := strconv.Atoi(c.Param("id"))
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
			}

			for _, element := range db {
				if element.Id == id {
					c.JSON(http.StatusOK, element)
					return
				}
			}
			c.JSON(http.StatusNotFound, nil)
		})
		users.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, db)
		})
		users.POST("/", func(c *gin.Context) {
			var person Person
			if e := c.ShouldBindJSON(&person); e != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
			} else {
				person.Id = len(db)
				db = append(db, person)
				c.JSON(http.StatusCreated, gin.H{
					"message": "ok",
					"data":    person,
				})
			}
		})
		users.DELETE("/:id", func(c *gin.Context) {
			id, e := strconv.Atoi(c.Param("id"))
			if e != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
			}
			dbCopy := make([]Person, 0)
			for _, element := range db {
				if element.Id != id {
					dbCopy = append(dbCopy, element)
				}
			}
			db = dbCopy

		})

	}
}
