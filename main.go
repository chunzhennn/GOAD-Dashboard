package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chunzhennn/GOAD-Dashboard/api/controllers"
	_ "github.com/chunzhennn/GOAD-Dashboard/docs"
	"github.com/chunzhennn/GOAD-Dashboard/pve"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	pveClient := pve.NewPVEClientFromConfig(config)
	pveController := controllers.NewPVEController(pveClient)

	router := gin.Default()

	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GOAD Dashboard API is running",
		})
	})

	// PVE API endpoints
	pveGroup := router.Group("/api/pve")
	{
		pveGroup.GET("/vms", pveController.GetVMs)
		pveGroup.POST("/vms/start", pveController.StartAllVMs)
		pveGroup.POST("/vms/stop", pveController.StopAllVMs)
		pveGroup.POST("/vms/reset", pveController.ResetAllVMs)
		pveGroup.GET("/reset", pveController.GetLastReset)
		pveGroup.POST("/reset", pveController.ResetLab)
	}

	swaggerEnabled := os.Getenv("ENABLE_SWAGGER")
	if swaggerEnabled == "1" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Printf("Starting GOAD Dashboard API server...")
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
