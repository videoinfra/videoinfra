package service

import (
	"context"
	pb "videoinfra/ffmpeg/service/gen"
)

type FfmpegAPIServerInterface struct {
	pb.UnimplementedFfmpegAPIServer
}

func (s* FfmpegAPIServerInterface) GenerateCallback(ctx context.Context, request *pb.GeneratePlaybackRequest) (*pb.GeneratePlaybackResponse, error) {
	// TODO(yilkal): Run ffmpeg and return the path to m3u8 file. Currently, simply returns the arguments.
	return &pb.GeneratePlaybackResponse{OriginalVideoPath:request.VideoPath, M3U8PlaybackPath:request.VideoPath}, nil;
}