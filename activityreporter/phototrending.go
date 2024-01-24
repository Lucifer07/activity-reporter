package activityreporter

import (
	"sort"
)

type photoTrending struct {
	Name string
	Like int
}

func (users *Users) Trending() []photoTrending {
	var trending []photoTrending
	count := 1
	users.sortTrending()
	for _, user := range users.user {
		if count <= 3 {
			photo := user.photo
			if photo != nil {
				trending = append(trending, photoTrending{user.name, len(photo.likes)})
			}
		}
		if count > 3 {
			break
		}
		count++
	}
	return trending
}
func (users *Users) sortTrending() *Users {
	sort.Slice(users.user, func(i, j int) bool {
		first := users.user[i].photo
		var firstLen, secondLen = -1, -1
		if first != nil {
			firstLen = len(users.user[i].photo.likes)
		}
		second := users.user[j].photo
		if second != nil {
			secondLen = len(users.user[j].photo.likes)
		}
		return firstLen > secondLen
	})
	return users
}
