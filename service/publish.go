package service

import (
	"bytes"
	"errors"
	"gocv.io/x/gocv"
	"goto2023/repository"
	"goto2023/structs"
	"image"
	"image/png"
	"os"
)

const PublicDir = "public/"
const VideoDir = PublicDir + "videos/"
const CoverDir = PublicDir + "covers/"

const serverAddr = "http://192.168.3.99:8080/"

// PublishAction capture the video cover and store video info to database
func PublishAction(title string, videoName string, userId int64) error {
	videoPath := VideoDir + videoName
	coverPath := CoverDir + videoName + ".png"

	// capture the video cover and save to coverPath
	coverImg, err := captureFrame(videoPath, 0.25)
	if err != nil {
		return errors.New("cannot capture video cover")
	}
	buff := new(bytes.Buffer)
	err = png.Encode(buff, coverImg)
	if err != nil {
		return errors.New("cannot save video cover")
	}
	err = os.WriteFile(coverPath, buff.Bytes(), os.ModePerm)
	if err != nil {
		return errors.New("cannot save video cover")
	}

	_, err = repository.AddVideo(userId, title, videoPath, coverPath)
	if err != nil {
		return errors.New("cannot save video info to database")
	}
	return nil
}

// CaptureFrame
// @param percent float64 "percent of video, should be less than 1"
func captureFrame(filePath string, percent float64) (i image.Image, err error) {
	//load video
	vc, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return i, err
	}

	frames := vc.Get(gocv.VideoCaptureFrameCount)

	frames = percent * frames

	// Set Video frames
	vc.Set(gocv.VideoCapturePosFrames, frames)

	img := gocv.NewMat()

	vc.Read(&img)

	imageObject, err := img.ToImage()
	if err != nil {
		return i, err
	}
	return imageObject, err
}

func PublishList(userId int64) ([]structs.Video, error) {
	videos := make([]structs.Video, 0, 10)
	rawVideos, err := repository.QueryVideosByUser(userId)
	if err != nil {
		return videos, err
	}
	user, err := QueryUserInfo(userId)
	if err != nil {
		return videos, err
	}
	if user == nil {
		return videos, errors.New("user not find")
	}

	for _, v := range rawVideos {
		videos = append(videos, structs.Video{
			Author:   *user,
			PlayUrl:  serverAddr + v.PlayUrl,
			CoverUrl: serverAddr + v.CoverUrl,
			Title:    v.Title,
		})
	}

	return videos, nil
}
