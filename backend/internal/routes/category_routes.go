package routes

import (
	"manager/internal/handlers"
	"manager/internal/repositories"
	"manager/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupCategoryRoutes(router *gin.RouterGroup, db *pgxpool.Pool) {
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryUsecase := usecases.NewCategoryUsecase(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecase)

	categories := router.Group("/categories")
	{
		categories.GET("", categoryHandler.GetAll)
		categories.GET("/:id", categoryHandler.GetByID)
		categories.POST("", categoryHandler.Create)
		categories.PUT("/:id", categoryHandler.Update)
		categories.DELETE("/:id", categoryHandler.Delete)
	}
}

