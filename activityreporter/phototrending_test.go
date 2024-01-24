package activityreporter_test

import (
	"strconv"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/activityreporter"
	"github.com/stretchr/testify/assert"
)

func TestTrending(t *testing.T) {
	t.Run("should have return photo trending type", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")
		target := "Bob:2 bob:1"

		bob.Follow(Bob)
		Bob.Follow(bob)
		bob.UploadPhoto()
		Bob.UploadPhoto()
		bob.Like(Bob)
		Bob.Like(bob)
		Bob.Like(Bob)
		trending := users.Trending()

		result := trending[0].Name + ":" + strconv.Itoa(trending[0].Like) + " " + trending[1].Name + ":" + strconv.Itoa(trending[1].Like)
		assert.Equal(t, target, result)
	})
}
