package activityreporter

import (
	"github.com/Lucifer07/activity-reporter/utility"
)

type Users struct {
	user []*User
}

func NewUsers() *Users {
	return &Users{}
}
func (users *Users) NewUser(name string) *User {
	user := &User{
		name: name,
	}
	users.user = append(users.user, user)
	return user
}

func (users *Users) CheckUser(name string) (*User, error) {
	for _, user := range users.user {
		if user.name == name {
			return user, nil
		}
	}
	return nil, utility.ErrUnknownUser(name)
}
