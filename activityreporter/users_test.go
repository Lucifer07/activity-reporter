package activityreporter_test

import (
	"testing"

	"github.com/Lucifer07/activity-reporter/activityreporter"
	"github.com/Lucifer07/activity-reporter/utility"
	"github.com/stretchr/testify/assert"
)

func TestNewUsers(t *testing.T) {
	t.Run("should have return users", func(t *testing.T) {
		target := &activityreporter.Users{}
		users := activityreporter.NewUsers()

		assert.Equal(t, target, users)
	})
}
func TestNewUser(t *testing.T) {
	t.Run("should have return user", func(t *testing.T) {
		users := activityreporter.NewUsers()
		user := users.NewUser("bob")

		target, _ := users.CheckUser("bob")

		assert.Equal(t, target, user)
	})
}
func TestCheckUser(t *testing.T) {
	t.Run("should have return user if succcessed", func(t *testing.T) {
		users := activityreporter.NewUsers()
		target := users.NewUser("bob")

		user, _ := users.CheckUser("bob")

		assert.Equal(t, target, user)
	})
	t.Run("should return error unknown user if failed", func(t *testing.T) {
		target := utility.ErrUnknownUser("charlie")
		users := activityreporter.NewUsers()

		users.NewUser("bob")
		_, err := users.CheckUser("charlie")

		assert.Equal(t, target, err)
	})
}
