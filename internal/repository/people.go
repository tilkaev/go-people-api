package repository

import (
	"github.com/samber/lo"
	"go-people-api/internal/model"
	"log"
)

func GetPeople() ([]model.Person, error) {
	var peoples []dbPerson
	err := DB.Select(&peoples, "SELECT * FROM people")
	if err != nil {
		return nil, err
	}

	return lo.Map(peoples, func(people dbPerson, _ int) model.Person {
		return model.Person{
			ID:          people.ID,
			FirstName:   people.Name,
			LastName:    people.LastName,
			Gender:      people.Gender,
			Nationality: people.Nationality,
			Age:         people.Age,
		}
	}), nil
}

func CreatePerson(name, lastName, gender, nationality string) error {
	query := `INSERT INTO people (first_name, last_name, gender, nationality)
						VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, name, lastName, gender, nationality)
	if err != nil {
		log.Println("❌ Ошибка при вставке данных:", err)
	}
	return err
}

type dbPerson struct {
	ID          uint   `db:"id"`
	Name        string `db:"first_name"`
	LastName    string `db:"last_name"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
	Age         int    `db:"age"`
}
