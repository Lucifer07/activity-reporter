package activityreporter_test

import (
	"strconv"
	"testing"

	"github.com/Lucifer07/activity-reporter/activityreporter"
	"github.com/stretchr/testify/assert"
)

func TestTrending(t *testing.T) {
	t.Run("should have return photo trending type", func(t *testing.T) {
		users := activityreporter.NewUsers()
		bob := users.NewUser("bob")
		Bob := users.NewUser("Bob")
		ras := users.NewUser("ras")
		gas := users.NewUser("gas")
		target := "Bob:2 bob:1"

		bob.Follow(Bob)
		Bob.Follow(bob)
		ras.Follow(bob)
		bob.Follow(ras)
		gas.Follow(Bob)
		Bob.Follow(gas)
		bob.UploadPhoto()
		Bob.UploadPhoto()
		gas.UploadPhoto()
		ras.UploadPhoto()
		bob.Like(Bob)
		Bob.Like(bob)
		Bob.Like(Bob)
		trending := users.Trending()

		result := trending[0].Name + ":" + strconv.Itoa(trending[0].Like) + " " + trending[1].Name + ":" + strconv.Itoa(trending[1].Like)
		assert.Equal(t, target, result)
	})
}
