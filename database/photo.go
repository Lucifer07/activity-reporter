package database

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/util"
)

type photo struct {
	likes []*user
}

func (dataUser *user) UploadPhoto() (*photo, error) {
	if dataUser.photo != nil {
		return nil, util.ErrUploadePhoto
	}
	dataUser.photo = &photo{}
	dataUser.activity = append(dataUser.activity, util.UploadedPhotoMessage("You"))
	dataUser.notifyObserverUploadPhoto()
	return dataUser.photo, nil
}
func (dataUser *user) Like(user *user) error {
	if dataUser.photo != nil {
		wasLikes := user.checkLikes(dataUser)
		if wasLikes != nil {
			return wasLikes
		}
		user.photo.likes = append(user.photo.likes, dataUser)
		dataUser.activity = append(dataUser.activity, util.LikeSelfPhoto)
		dataUser.notifyObserverLikePhoto(user)
		return nil
	}
	isNotFollower := user.checkFollower(dataUser)
	if isNotFollower != nil {
		return isNotFollower
	}
	if user != dataUser && user.photo == nil {
		return util.ErrDoesntHavePhoto(user.name)
	}
	return util.ErrDontHavePhoto
}
func (dataUser *user) checkFollower(user *user) error {
	for _, usr := range dataUser.followers {
		if usr == user {
			return nil
		}
	}
	return util.ErrUnableToLike(dataUser.name)
}
func (dataUser *user) checkLikes(user *user) error {
	for _, usr := range dataUser.photo.likes {
		if usr == user {
			return util.ErrAlreadyLike
		}
	}
	return nil
}
