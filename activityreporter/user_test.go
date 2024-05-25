package activityreporter_test

import (
	"testing"

	"github.com/Lucifer07/activity-reporter/activityreporter"
	"github.com/Lucifer07/activity-reporter/utility"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	t.Run("should return nil if successed followed", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")

		err := bob.Follow(Bob)

		assert.Equal(t, nil, err)
	})
	t.Run("should return error already followed the user if you already followed the user", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")
		target := utility.ErrAlreadyFollowed

		bob.Follow(Bob)
		err := bob.Follow(Bob)

		assert.Equal(t, target, err)
	})
	t.Run("should return error cant followed themself if user try to follow themself", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		target := utility.ErrFollowThemselves

		err := bob.Follow(bob)

		assert.Equal(t, target, err)
	})
}
func TestUploadPhoto(t *testing.T) {
	t.Run("should be return photo if successed upload photo", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")

		Bobsphoto, _ := Bob.UploadPhoto()
		bobsphoto, _ := bob.UploadPhoto()

		assert.IsType(t, Bobsphoto, bobsphoto)
	})
	t.Run("should return error cannot upload more than once if user try upload more than once", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		target := utility.ErrUploadePhoto

		bob.UploadPhoto()
		_, err := bob.UploadPhoto()

		assert.Equal(t, target, err)
	})
}
func TestLike(t *testing.T) {
	t.Run("should return nil if successed like some photo", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")

		bob.Follow(Bob)
		Bob.UploadPhoto()
		err := bob.Like(Bob)

		assert.Equal(t, nil, err)
	})
	t.Run("should return already liked the photo if user try to like himself photo twice", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		target := utility.ErrAlreadyLike
		bob.UploadPhoto()
		bob.Like(bob)
		err := bob.Like(bob)

		assert.Equal(t, target, err)
	})
	t.Run("should return already liked the photo if user try to like some photo twice", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")
		target := utility.ErrAlreadyLike

		bob.Follow(Bob)
		Bob.UploadPhoto()
		bob.Like(Bob)
		err := bob.Like(Bob)

		assert.Equal(t, target, err)
	})
	t.Run("should return you don't have a photo if user try to like him self but dont have a photo", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		target := utility.ErrDontHavePhoto

		err := bob.Like(bob)

		assert.Equal(t, target, err)
	})
	t.Run("should return unable to like some photo when user try to like some photo but he not yet follow that user", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")
		target := utility.ErrUnableToLike("Bob")

		Bob.UploadPhoto()
		err := bob.Like(Bob)

		assert.Equal(t, target, err)
	})
	t.Run("should return user doesn't have a photo when user follower try to like the user but user dont have photo", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")
		target := utility.ErrDoesntHavePhoto("Bob")

		bob.Follow(Bob)
		bob.Like(Bob)
		err := bob.Like(Bob)

		assert.Equal(t, target, err)
	})

}
func TestCheckActivity(t *testing.T) {
	users := activityreporter.NewUsers()
	user := users.NewUser("bob")
	target := []string{"You uploaded photo"}
	user.UploadPhoto()
	userActivity := user.Activity()
	assert.Equal(t, target, userActivity)
}
func TestName(t *testing.T) {
	users := activityreporter.NewUsers()
	target := "bob"
	user := users.NewUser("bob")
	userName := user.Name()
	assert.Equal(t, target, userName)
}
