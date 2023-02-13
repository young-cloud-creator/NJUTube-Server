package repository

import (
	"errors"
	"time"
)

func QueryVideoById(id int64) (*DBVideo, error) {
	var video DBVideo
	err := database.Model(&video).Where("id = ?", id).Find(&video).Error
	if err != nil {
		return nil, errors.New("video not exists")
	}
	return &video, nil
}

func AddVideo(userId int64, title string, playUrl string, coverUrl string) (*DBVideo, error) {
	video := DBVideo{
		AuthorId:   userId,
		Title:      title,
		PlayUrl:    playUrl,
		CoverUrl:   coverUrl,
		CreateTime: time.Now(),
	}
	err := database.Model(&video).Create(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}
