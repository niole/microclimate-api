package persister

import (
	api "api/protobuf"
)

type Scannable = func(dest ...interface{}) error

func ScanOneUser(result Scannable, user *api.User) error {
	var id string
	var email string
	var name string

	err := result(&id, &email, &name)

	*user = api.User{
		Id:    id,
		Email: email,
		Name:  name,
	}

	return err
}
