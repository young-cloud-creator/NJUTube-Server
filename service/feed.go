package service

import (
	"goto2023/repository"
	"goto2023/structs"
	"time"
)

func GenerateFeed(latestTime int64, userId int64) ([]structs.Video, int64, error) {
	dateTime := time.Unix(latestTime, 0).Local()
	rawVideos, err := repository.QueryVideosByTime(dateTime, 10)
	if err != nil {
		return nil, -1, nil
	}
	videos := make([]structs.Video, 0, len(rawVideos))

	nextTime := time.Now().Unix()
	if len(rawVideos) > 0 {
		nextTime = rawVideos[len(rawVideos)-1].CreateTime.Unix()
	}

	for _, v := range rawVideos {
		if v == nil {
			continue
		}

		user, err := QueryUserInfo(v.AuthorId)
		if err != nil {
			continue
		}

		isFavorite, _ := repository.IsFavorite(userId, v.Id)
		favoriteCount, _ := repository.CountFavorite(v.Id)
		commentCount, _ := repository.CountComment(v.Id)
		videos = append(videos, structs.Video{
			Id:            v.Id,
			Author:        *user,
			PlayUrl:       serverAddr + v.PlayUrl,
			CoverUrl:      serverAddr + v.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         v.Title,
		})
	}

	return videos, nextTime, nil
}
