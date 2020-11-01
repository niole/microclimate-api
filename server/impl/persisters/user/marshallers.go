package user

import (
	api "api/protobuf"
)

type Scannable = func(dest ...interface{}) error

func ScanOneUser(result Scannable, user *api.User) error {
	var id string
	var email string

	err := result(&id, &email)

	*user = api.User{
		Id:    id,
		Email: email,
	}

	return err
}
