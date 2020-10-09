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
	"google.golang.org/grpc"
	pb "videoinfra/videos/service/gen"
)

func InitServer(videos_db_handle *gorm.DB) *service.VideoAPIServer {
 return &service.VideoAPIServer{DB_context: storage.VideoDBContext{Videos_db_handle: videos_db_handle}}
 
}

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

func main() {
	// Open connection to DB.
	data_source_name := "user=postgres password=gorm dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(data_source_name), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	// Start the server.
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterVideoAPIServer(grpcServer, InitServer(db))
	grpcServer.Serve(lis)
}