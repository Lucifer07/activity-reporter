package activityreporter

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"
)

type User struct {
	name       string
	photo      *photo
	followers  []Observer
	followings []*User
	activity   []string
}

func (userSubject *User) Follow(user *User) error {
	if userSubject == user {
		return utility.ErrFollowThemselves
	}
	errFollowing := userSubject.checkFollowing(user)
	if errFollowing != nil {
		return errFollowing
	}
	userSubject.followings = append(userSubject.followings, user)
	user.UpdateFollower(userSubject)
	return nil
}

func (userSubject *User) checkFollowing(user *User) error {
	for _, usr := range userSubject.followings {
		if usr == user {
			return utility.ErrAlreadyFollowed
		}
	}
	return nil
}
func (userSubject *User) UploadPhoto() (*photo, error) {
	if userSubject.photo != nil {
		return nil, utility.ErrUploadePhoto
	}
	userSubject.photo = &photo{}
	userSubject.activity = append(userSubject.activity, utility.UploadedPhotoMessage(utility.Subject))
	userSubject.NotifyObserverUploadPhoto()
	return userSubject.photo, nil
}
func (userSubject *User) Like(user *User) error {
	if user != userSubject {
		return userSubject.likeOthersPhoto(user)
	}
	return userSubject.likeOwnPhoto()
}
func (userSubject *User) likeOwnPhoto() error {
	if userSubject.photo != nil {
		errLike := userSubject.checkLikes(userSubject)
		if errLike != nil {
			return errLike
		}
		userSubject.activity = append(userSubject.activity, utility.LikeSelfPhoto)
		userSubject.photo.likes = append(userSubject.photo.likes, userSubject)
		userSubject.NotifyObserverLike(userSubject)
		return nil
	}
	return utility.ErrDontHavePhoto
}
func (userSubject *User) likeOthersPhoto(user *User) error {
	isNotFollower := user.checkFollower(userSubject)
	if isNotFollower != nil {
		return isNotFollower
	}
	if user.photo != nil {
		errLike := user.checkLikes(userSubject)
		if errLike != nil {
			return errLike
		}
		userSubject.activity = append(userSubject.activity, utility.LikePhotoMessage(utility.Subject, user.name))
		user.photo.likes = append(user.photo.likes, userSubject)
		userSubject.NotifyObserverLike(user)
		return nil
	}
	return utility.ErrDoesntHavePhoto(user.name)
}
func (userSubject *User) checkFollower(user *User) error {
	for _, usr := range userSubject.followers {
		if usr == user {
			return nil
		}
	}
	return utility.ErrUnableToLike(userSubject.name)
}
func (userSubject *User) checkLikes(user *User) error {
	for _, usr := range userSubject.photo.likes {
		if usr == user {
			return utility.ErrAlreadyLike
		}
	}
	return nil
}
func (userSubject *User) Activity() []string {
	return userSubject.activity
}
func (userSubject *User) Name() string {
	return userSubject.name
}
