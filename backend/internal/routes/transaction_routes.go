package routes

import (
	"manager/internal/handlers"
	"manager/internal/repositories"
	"manager/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupTransactionRoutes(router *gin.RouterGroup, db *pgxpool.Pool) {
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionUsecase := usecases.NewTransactionUsecase(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionUsecase)

	transactions := router.Group("/transactions")
	{
		transactions.GET("", transactionHandler.GetAll)
		transactions.GET("/:id", transactionHandler.GetByID)
		transactions.POST("", transactionHandler.Create)
		transactions.PUT("/:id", transactionHandler.Update)
		transactions.DELETE("/:id", transactionHandler.Delete)
		transactions.POST("/:id/installments/:installment/pay", transactionHandler.PayInstallment)
	}
}

