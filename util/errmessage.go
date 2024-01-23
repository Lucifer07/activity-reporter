package util

import "errors"

var (
	ErrKeyword          = errors.New("invalid keyword")
	ErrAlreadyFollowed  = errors.New("you already followed the user")    //1
	ErrUploadePhoto     = errors.New("you cannot upload more than once") //1
	ErrAlreadyLike      = errors.New("you already liked the photo") 
	ErrDontHavePhoto    = errors.New("you don't have a photo")          //1
	ErrFollowThemselves = errors.New("a user cannot follow themselves") //1
)

func ErrUnknownUser(name string) error {
	return errors.New("unknown user " + name)
}

func ErrDoesntHavePhoto(name string) error { //1
	return errors.New(name + " doesn't have a photo")
}
func ErrUnableToLike(name string) error { //1
	return errors.New("unable to like " + name + "'s photo")

}
