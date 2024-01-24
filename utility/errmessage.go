package utility

import "errors"

var (
	ErrKeyword          = errors.New("invalid keyword")
	ErrAlreadyFollowed  = errors.New("you already followed the user")
	ErrUploadePhoto     = errors.New("you cannot upload more than once")
	ErrAlreadyLike      = errors.New("you already liked the photo")
	ErrDontHavePhoto    = errors.New("you don't have a photo")
	ErrFollowThemselves = errors.New("a user cannot follow themselves")
)

func ErrUnknownUser(name string) error {
	return errors.New("unknown user " + name)
}

func ErrDoesntHavePhoto(name string) error {
	return errors.New(name + " doesn't have a photo")
}
func ErrUnableToLike(name string) error {
	return errors.New("unable to like " + name + "'s photo")

}
