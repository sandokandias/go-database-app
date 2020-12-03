package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sandokandias/go-database-app/pkg/godb/order"
	"github.com/sandokandias/go-database-app/pkg/godb/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	log.Println("Starting godb app...")
	db := connectDB()
	defer db.Close()

	orderStorage := postgres.NewOrderStorage(db)
	orderService := order.NewService(orderStorage)
	orderHandler := order.NewHandler(orderService)

	http.HandleFunc("/orders", orderHandler.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() *pgxpool.Pool {
	fmt.Println("connecting to ", os.Getenv("DATABASE_URL"))
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db
}
