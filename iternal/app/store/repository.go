package store

import "http-rest-api/iternal/app/model"

//UserRepository...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
