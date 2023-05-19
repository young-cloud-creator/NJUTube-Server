package service

import (
	"nju-tube/repository"
	"nju-tube/structs"
)

func DoFavoriteAction(userId int64, videoId int64, isFavorite bool) error {
	if isFavorite {
		return repository.AddFavorite(userId, videoId)
	}
	return repository.CancelFavorite(userId, videoId)
}

func FavoriteList(userId int64, selfId int64) ([]structs.Video, error) {
	vidArray, err := repository.QueryFavoriteVIDByUser(userId)
	if err != nil {
		return nil, err
	}
	videos := make([]structs.Video, 0, len(vidArray))

	for _, vid := range vidArray {
		dbVideo, err := repository.QueryVideoById(vid)
		if err != nil {
			continue
		}
		user, err := QueryUserInfo(dbVideo.AuthorId)
		if err != nil {
			continue
		}
		if user == nil {
			continue
		}

		isFavorite, _ := repository.IsFavorite(selfId, dbVideo.Id)
		favoriteCount, _ := repository.CountFavorite(dbVideo.Id)
		commentCount, _ := repository.CountComment(dbVideo.Id)
		videos = append(videos, structs.Video{
			Id:            dbVideo.Id,
			Author:        *user,
			PlayUrl:       dbVideo.PlayUrl,
			CoverUrl:      dbVideo.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         dbVideo.Title,
			UploadDate:    dbVideo.CreateTime.Format("2006-1-2-15-04"),
		})
	}

	return videos, nil
}
