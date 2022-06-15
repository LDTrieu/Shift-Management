package shift 

import (
	"fmt"
	"net/http"

	"github.com/ldtrieu/staffany-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Route(g *gin.RouterGroup) {
	shifts := g.Group("/shifts")
	shifts.PUT("/:id", s.UpdateShift)
	shifts.DELETE("/:id", s.DeleteShift)
}

type ShiftParams struct {
	DateID uint `json:"date_id"`
	UserID uint `json:"user_id"`

	QuarterStart uint `json:"quarter_start"`
	NumQuarter   uint `json:"num_quarter"`

	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *Service) UpdateShift(c *gin.Context) {
	var params ShiftParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shiftIDStr := c.Param("id")
	shiftID := utils.StringToUint(shiftIDStr)

	shift, err := s.repo.FindByID(shiftID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shift.QuarterStart = params.QuarterStart
	shift.Title = params.Title
	shift.Description = params.Description
	shift.NumQuarter = params.NumQuarter
	fmt.Printf("shift = %+v\n", shift)

	err = s.repo.Update(shift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}

func (s *Service) DeleteShift(c *gin.Context) {
	shiftIDStr := c.Param("id")
	shiftID := utils.StringToUint(shiftIDStr)
	err := s.repo.DeleteByID(shiftID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}