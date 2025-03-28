package handlers

import (
	"github.com/gin-gonic/gin"
	"go-people-api/internal/model"
	"go-people-api/internal/repository"
	"net/http"
)

// Получить всех людей
func GetPeople(c *gin.Context) {
	people, err := repository.GetPeople()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, people)
}

// Добавить нового человека
func CreatePerson(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	_, err := repository.DB.NamedExec(`INSERT INTO people (first_name, last_name, gender, nationality, age) 
		VALUES (:first_name, :last_name, :gender, :nationality, :age)`, &person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Человек добавлен"})
}
