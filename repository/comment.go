package repository

import "time"

func CountComment(videoId int64) (int64, error) {
	var count int64 = 0
	err := database.Model(&DBComment{}).Where("video_id = ?", videoId).Count(&count).Error
	return count, err
}

func QueryCommentsByVID(videoId int64) ([]*DBComment, error) {
	var comments = make([]*DBComment, 0)
	err := database.Model(&DBComment{}).Where("video_id = ?", videoId).Order("create_time DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func QueryCommentById(commentId int64) (*DBComment, error) {
	var comment DBComment
	err := database.Model(&comment).Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func AddComment(userId int64, videoId int64, content string) (*DBComment, error) {
	comment := DBComment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    content,
		CreateTime: time.Now(),
	}
	err := database.Model(&DBComment{}).Create(&comment).Error
	return &comment, err
}

func DeleteComment(commentId int64) error {
	return database.Model(&DBComment{}).Where("id = ?", commentId).Delete(&DBComment{}).Error
}
