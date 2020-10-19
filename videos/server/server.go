package main

import (
	"fmt"
	"flag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"net"
	"log"
	"videoinfra/videos/storage"
	"videoinfra/videos/service"
	ffmpeg "videoinfra/ffmpeg/service/gen"
	"google.golang.org/grpc"
	pb "videoinfra/videos/service/gen"
)

func InitServer(videos_db_handle *gorm.DB, client ffmpeg.FfmpegAPIClient) *service.VideoAPIServer {
 return &service.VideoAPIServer{DB_context: storage.VideoDBContext{Videos_db_handle: videos_db_handle}, Ffmpeg_client: client}
}

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
	ffmpegServerAddr  = flag.String("ffmpeg_server_addr", "localhost:10001", "The server address in the format of host:port")
)

func main() {
	// Open connection to DB.
	data_source_name := "user=postgres password=gorm dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(data_source_name), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	// Create client of ffmpeg service. I wonder if this should be created per request?
	// Will there be race conditions? If multiple request GeneratePlayback requests are made at the same time?
	var opts_ffmpeg_client []grpc.DialOption
	opts_ffmpeg_client = append(opts_ffmpeg_client, grpc.WithInsecure())
	conn, _ := grpc.Dial(*ffmpegServerAddr, opts_ffmpeg_client...)
	defer conn.Close()
	ffmpeg_client := ffmpeg.NewFfmpegAPIClient(conn)

	// Start the server.
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterVideoAPIServer(grpcServer, InitServer(db, ffmpeg_client))
	grpcServer.Serve(lis)
}