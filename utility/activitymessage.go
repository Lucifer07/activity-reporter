package utility

var (
	LikeSelfPhoto = "You liked your photo"
	Subject       = "You"
)

func LikePhotoMessage(subject string, object string) string {
	return subject + " liked " + object + "'s photo"

}
func UploadedPhotoMessage(subject string) string {
	return subject + " uploaded photo"

}
