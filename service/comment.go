package service

import (
	"goto2023/repository"
	"goto2023/structs"
)

func AddComment(userId int64, videoId int64, content string) (*structs.Comment, error) {
	rawComment, err := repository.AddComment(userId, videoId, content)
	if err != nil {
		return nil, err
	}
	user, err := QueryUserInfo(rawComment.UserId)
	if err != nil {
		return nil, err
	}
	comment := structs.Comment{
		Id:         rawComment.Id,
		User:       *user,
		Content:    rawComment.Content,
		CreateDate: rawComment.CreateTime.Format("01-02"),
	}
	return &comment, nil
}

func DeleteComment(userId int64, videoId int64, commentId int64) (bool, string) {
	comment, err := repository.QueryCommentById(commentId)
	if err != nil {
		return false, "Database Failed"
	}
	if comment.UserId != userId {
		return false, "User Does Not Have Permission"
	}
	if comment.VideoId != videoId {
		return false, "Wrong Video ID"
	}
	err = repository.DeleteComment(commentId)
	if err != nil {
		return false, "Database Failed"
	}
	return true, ""
}

func CommentList(videoId int64) []structs.Comment {
	rawComments, err := repository.QueryCommentsByVID(videoId)
	if err != nil {
		return nil
	}
	comments := make([]structs.Comment, 0, len(rawComments))
	for _, comment := range rawComments {
		user, err := QueryUserInfo(comment.UserId)
		if err != nil {
			continue
		}
		comments = append(comments, structs.Comment{
			Id:         comment.Id,
			User:       *user,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format("01-02"),
		})
	}
	return comments
}
