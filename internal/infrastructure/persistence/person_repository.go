package persistence

import (
	"go-people-api/internal/domain/person"
	"go-people-api/internal/repository"
)

type PersonRepository struct{}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{}
}

func (r *PersonRepository) Create(person *person.Person) error {
	query := `INSERT INTO orders (
                    first_name, last_name, gender, nationality, age
                    ) VALUES ($1, $2, $3, $4, $5)`
	return repository.DB.QueryRow(
		query,
		person.FirstName,
		person.LastName,
		person.Gender,
		person.Nationality,
		person.Age,
	).Scan(&person.ID)
}

func (r *PersonRepository) GetByID(id int) (*person.Person, error) {
	var o person.Person
	query := `SELECT id, product_id, quantity, status, created_at FROM orders WHERE id = $1`
	err := repository.DB.Get(&o, query, id)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *PersonRepository) GetAll() (*[]person.Person, error) {
	var people []person.Person
	query := `SELECT * FROM people`
	err := repository.DB.Select(&people, query)
	if err != nil {
		return nil, err
	}
	return &people, nil
}
