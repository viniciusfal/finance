package handlers

import (
	"net/http"

	"manager/internal/usecases"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	usecase *usecases.DashboardUsecase
}

func NewDashboardHandler(usecase *usecases.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{usecase: usecase}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	summary, err := h.usecase.GetSummary(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}

