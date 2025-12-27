package main

import (
	"log"
	"os"

	"manager/internal/config"
	"manager/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Inicializar conexão com banco de dados
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Configurar Gin
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Criar router
	router := gin.Default()

	// Middleware de logging para debug
	router.Use(func(c *gin.Context) {
		log.Printf("[%s] %s %s", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr)
		c.Next()
	})

	// Configurar CORS - Aceita qualquer origem (*)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Backend is running"})
	})

	// Rota raiz para teste
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Financy API",
			"version": "1.0.0",
			"endpoints": []string{
				"/api/transactions",
				"/api/categories",
				"/api/dashboard/summary",
			},
		})
	})

	// Configurar rotas
	log.Println("Setting up routes...")
	routes.SetupRoutes(router, db)
	log.Println("Routes configured successfully")

	// Log das rotas registradas (sempre logar para debug)
	log.Println("Registered routes:")
	for _, route := range router.Routes() {
		log.Printf("  %s %s", route.Method, route.Path)
	}

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on 0.0.0.0:%s", port)
	log.Printf("API available at http://0.0.0.0:%s/api", port)
	log.Printf("Health check: http://0.0.0.0:%s/health", port)

	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
