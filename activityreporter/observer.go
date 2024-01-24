package activityreporter

import "git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"

type observer interface {
	updateLikePhoto(userObject *user, ownerPhoto *user)
	updateUploadPhoto(userObject *user)
}

func (userSubject *user) updateUploadPhoto(userObject *user) {
	userSubject.activity = append(userSubject.activity, utility.UploadedPhotoMessage(userObject.name))
}
func (userSubject *user) updateLikePhoto(userObject *user, ownerPhoto *user) {
	userSubject.activity = append(userSubject.activity, utility.LikePhotoMessage(userObject.name, ownerPhoto.name))
}

type publisher interface {
	updateFollower(observer)
	notifyObserverLike(ownerPhoto *user)
	notifyObserverUploadPhoto()
}

func (userSubject *user) notifyObserverUploadPhoto() {
	for _, follower := range userSubject.followers {
		follower.updateUploadPhoto(userSubject)
	}
}
func (userSubject *user) notifyObserverLike(ownerPhoto *user) {
	if userSubject != ownerPhoto {
		ownerPhoto.activity = append(ownerPhoto.activity, utility.LikePhotoMessage(userSubject.name, "your"))
	}
	for _, follower := range userSubject.followers {
		if follower != ownerPhoto {
			follower.updateLikePhoto(userSubject, ownerPhoto)
		}
	}
}
func (userSubject *user) updateFollower(user *user) {
	userSubject.followers = append(userSubject.followers, user)
}
