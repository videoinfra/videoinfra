package service

import (
	"context"
	pb "videoinfra/videos/service/gen"
	storage "videoinfra/videos/storage"
	ffmpeg "videoinfra/ffmpeg/service/gen"
	"time"
	"log"
)

type VideoAPIServer struct {
	pb.UnimplementedVideoAPIServer
	DB_context storage.VideoDBContext
	Ffmpeg_client ffmpeg.FfmpegAPIClient
}

func (s* VideoAPIServer) GeneratePlayback(video_path string) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Response, err := s.Ffmpeg_client.GeneratePlayback(ctx, &ffmpeg.GeneratePlaybackRequest{VideoPath: video_path})
	if err != nil {
		log.Fatalf("%v.GenerateCallback(_) = _, %v: ", s.Ffmpeg_client, err)
	}
	log.Println("Orignial video path = " + Response.OriginalVideoPath)
	log.Println("Playback video path = " + Response.M3U8PlaybackPath)

	return nil;
}

func (s* VideoAPIServer) UploadVideo(ctx context.Context, request *pb.UploadVideoRequest) (*pb.UploadVideoResponse, error) {
	s.GeneratePlayback(request.Title)
	video_draft := storage.VideoDraft{VideoId: "vlabc", Title: request.Title, AccountId: "account", Filepath: "filepath"}
	return &pb.UploadVideoResponse{}, s.DB_context.CreateVideo(video_draft)
}