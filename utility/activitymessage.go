package utility

var (
	LikeSelfPhoto = "You liked your photo"
	Subject       = "You"
)

func LikePhotoMessage(subject string, object string) string {
	if object != "your" {
		return subject + " liked " + object + "'s photo"
	}
	return subject + " liked " + object + " photo"
}
func UploadedPhotoMessage(subject string) string {
	return subject + " uploaded photo"

}
