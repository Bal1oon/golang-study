package repository

import (
	"sync"
)

// network, repository 중간 역할

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	// db mongo.database 같은 데이터 값을 가지고 있어야 함

	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})

	return repositoryInstance
}
