package util

var (
	LikeSelfPhoto = "You liked your photo"
)

func LikePhotoMessage(subject string, object string) string {
	return subject + " liked " + object + " photo"

}
func UploadedPhotoMessage(subject string) string {
	return subject + " uploaded photo"

}
