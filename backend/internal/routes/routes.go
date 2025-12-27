package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	api := router.Group("/api")
	{
		SetupTransactionRoutes(api, db)
		SetupCategoryRoutes(api, db)
		SetupDashboardRoutes(api, db)
	}
}
