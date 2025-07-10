package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Project struct {
	Name        string   `form:"name" json:"name"`
	Description string   `form:"description" json:"description`
	Language    string   `form:"language" json:"language"`
	Modules     []string `form:"modules" json:"modules"`
}

func Server() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "web/static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Lyra: Fiware Project Creator",
		})
	})

	router.GET("/project", func(c *gin.Context) {
		c.HTML(200, "project.html", gin.H{
			"title": "Lyra: Create Project",
		})
	})

	router.POST("/project", func(c *gin.Context) {
		var project Project

		if err := c.ShouldBind(&project); err != nil {
			c.JSON(400, gin.H{"error": "Project name is required"})
			fmt.Printf("Failed to bind project data: %v", err)
			return
		}

		c.JSON(200, gin.H{
			"message": "Project created successfully",
		})
	})

	router.Run()
}
