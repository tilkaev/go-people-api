package person

import (
	"go-people-api/internal/domain/person"
	"strconv"
)

type UseCase struct {
	personRepo person.Repository
}

func NewUseCase(repo person.Repository) *UseCase {
	return &UseCase{personRepo: repo}
}

func (uc *UseCase) GetPerson(id int) (*person.Person, error) {
	return uc.personRepo.GetByID(id)
}

func (uc *UseCase) CreatePerson(first_name, last_name, gender, nationality, age string) (*person.Person, error) {

	num, err := strconv.Atoi(age)
	if err != nil {
		return nil, err
	}

	newPerson := &person.Person{
		FirstName:   first_name,
		LastName:    last_name,
		Gender:      gender,
		Nationality: nationality,
		Age:         num,
	}

	err = uc.personRepo.Create(newPerson)
	if err != nil {
		return nil, err
	}

	return newPerson, nil
}
