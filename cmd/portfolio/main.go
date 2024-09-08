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
	err := godotenv.Load("/Users/vasilieiosvamvakas/Documents/projects/gomp/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require pgbouncer=true",
		"aws-0-eu-west-2.pooler.supabase.com",
		5432,
		"postgres.oucsyczvcrrwtdieynxt",
		os.Getenv("DB_PASSWORD"), // Replace with your actual password
		"postgres")

	pool := database.ConnectToDB(connectionStr)
	projects, err := database.GetProjects()
	if err != nil {
		log.Fatalf("error getting projects")
	}

	fmt.Printf("%v", projects)

	r := gin.Default()
	r.LoadHTMLGlob("internal/templates/*")
	// Serve static files
	r.Static("/static", "web/static")
	r.GET("/", handlers.HandleMain)

	r.Run(":8080")
	pool.Close()
}
