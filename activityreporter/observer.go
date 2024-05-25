package activityreporter

import "github.com/Lucifer07/activity-reporter/utility"

type Observer interface {
	UpdateLikePhoto(userObject *User, ownerPhoto *User)
	UpdateUploadPhoto(userObject *User)
}

func (userSubject *User) UpdateUploadPhoto(userObject *User) {
	userSubject.activity = append(userSubject.activity, utility.UploadedPhotoMessage(userObject.name))
}
func (userSubject *User) UpdateLikePhoto(userObject *User, ownerPhoto *User) {
	userSubject.activity = append(userSubject.activity, utility.LikePhotoMessage(userObject.name, ownerPhoto.name))
}

type publisher interface {
	UpdateFollower(Observer)
	NotifyObserverLike(ownerPhoto *User)
	NotifyObserverUploadPhoto()
}

func (userSubject *User) NotifyObserverUploadPhoto() {
	for _, follower := range userSubject.followers {
		follower.UpdateUploadPhoto(userSubject)
	}
}
func (userSubject *User) NotifyObserverLike(ownerPhoto *User) {
	if userSubject != ownerPhoto {
		ownerPhoto.activity = append(ownerPhoto.activity, utility.LikePhotoMessage(userSubject.name, "your"))
	}
	for _, follower := range userSubject.followers {
		if follower != ownerPhoto {
			follower.UpdateLikePhoto(userSubject, ownerPhoto)
		}
	}
}
func (userSubject *User) UpdateFollower(o Observer) {
	userSubject.followers = append(userSubject.followers, o)
}
