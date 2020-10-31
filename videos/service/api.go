package service

import (
	"context"
	"fmt"
	"log"
	"time"
	ffmpeg "videoinfra/ffmpeg/service"
)

type VideoAPIServerImpl struct {
	UnimplementedVideoAPIServer
	DbContext    VideoDBContext
	FfmpegClient ffmpeg.FfmpegAPIClient
}

func (s* VideoAPIServerImpl) GeneratePlayback(videoPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Response, err := s.FfmpegClient.GeneratePlayback(ctx, &ffmpeg.GeneratePlaybackRequest{VideoPath: videoPath})
	if err != nil {
		log.Fatalf("%v.GenerateCallback(_) = _, %v: ", s.FfmpegClient, err)
	}
	log.Println("Original video path = " + Response.OriginalVideoPath)
	log.Println("Playback video path = " + Response.M3U8PlaybackPath)

	return nil;
}

func (s* VideoAPIServerImpl) CreatePlayback(ctx context.Context, request *CreatePlaybackRequest) (*CreatePlaybackResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	generateResponse, err := s.FfmpegClient.GeneratePlayback(ctx, &ffmpeg.GeneratePlaybackRequest{VideoPath: request.VideoPath})
	if err != nil {
		return &CreatePlaybackResponse{}, err
	}
	playbackId := generateResponse.M3U8PlaybackPath
	draftPlayback := Playback{VideoId: request.VideoId, PlaybackId: playbackId,
		Policy: request.Policy}
	err = s.DbContext.CreatePlayback(draftPlayback)
	return &CreatePlaybackResponse{PlaybackId: playbackId}, err
}

func (s* VideoAPIServerImpl) CreateVideo(ctx context.Context, request *CreateVideoRequest) (*CreateVideoResponse, error) {
	draftVideo := Video{VideoId: "", Title: request.Title, AccountId: request.AccountId, Filepath: "",
		VideoState: VideoState_PROCESSING, VideoProperties: &VideoProperties{}}
	videoId, err := s.DbContext.CreateVideo(draftVideo)
	response := &CreateVideoResponse{VideoId: videoId}
	return response, err
}

func (s* VideoAPIServerImpl) GetVideo(ctx context.Context, request *GetVideoRequest) (*GetVideoResponse, error) {
	video, err := s.DbContext.GetVideo(request.VideoId)
	return &GetVideoResponse{Video: &video}, err
}

func (s* VideoAPIServerImpl) GetVideosByAccount(ctx context.Context, request *GetVideosByAccountRequest) (*GetVideosByAccountResponse, error) {
	videos, err := s.DbContext.GetVideosByAccount(request.AccountId)
	returnedVideos := []*Video{}
	for _, video := range videos {
		returnedVideo := video
		returnedVideos = append(returnedVideos, &returnedVideo)
	}

	fmt.Println("RPC handler return")
	for _, video := range returnedVideos {
		fmt.Println(video.Title)
	}

	return &GetVideosByAccountResponse{Videos: returnedVideos}, err
}