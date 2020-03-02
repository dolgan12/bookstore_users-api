package users

import (
	"fmt"

	"github.com/dolgan12/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get fetches a user from persistance
func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("useer %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

//Save saves user to persistant layer
func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
