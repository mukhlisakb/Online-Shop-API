package main

import (
	"Online-Shop-API/database"
	"Online-Shop-API/handler"
	"Online-Shop-API/middleware"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DB_URI"))
	if err != nil {
		fmt.Printf("Failed to connection to the Database...%v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Failed to verification to the Database...%v\n", err)
		os.Exit(1)
	}

	if _, err = database.DatabaseMigrate(db); err != nil {
		fmt.Printf("Failed to migrate the Database...%v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/api/v1/products", handler.ListProduct(db))
	r.GET("/api/v1/products/:id", handler.GetProducts(db))
	r.POST("/api/v1/checkout", handler.CheckOutOrder(db))

	r.POST("/api/v1/order/:id/confirm", handler.Confirm(db))
	r.GET("/api/v1/products/id", handler.GetOrder(db))

	r.POST("/admin/products", middleware.AdminOnly(), handler.CreateProducts(db))
	r.PUT("/admin/products/:id", middleware.AdminOnly(), handler.UpdateProducts(db))
	r.DELETE("/admin/products/:id", middleware.AdminOnly(), handler.DeleteProducts(db))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err = server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to running the Server...%v\n", err)
		os.Exit(1)
	}
}
