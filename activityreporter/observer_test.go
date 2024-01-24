package activityreporter_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/activityreporter"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/activityreporter/mocks"
)

func TestUpdateLikePhoto(t *testing.T) {
	users := activityreporter.NewUsers()
	user := users.NewUser("bob")
	ownerPhoto := users.NewUser("jhon")
	ownerPhoto.UploadPhoto()
	MockO := new(mocks.Observer)
	MockO.On("UpdateLikePhoto", user, ownerPhoto).Return()
	user.UpdateFollower(MockO)
	user.Like(ownerPhoto)
	user.NotifyObserverLike(ownerPhoto)
	MockO.AssertNumberOfCalls(t, "UpdateLikePhoto", 1)
}

func TestUpdateUploadPhoto(t *testing.T) {
	users := activityreporter.NewUsers()
	user := users.NewUser("bob")
	MockO := new(mocks.Observer)
	MockO.On("UpdateUploadPhoto", user).Return()
	user.UpdateFollower(MockO)
	user.UploadPhoto()
	MockO.AssertNumberOfCalls(t, "UpdateUploadPhoto", 1)
}