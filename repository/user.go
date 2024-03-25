package repository

import (
	"golang-study/types"
	"golang-study/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{}, // Slice 초기화
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	isExisted := false
	for _, user := range u.userMap {
		if user.Name == newUser.Name {
			isExisted = true
			break
		}
	}
	if !isExisted {
		u.userMap = append(u.userMap, newUser)
		return nil
	} else {
		return errors.Errorf(errors.AlreadyExisted, newUser.Name)
	}
	// u.userMap = append(u.userMap, newUser)
	// return nil
}

func (u *UserRepository) Update(name string, newAge int64) error {
	isExisted := false
	for _, user := range u.userMap {
		if user.Name == name {
			user.Age = newAge
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(userName string) error {
	isExisted := false
	for index, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...)
			isExisted = true
			continue
		}
	}

	if !isExisted {
		// 에러 코드 작성
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
