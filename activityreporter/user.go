package activityreporter

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"
)

type user struct {
	name       string
	photo      *photo
	followers  []*user
	followings []*user
	activity   []string
}

func (userSubject *user) Follow(user *user) error {
	if userSubject == user {
		return utility.ErrFollowThemselves
	}
	errFollowing := userSubject.checkFollowing(user)
	if errFollowing != nil {
		return errFollowing
	}
	userSubject.followings = append(userSubject.followings, user)
	user.updateFollower(userSubject)
	return nil
}

func (userSubject *user) checkFollowing(user *user) error {
	for _, usr := range userSubject.followings {
		if usr == user {
			return utility.ErrAlreadyFollowed
		}
	}
	return nil
}
func (userSubject *user) UploadPhoto() (*photo, error) {
	if userSubject.photo != nil {
		return nil, utility.ErrUploadePhoto
	}
	userSubject.photo = &photo{}
	userSubject.activity = append(userSubject.activity, utility.UploadedPhotoMessage("You"))
	userSubject.notifyObserverUploadPhoto()
	return userSubject.photo, nil
}
func (userSubject *user) Like(user *user) error {
	if user == userSubject && user.photo == nil {
		return utility.ErrDontHavePhoto
	}
	if user != userSubject && user.photo == nil {
		return utility.ErrDoesntHavePhoto(user.name)
	}
	if (user == userSubject && user.photo != nil) || (user != userSubject && user.photo != nil) {
		errLike := user.checkLikes(userSubject)
		if errLike != nil {
			return errLike
		}
	}
	if user != userSubject && user.photo != nil {
		isNotFollower := user.checkFollower(userSubject)
		if isNotFollower != nil {
			return isNotFollower
		}
		userSubject.activity = append(userSubject.activity, utility.LikePhotoMessage(utility.Subject, user.name))
	}
	if user == userSubject && user.photo != nil {
		userSubject.activity = append(userSubject.activity, utility.LikeSelfPhoto)
	}
	user.photo.likes = append(user.photo.likes, userSubject)
	userSubject.notifyObserverLike(user)
	return nil

}
func (userSubject *user) checkFollower(user *user) error {
	for _, usr := range userSubject.followers {
		if usr == user {
			return nil
		}
	}
	return utility.ErrUnableToLike(userSubject.name)
}
func (userSubject *user) checkLikes(user *user) error {
	for _, usr := range userSubject.photo.likes {
		if usr == user {
			return utility.ErrAlreadyLike
		}
	}
	return nil
}
func (userSubject *user) Activity() []string {
	return userSubject.activity
}
func (userSubject *user) Name() string {
	return userSubject.name
}
