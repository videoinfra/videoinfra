package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	ffmpeg "videoinfra/ffmpeg/service"
)

func InitServer(videosDbHandle *sql.DB, client ffmpeg.FfmpegAPIClient) *VideoAPIServerImpl {
	return &VideoAPIServerImpl{DbContext: VideoDBContext{VideosDBHandle: videosDbHandle}, FfmpegClient: client}
}

var (
	tls              = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile         = flag.String("cert_file", "", "The TLS cert file")
	keyFile          = flag.String("key_file", "", "The TLS key file")
	jsonDBFile       = flag.String("json_db_file", "", "A json file containing a list of features")
	port             = flag.Int("port", 10000, "The server port")
	ffmpegServerAddr = flag.String("ffmpeg_server_addr", "localhost:10001", "The server address in the format of host:port")
)

func main() {
	// Open connection to DB.
	psqlInfo := "host=localhost port=5432 user=postgres password=gorm dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create client of ffmpeg service. I wonder if this should be created per request?
	// Will there be race conditions? If multiple request GeneratePlayback requests are made at the same time?
	var optsFfmpegClient []grpc.DialOption
	optsFfmpegClient = append(optsFfmpegClient, grpc.WithInsecure())
	conn, _ := grpc.Dial(*ffmpegServerAddr, optsFfmpegClient...)
	defer conn.Close()
	ffmpegClient := ffmpeg.NewFfmpegAPIClient(conn)

	// Start the server.
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Println("Starting server on port", *port)

	RegisterVideoAPIServer(grpcServer, InitServer(db, ffmpegClient))
	grpcServer.Serve(lis)
}
