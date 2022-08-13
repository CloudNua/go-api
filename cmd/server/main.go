package main

import (
	"fmt"

	"github.com/CloudNua/go-api/internal/db"
)

func Run() error {
	fmt.Println("Starting up....")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println()
	}
}
