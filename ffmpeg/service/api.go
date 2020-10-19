package service

import (
	"context"
	"bytes"
	"os"
	"os/exec"
	"log"
	"math/rand"
	"time"
	pb "videoinfra/ffmpeg/service/gen"
)

type FfmpegAPIServerInterface struct {
	
	pb.UnimplementedFfmpegAPIServer

}

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}

// TODO(yilkal) rename this rpc. I am sure that there are many better ones out there.
func (s* FfmpegAPIServerInterface) GenerateCallback(ctx context.Context, request *pb.GeneratePlaybackRequest) (*pb.GeneratePlaybackResponse, error) {
	// First create a random string.
	generated_path_name := String(10)

	// First step, create a folder for the ffmpeg outputs.
	err := os.Mkdir(generated_path_name, os.ModeDir)
	if err != nil {
		log.Printf("Error : ", err)
	}

	// Generate the playback
	cmd := exec.Command("ffmpeg" , "-i", request.VideoPath, "-c:a", "aac", "-strict", "experimental", "-c:v", "libx264", "-s", "240x320", "-aspect", "16:9", "-f", "hls", "-hls_list_size", "1000000", "-hls_time", "2", generated_path_name + "/240_out.m3u8");
	var out bytes.Buffer
	cmd.Stdout = &out

	ffmpeg_err := cmd.Run()
	if ffmpeg_err != nil {
   		log.Printf("Error :", ffmpeg_err)
	}

	log.Printf("in all caps: %q", out.String())

	// calculate m3u8 file path
	path, _ := os.Getwd()
	m3u8_file_path := path + "/" + generated_path_name + "/240_out.m3u8"

	return &pb.GeneratePlaybackResponse{OriginalVideoPath:request.VideoPath, M3U8PlaybackPath:m3u8_file_path}, nil;
}