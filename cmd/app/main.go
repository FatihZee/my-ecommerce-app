package main

import (
    "log"
    "my-ecommerce-app/config"
    "my-ecommerce-app/internal/handlers"
    "my-ecommerce-app/internal/middleware"
    "my-ecommerce-app/pkg/database"

    "github.com/gin-gonic/gin"
)

func main() {
    config, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    err = database.Connect(config)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }

    database.Migrate()

    r := gin.Default()
    r.Use(middleware.AuthMiddleware(config))

    // Setup routes
    productHandler := handlers.NewProductHandler()
    r.POST("/products", productHandler.CreateProduct)
    // Tambahkan route lainnya

    r.Run(":8080")
}
