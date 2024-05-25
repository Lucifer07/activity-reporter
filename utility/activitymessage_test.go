package utility_test

import (
	"testing"

	"github.com/Lucifer07/activity-reporter/utility"
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
func TestLikeOwnPhotoMessage(t *testing.T) {
	t.Run("should have return subject liked object photo", func(t *testing.T) {
		subject := "Bob"
		object := "your"
		target := subject + " liked " + "your" + " photo"
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
