package main

import (
	"go-people-api/internal/application/person"
	"go-people-api/internal/infrastructure/persistence"
	"go-people-api/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {

	personRepo := persistence.NewPersonRepository()
	personUS := person.NewUseCase(personRepo)

	orderHandler := http.NewOrderHandler(personUS)

	r := gin.Default()
	r.POST("/addPeople", orderHandler.CreatePeople)

	r.Run(":8080")

	/*
		repository.InitDB()

		// Создаём новый роутер Gin
		r := gin.Default()

		r.GET("/people", handlers.GetPeople)
		r.POST("/people", handlers.CreatePerson)

		// Запускаем сервер на порту 8080
		r.Run(":8080")*/

	/*
		// Маршрут для проверки работы сервера
		r.GET("/getUsers", func(c *gin.Context) {
			fmt.Println(Storage.GetPeople())

			c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
		})

		// Маршрут для проверки работы сервера
		r.GET("/createUser", func(c *gin.Context) {
			fmt.Println(Storage.CreatePerson(
				"123",
				"123",
				"123",
				"123",
			))

			c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
		})

		r.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
		})

		r.GET("/search", func(c *gin.Context) {
			query := c.Query("q") // search?q=Go
			c.JSON(http.StatusOK, gin.H{"search": query})
		})
	*/

}
