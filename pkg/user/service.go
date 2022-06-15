package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Route(g *gin.RouterGroup) {
	users := g.Group("/users")
	users.POST("/sign_up", s.SignUp)
	users.POST("/login", s.Login)
}

func (s *Service) SignUp(c *gin.Context) {
	var params UserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := s.repo.Create(params.Username, params.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
	return
}

func (s *Service) Login(c *gin.Context) {
	var params UserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// naive user login validation, in production i will never do this
	user, err := s.repo.FindByUserName(params.Username)
	if err != nil {
		if err == RecordNotFound {
			id, err := s.repo.Create(params.Username, params.Password)

			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, gin.H{"success": true, "user_id": id})
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return

	}

	// if user.Password != params.Password {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"success": true, "user_id": user.ID})
	return
}