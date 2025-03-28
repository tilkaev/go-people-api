package http

import (
	"github.com/gin-gonic/gin"
	"go-people-api/internal/application/person"
	"net/http"
)

type PersonHandler struct {
	useCase *person.UseCase
}

func NewOrderHandler(useCase *person.UseCase) *PersonHandler {
	return &PersonHandler{useCase: useCase}
}

func (h *PersonHandler) CreatePeople(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	gender := c.Query("gender")
	nationality := c.Query("nationality")
	age := c.Query("age")

	newPerson, err := h.useCase.CreatePerson(first_name, last_name, gender, nationality, age)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newPerson)
}
