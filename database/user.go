package database

import "git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/util"

type user struct {
	name       string
	photo      *photo
	followers  []*user
	followings []*user
	activity   []string
}

func (dataUser *user) Follow(user *user) error {
	errFollowing := dataUser.checkFollowing(user)
	if errFollowing != nil {
		return errFollowing
	}
	dataUser.followings = append(dataUser.followings, user)
	user.updateFollower(dataUser)
	return nil
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
func (dataUser *user) checkFollowing(user *user) error {
	for _, usr := range dataUser.followings {
		if usr == user {
			return util.ErrAlreadyFollowed
		}
	}
	return nil
}
