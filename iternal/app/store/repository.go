package store

import "http-rest-api/iternal/app/model"

//UserRepository...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
