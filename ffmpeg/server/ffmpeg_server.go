package main

import (
	"fmt"
	"flag"
	"net"
	"log"
	"google.golang.org/grpc"
	"videoinfra/ffmpeg/service"
	pb "videoinfra/ffmpeg/service/gen"
)

func InitServer() *service.FfmpegAPIServerInterface {
 return &service.FfmpegAPIServerInterface{}
}

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
)

func main() {
	// Start the server.
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFfmpegAPIServer(grpcServer, InitServer())
	grpcServer.Serve(lis)
}