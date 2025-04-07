package person

type Repository interface {
	GetByID(id int) (*Person, error)
	Create(user *Person) error
}
