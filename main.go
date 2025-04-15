package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/chunzhennn/GOAD-Dashboard/docs"
	"github.com/chunzhennn/GOAD-Dashboard/internal/api/controllers"
	"github.com/chunzhennn/GOAD-Dashboard/internal/config"
	"github.com/chunzhennn/GOAD-Dashboard/internal/platform/pfsense"
	"github.com/chunzhennn/GOAD-Dashboard/internal/platform/proxmox"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

//go:embed ui/dist
var uiFS embed.FS

//go:embed docs/swagger.json
var swaggerJSON []byte

// @title GOAD Dashboard API
// @version 1.0
// @description GOAD Dashboard API
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	pveClient := proxmox.NewPVEClientFromConfig(config)
	pveController := controllers.NewPVEController(pveClient)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// PVE API endpoints
	router.Route("/api/pve", func(r chi.Router) {
		// GET group
		r.Group(func(r chi.Router) {
			r.Get("/vms", pveController.GetVMs)
			r.Get("/reset", pveController.GetLastReset)
		})

		// POST group
		r.Group(func(r chi.Router) {
			r.Post("/vms/start", pveController.StartAllVMs)
			r.Post("/vms/stop", pveController.StopAllVMs)
			r.Post("/vms/reset", pveController.ResetAllVMs)
			r.Post("/reset", pveController.ResetLab)
		})
	})

	pfsenseClient := pfsense.NewPfsenseClient(config)
	pfsenseController := controllers.NewPfsenseController(pfsenseClient)

	// PFSENSE API endpoints
	router.Route("/api/pfsense", func(r chi.Router) {
		r.Get("/openvpn/connections", pfsenseController.GetOpenVPNConnections)
	})

	swaggerEnabled := os.Getenv("ENABLE_SWAGGER")
	if swaggerEnabled == "1" {
		router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
			w.Write(swaggerJSON)
		})
		router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", config.GetPort()))))
	}

	ui, _ := fs.Sub(uiFS, "ui/dist")
	router.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(ui)).ServeHTTP(w, r)
	}))

	log.Printf("Starting GOAD Dashboard API server...")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.GetPort()), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
