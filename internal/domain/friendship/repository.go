package friendship

type Repository interface {
	GetByID(id int) (*Friendship, error)
	Create(user *Friendship) error
}
