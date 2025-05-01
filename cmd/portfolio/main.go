package main

import (
	"fmt"
	"log"
	"os"

	"github.com/billvamva/gomp/database"
	"github.com/billvamva/gomp/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require pgbouncer=true",
		"aws-0-eu-west-2.pooler.supabase.com",
		5432,
		"postgres.oucsyczvcrrwtdieynxt",
		os.Getenv("DB_PASSWORD"),
		"postgres")

	pool := database.ConnectToDB(connectionStr)
	projects, err := database.GetProjects()
	if err != nil {
		log.Fatalf("error getting projects")
	}

	fmt.Printf("%v", projects)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.LoadHTMLGlob("internal/templates/*")
	// Serve static files
	r.Static("/static", "web/static")
	r.GET("/", handlers.HandleMain)
	r.GET("/post/:id", handlers.HandlePost)

	r.Run(":8080")
	pool.Close()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
