package repository

import "time"

// gorm models

type DBUser struct {
	Id       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:passwd"`
}

func (DBUser) TableName() string {
	return "user"
}

type DBVideo struct {
	Id         int64     `gorm:"column:id"`
	AuthorId   int64     `gorm:"column:author"`
	Title      string    `gorm:"column:title"`
	PlayUrl    string    `gorm:"column:play_url"`
	CoverUrl   string    `gorm:"column:cover_url"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (DBVideo) TableName() string {
	return "video"
}

type DBFavorite struct {
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (DBFavorite) TableName() string {
	return "favorite"
}

type DBComment struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (DBComment) TableName() string {
	return "comment"
}
