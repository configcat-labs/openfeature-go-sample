package main

import (
	"net/http"

	configcat "github.com/configcat/go-sdk/v8"
	"github.com/gin-gonic/gin"
)

type course struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var courses = []course{
	{ID: "1", Title: "Intro to HTML", Author: "Guy Person", Price: 49.99},
	{ID: "2", Title: "Intro to JavaScript", Author: "Mister Coder", Price: 59.99},
	{ID: "3", Title: "Building APIs with Go", Author: "Sir Programmer", Price: 60.99},
}

func main() {
	// Initialize the ConfigCat client. Don't forget to replace "YOUR-SDK-KEY" with your ConfigCat SDK key
	client := configcat.NewClient("YOUR-SDK-KEY")

	router := gin.Default()
	router.GET("/courses", func(c *gin.Context) {
		// fetch the value of the feature flag
		getCoursesEnabled := client.GetBoolValue("get_courses_enabled", false, nil)

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
