package main

import (
	"fmt"

	"github.com/CloudNua/go-api/internal/comment"
	"github.com/CloudNua/go-api/internal/db"
	transportHttp "github.com/CloudNua/go-api/internal/transport/http"
)

// Run - is going to be responsible for
// the instantiation and startup of the
// go application
func Run() error {
	fmt.Println("starting up application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
