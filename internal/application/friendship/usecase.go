package friendship

import (
	"go-people-api/internal/domain/friendship"
)

type UseCase struct {
	repo friendship.Repository
}

func NewUseCase(repo friendship.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) GetUser(id int) (*friendship.Friendship, error) {
	return uc.repo.GetByID(id)
}
