package utility_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"
	"github.com/stretchr/testify/assert"
)

func TestLikePhotoMessage(t *testing.T) {
	t.Run("should have return subject liked object photo", func(t *testing.T) {
		subject := "Bob"
		object := "jhon"
		target := subject + " liked " + object + "'s photo"
		message := utility.LikePhotoMessage(subject, object)
		assert.Equal(t, target, message)
	})
}
func TestUploadedPhotoMessage(t *testing.T) {
	t.Run("should have return subject uploaded photo", func(t *testing.T) {
		subject := "Bob"
		target := subject + " uploaded photo"
		message := utility.UploadedPhotoMessage(subject)
		assert.Equal(t, target, message)
	})
}
