package service

import (
	"nju-tube/repository"
	"nju-tube/structs"
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
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         v.Title,
			UploadDate:    v.CreateTime.Format("2006-1-2-15-04"),
		})
	}

	return videos, nextTime, nil
}
