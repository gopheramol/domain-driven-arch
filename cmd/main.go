package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/gopheramol/domain-driven-arch/config"
	adHandler "github.com/gopheramol/domain-driven-arch/internal/ad/handler"
	adRepository "github.com/gopheramol/domain-driven-arch/internal/ad/repository"
	adService "github.com/gopheramol/domain-driven-arch/internal/ad/service"

	"github.com/gopheramol/domain-driven-arch/internal/user/handler"
	"github.com/gopheramol/domain-driven-arch/internal/user/repository"
	"github.com/gopheramol/domain-driven-arch/internal/user/service"
	"github.com/gopheramol/domain-driven-arch/pkg/db/postgres"
	"github.com/gopheramol/domain-driven-arch/web/middleware"
)

func main() {
	// Load config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	// Connect to PostgreSQL database
	db, err := postgres.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}

	defer db.Close()

	// Initialize user repository and service
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize ad repository and service
	adRepo := adRepository.NewAdRepository(db)
	adService := adService.NewAdService(adRepo)
	adHandler := adHandler.NewAdHandler(adService)

	// Initialize Gin router
	router := gin.Default()
	// Apply the logger middleware to log incoming requests and responses
	router.Use(middleware.LoggerMiddleware())

	// Define user routes
	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/", userHandler.CreateUser)
			userGroup.GET("/", userHandler.GetAllUsers)
			userGroup.GET("/:id", userHandler.GetUserByID)
			userGroup.PUT("/:id", userHandler.UpdateUser)
			userGroup.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// Define ad routes
	{
		adGroup := v1.Group("/ads")
		{
			adGroup.POST("/", adHandler.CreateAd)
			adGroup.GET("/", adHandler.GetAllAds)
			adGroup.GET("/:id", adHandler.GetAdByID)
			adGroup.PUT("/:id", adHandler.UpdateAd)
			adGroup.DELETE("/:id", adHandler.DeleteAd)
			adGroup.GET("/user/:user_id", adHandler.GetAdByUserID)
		}
	}
	// Start the server
	log.Printf("Server listening on %s", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, router))
}
