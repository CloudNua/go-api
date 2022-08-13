package main

import (
	"context"
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
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected to db")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println()
	}
}
