package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id          int
	Name        string            `json:"name,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
	Description string            `json:"description,omitempty"`
	Url         string            `json:"url,omitempty"`
}

var (
	pool *pgxpool.Pool
	err  error
)

func ConnectToDB(connectionStr string) *pgxpool.Pool {
	pool, err = pgxpool.New(context.Background(), connectionStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	return pool
}

func AddProject(project Project) error {
	if pool == nil {
		log.Fatalf("connection with db not established")
	}
	tagBytes, err := json.Marshal(project.Tags)
	if err != nil {
		log.Fatalf("could not write tags to json format")
	}
	_, err = pool.Exec(context.Background(), "insert into projects(id, name, description, url, tags) values($1, $2, $3, $4, $5)",
		project.Id, project.Name, project.Description, project.Url, tagBytes)
	return err
}

func GetProjects() ([]Project, error) {
	if pool == nil {
		log.Fatalf("connection with db not established")
	}
	rows, err := pool.Query(context.Background(), "select * from projects")
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 0)

	for rows.Next() {
		var id int
		var name, description, url string
		var tagsBytes []byte
		rows.Scan(&id, &name, &description, &url, &tagsBytes)

		tags := map[string]string{}

		err := json.Unmarshal(tagsBytes, &tags)
		if err != nil {
			return nil, fmt.Errorf("failed decoding tag json")
		}

		projects = append(projects, Project{id, name, tags, description, url})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return projects, nil
}
