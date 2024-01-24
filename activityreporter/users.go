package activityreporter

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"
)

type Users struct {
	user []*user
}

func NewUsers() *Users {
	return &Users{}
}
func (users *Users) NewUser(name string) *user {
	user := &user{
		name: name,
	}
	users.user = append(users.user, user)
	return user
}

func (users *Users) CheckUser(name string) (*user, error) {
	for _, user := range users.user {
		if user.name == name {
			return user, nil
		}
	}
	return nil, utility.ErrUnknownUser(name)
}
