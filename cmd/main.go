package main

import (
	"log"
	"net/http"

	"github.com/sandokandias/go-database-app/pkg/godb/db"
	"github.com/sandokandias/go-database-app/pkg/godb/order"
	"github.com/sandokandias/go-database-app/pkg/godb/postgres"
)

func main() {
	log.Println("Starting godb app...")
	dbpool := postgres.Connect()
	defer dbpool.Close()

	txManager := db.NewTxManager(dbpool)
	customerStorage := postgres.NewCustomerStorage(dbpool)
	orderStorage := postgres.NewOrderStorage(dbpool)
	orderService := order.NewService(txManager, orderStorage, customerStorage)
	orderHandler := order.NewHandler(orderService)

	http.HandleFunc("/orders", orderHandler.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
