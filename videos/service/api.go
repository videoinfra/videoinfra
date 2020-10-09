package service

import (
	"context"
	pb "videoinfra/videos/service/gen"
	storage "videoinfra/videos/storage"
)

type VideoAPIServer struct {
	pb.UnimplementedVideoAPIServer
	DB_context storage.VideoDBContext
}

func (s* VideoAPIServer) UploadVideo(ctx context.Context, request *pb.UploadVideoRequest) (*pb.UploadVideoResponse, error) {
	video_draft := storage.VideoDraft{VideoId: "vlabc", Title: request.Title, AccountId: "account", Filepath: "filepath"}
	return &pb.UploadVideoResponse{}, s.DB_context.CreateVideo(video_draft)
}