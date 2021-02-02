package persistence

import (
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/repository"
	"gopkg.in/gorp.v1"
)

type UserRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewUserRepository(dbMap *gorp.DbMap) repository.UserRepository {
	return &UserRepositoryImpl{Table: "user", DbMap: dbMap}
}

func (r *UserRepositoryImpl) FindByUserName(userName string) (*entity.User, error) {
	var user entity.User
	err := r.DbMap.SelectOne(&user, "SELECT * FROM "+r.Table+" WHERE user_name=? AND available=1", userName)

	if err == nil {
		return &user, nil
	}
	return nil, err
}

func (r *UserRepositoryImpl) FindByUserNameAndPassword(userName string, password string) (*entity.User, error) {
	var user entity.User
	err := r.DbMap.SelectOne(&user, "SELECT * FROM "+r.Table+" WHERE user_name=? AND password=? AND available=1", userName, password)

	if err == nil {
		return &user, nil
	}
	return nil, err
}

func (r *UserRepositoryImpl) FindAll() ([]entity.User, error) {
	var users []entity.User
	_, err := r.DbMap.Select(&users, "SELECT * FROM "+r.Table+" WHERE available=1")

	if err == nil {
		return users, nil
	}

	return nil, err
}
