package repository

import "golang-study/types"

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{}, // Slice 초기화
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Update(beforeUser *types.User, updatedUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(user *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
