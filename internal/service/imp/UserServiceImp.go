package imp

import (
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/storage"
	"context"
	"fmt"
)

type UserServiceImp struct {
	rep storage.UserRepository
}

func NewUserService(rep storage.UserRepository) *UserServiceImp {
	return &UserServiceImp{rep: rep}
}

func (s UserServiceImp) Find(ctx context.Context, username string) (user.UserDB, error) {

	u, err := s.rep.Get(ctx, username)

	if err != nil {
		return user.UserDB{}, fmt.Errorf("failed to find user: %w", err)
	}
	if u == nil {
		return user.UserDB{}, fmt.Errorf("user not found")
	}

	return *u, nil
}
