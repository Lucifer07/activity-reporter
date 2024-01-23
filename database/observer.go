package database

import "git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/util"

type observer interface {
	updateLike(*user) //0
	updateUploadPhoto(*user)
}

func (dataUser *user) updateUploadPhoto(user *user) {
	dataUser.activity = append(dataUser.activity, util.UploadedPhotoMessage(user.name))
}

type publisher interface {
	updateFollower(observer)
	notifyObserverLike()
	notifyObserverUploadPhoto()
}

func (dataUser *user) notifyObserverUploadPhoto() {
	for _, follower := range dataUser.followers {
		follower.updateUploadPhoto(dataUser)
	}
}
func (dataUser *user) updateFollower(user *user) {
	dataUser.followers = append(dataUser.followers, user)
}
