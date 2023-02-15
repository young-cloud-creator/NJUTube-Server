package repository

func CountComment(videoId int64) (int64, error) {
	var count int64 = 0
	err := database.Model(&DBComment{}).Where("video_id = ?", videoId).Count(&count).Error
	return count, err
}
