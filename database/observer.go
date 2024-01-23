package database

import "git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/util"

type observer interface {
	updateLike(*user)
	updateUploadPhoto(*user)
}

func (dataUser *user) updateUploadPhoto(user *user) {
	dataUser.activity = append(dataUser.activity, util.UploadedPhotoMessage(user.name))
}
func (dataUser *user) updateLikePhoto(user *user, photoOwner *user) {
	dataUser.activity = append(dataUser.activity, util.LikePhotoMessage(user.name, photoOwner.name))
}

type publisher interface {
	updateFollower(observer)
	notifyObserverLike(photoOwner *user)
	notifyObserverUploadPhoto()
}

func (dataUser *user) notifyObserverUploadPhoto() {
	for _, follower := range dataUser.followers {
		follower.updateUploadPhoto(dataUser)
	}
}
func (dataUser *user) notifyObserverLikePhoto(photoOwner *user) {
	for _, follower := range dataUser.followers {
		follower.updateLikePhoto(dataUser, photoOwner)
	}
}
func (dataUser *user) updateFollower(user *user) {
	dataUser.followers = append(dataUser.followers, user)
}
