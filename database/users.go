package database

import (
	"sort"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/util"
)

type photoTrending struct {
	name string
	like int
}
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
	return nil, util.ErrUnknownUser(name)
}
func (users *Users) Trending() []photoTrending {
	sort.Slice(users.user, func(i, j int) bool {
		first := users.user[i].photo
		var firstLen, secondLen = 0, 0
		if first != nil {
			firstLen = len(users.user[i].photo.likes)
		}
		second := users.user[j].photo
		if second != nil {
			secondLen = len(users.user[j].photo.likes)
		}
		return firstLen > secondLen
	})
	firstLike, secondLike := 0, 0
	firstPhotoPosition := users.user[0].photo
	secondPhotoPosition := users.user[1].photo
	if firstPhotoPosition != nil {
		firstLike = len(firstPhotoPosition.likes)
	}
	if secondPhotoPosition != nil {
		secondLike = len(secondPhotoPosition.likes)
	}
	firstTrending := photoTrending{name: users.user[0].name, like: firstLike}
	secondTrending := photoTrending{name: users.user[1].name, like: secondLike}
	return []photoTrending{firstTrending, secondTrending}

}
