package repository

import (
	"errors"
	"gorm.io/gorm"
)

func AddFavorite(userId int64, videoId int64) error {
	err := database.Model(&DBFavorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).First(&DBFavorite{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("database failed")
	}
	if err == nil {
		return errors.New("already has favorite item")
	}
	return database.Model(&DBFavorite{}).Create(&DBFavorite{
		UserId:  userId,
		VideoId: videoId,
	}).Error
}

func CancelFavorite(userId int64, videoId int64) error {
	return database.Model(&DBFavorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&DBFavorite{}).Error
}

func IsFavorite(userId int64, videoId int64) (bool, error) {
	err := database.Model(&DBFavorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).First(&DBFavorite{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

func QueryFavoriteVIDByUser(userId int64) ([]int64, error) {
	videos := make([]int64, 0)
	err := database.Model(&DBFavorite{}).Where("user_id = ?", userId).Select("video_id").Find(&videos).Error
	return videos, err
}

func CountFavorite(videoId int64) (int64, error) {
	var count int64 = 0
	err := database.Model(&DBFavorite{}).Where("video_id = ?", videoId).Count(&count).Error
	return count, err
}
