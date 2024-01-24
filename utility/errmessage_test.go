package utility_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/assignment-activity-reporter/-/tree/dev/utility"
	"github.com/stretchr/testify/assert"
)

func TestErrUnknownUser(t *testing.T) {
	t.Run("should have return unknown user subject name", func(t *testing.T) {
		subject := "Bob"
		target := errors.New("unknown user " + subject)
		message := utility.ErrUnknownUser(subject)
		assert.Equal(t, target, message)
	})
}
func TestErrDoesntHavePhoto(t *testing.T) {
	t.Run("should have return subject name doesn't have a photo", func(t *testing.T) {
		subject := "Bob"
		target := errors.New(subject + " doesn't have a photo")
		message := utility.ErrDoesntHavePhoto(subject)
		assert.Equal(t, target, message)
	})
}
func TestErrUnableToLike(t *testing.T) {
	t.Run("should have return unable to like subject name's photo", func(t *testing.T) {
		subject := "Bob"
		target := errors.New("unable to like " + subject + "'s photo")
		message := utility.ErrUnableToLike(subject)
		assert.Equal(t, target, message)
	})
}
