package main

import (
	"log"
	"net/http"

	"github.com/sandokandias/go-database-app/pkg/godb/order"
	"github.com/sandokandias/go-database-app/pkg/godb/postgres"
)

func main() {
	log.Println("Starting godb app...")
	db := postgres.Connect()
	defer db.Close()

	txManager := postgres.NewTxManager(db)
	orderStorage := postgres.NewOrderStorage(txManager)
	orderService := order.NewService(orderStorage)
	orderHandler := order.NewHandler(orderService)

	http.HandleFunc("/orders", orderHandler.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
