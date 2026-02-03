package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"codeid.hr-api/api/routes"
	configs "codeid.hr-api/internal/config"
	"codeid.hr-api/internal/models"
	"codeid.hr-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//1. set environment (bisa cmd atau system environment)
	os.Setenv("APP_ENV", "development")
	//2.Load configuration
	config := configs.Load()
	//1. current code lebih minimalis
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatal("failed to initialize database:%w", err)
	}
	defer database.CloseDB(db)
	// Run auto migration
	if err := database.AutoMigrate(db, &models.Region{}, &models.Country{}); err != nil {
		log.Printf("Warning: Auto migration failed: %v", err)
	}
	// Set Gin mode based on environment
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// Setup routes
	router := gin.Default()
	routes.SetupRoutes(router, db.DB)
	// Start server
	log.Printf("Server starting on %s in %s mode", config.Server.Address,
		config.Environment)
	go func() {
		if err := router.Run(config.Server.Address); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	// 1. Graceful shutdown : nunggu operasi selesai baru shutdown server
	// 2. Tanpa Graceful Shutdowon : close connection,ada kemungkinan operasi seperti query masih jalan
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

}

/* import (
	"log"
	"net/http"

	"codeid.hr-api/internal/handlers"
	"codeid.hr-api/internal/repositories"
	"codeid.hr-api/internal/services"
	"codeid.hr-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize database connection
	db, err := database.SetupDB()
	if err != nil {
		log.Fatal("Failed to Connect %w", err)
	}

	//1.1. init automigrate --> baru
	// InitAutoMigrate parameter *gorm.DB, disupply dari variabel db
	database.InitAutoMigrate(db)

	//5. Initialize repositories
	regionRepo := repositories.NewRegionRepository(db)
	//6. Initialize services
	regionService := services.NewRegionService(regionRepo)
	//7. Initialize handlers
	regionHandler := handlers.NewRegionHandler(regionService)

	// setup router
	router := gin.Default()

	//8. grouping subroute with prefix /api
	api := router.Group("/api")
	{
		// region routes endpoints
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
		}
	}
	log.Println("server starting on port :8080")
	//4. run webserver at port 8080
	router.Run(":8080")
}

func helloworldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
		"status":  "running",
	})
}
*/
