package service

import (
	"bytes"
	"errors"
	"fmt"
	"goto2023/repository"
	"goto2023/structs"
	"image"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
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

	// use gocv
	// coverImg, err := captureFrame(videoPath, 0.25)
	// use ffmpeg
	coverImg, err := captureFrameFFmpeg(videoPath)

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

/*
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
*/

func captureFrameFFmpeg(filePath string) (i image.Image, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "png"}).
		WithOutput(buf).
		Run()
	if err != nil {
		return nil, err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func PublishList(userId int64) ([]structs.Video, error) {
	rawVideos, err := repository.QueryVideosByUser(userId)
	if err != nil {
		return nil, err
	}
	user, err := QueryUserInfo(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not find")
	}
	videos := make([]structs.Video, 0, len(rawVideos))

	for _, v := range rawVideos {
		videos = append(videos, structs.Video{
			Id:       v.Id,
			Author:   *user,
			PlayUrl:  serverAddr + v.PlayUrl,
			CoverUrl: serverAddr + v.CoverUrl,
			Title:    v.Title,
		})
	}

	return videos, nil
}
