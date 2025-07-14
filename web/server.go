package web

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lyra/globals"
	"lyra/internal"
	"lyra/types"
	_ "lyra/types"
	"os"
)

func Server() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "web/static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Lyra: Fiware Project Creator",
		})
	})

	router.GET("/project", func(c *gin.Context) {
		c.HTML(200, "project.html", gin.H{
			"title":     "Lyra: Create Project",
			"languages": globals.AvailableLanguages,
			"licenses":  globals.AvailableLicenses,
			"modules":   globals.AvailableModules,
		})
	})

	router.GET("/project-creating", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.HTML(400, "error.html", gin.H{
				"title":   "Error",
				"message": "Project ID is required.",
			})
			return
		}
		projectFile := fmt.Sprintf("projects/%s.json", id)
		if _, err := os.Stat(projectFile); os.IsNotExist(err) {
			c.HTML(404, "error.html", gin.H{
				"title":   "Error",
				"message": "Project not found.",
			})
			return
		}

		file, err := os.ReadFile(projectFile)
		if err != nil {
			c.HTML(500, "error.html", gin.H{
				"title":   "Error",
				"message": "Failed to read project file.",
			})
			fmt.Printf("Failed to read project file: %v", err)
			return
		}

		var project types.Project
		if err := json.Unmarshal(file, &project); err != nil {
			c.HTML(500, "error.html", gin.H{
				"title":   "Error",
				"message": "Failed to parse project data.",
			})
			fmt.Printf("Failed to parse project data: %v", err)
			return
		}

		c.HTML(200, "project-creating.html", gin.H{
			"title":   "Lyra: Creating Project",
			"project": project,
		})
	})

	router.POST("/project", func(c *gin.Context) {
		var project types.Project

		if err := c.ShouldBind(&project); err != nil {
			c.JSON(400, gin.H{"error": "Project name is required"})
			fmt.Printf("Failed to bind project data: %v", err)
			return
		}

		buf := make([]byte, 16)
		_, err := rand.Read(buf)
		key := fmt.Sprintf("%x", buf)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate random ID"})
			fmt.Printf("Failed to generate random ID: %v", err)
			return
		}
		file, err := os.Create(fmt.Sprintf("projects/%s.json", key))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create project file"})
			fmt.Printf("Failed to create project file: %v", err)
			return
		}
		defer file.Close()
		projectJSON, err := json.Marshal(project)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to marshal project data"})
			fmt.Printf("Failed to marshal project data: %v", err)
			return
		}

		if _, err := file.Write(projectJSON); err != nil {
			c.JSON(500, gin.H{"error": "Failed to write project data"})
			fmt.Printf("Failed to write project data: %v", err)
			return
		}

		err = internal.ScaffoldProject(project)
		if err != nil {
			return
		}
		c.Redirect(302, "/project-creating?id="+key)
	})

	router.Run()
}
