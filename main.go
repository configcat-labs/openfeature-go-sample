package main

import (
	"context"
	"net/http"

	configcat "github.com/configcat/go-sdk/v8"
	"github.com/gin-gonic/gin"
	provider "github.com/open-feature/go-sdk-contrib/providers/configcat/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

type course struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Instructor string  `json:"instructor"`
	Price      float64 `json:"price"`
}

var courses = []course{
	{ID: "1", Title: "Intro to HTML", Instructor: "Guy Person", Price: 49.99},
	{ID: "2", Title: "Intro to JavaScript", Instructor: "Mister Coder", Price: 59.99},
	{ID: "3", Title: "Building APIs with Go", Instructor: "Sir Programmer", Price: 60.99},
}

func main() {

	ccProvider := provider.NewProvider(configcat.NewClient("YOUR-SDK-KEY"))
	openfeature.SetProvider(ccProvider)
	client := openfeature.NewClient("app")

	router := gin.Default()
	router.GET("/courses", func(c *gin.Context) {
		getCoursesEnabled, _ := client.BooleanValue(
			context.Background(), "get_courses_enabled", false, openfeature.EvaluationContext{},
		)
		if getCoursesEnabled {
			c.IndentedJSON(http.StatusOK, courses)
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Not Found",
			})
		}
	})

	router.Run("localhost:8000")
}
