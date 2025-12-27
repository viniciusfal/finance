package routes

import (
	"manager/internal/handlers"
	"manager/internal/repositories"
	"manager/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDashboardRoutes(router *gin.RouterGroup, db *pgxpool.Pool) {
	transactionRepo := repositories.NewTransactionRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	dashboardUsecase := usecases.NewDashboardUsecase(transactionRepo, categoryRepo)
	dashboardHandler := handlers.NewDashboardHandler(dashboardUsecase)

	dashboard := router.Group("/dashboard")
	{
		dashboard.GET("/summary", dashboardHandler.GetSummary)
	}
}

