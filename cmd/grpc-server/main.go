package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/alexandrhub/auth/internal/config"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

type server struct {
	pb.UnimplementedUserV1Server
}

func main() {
	conf := config.MustConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Grpc.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserV1Server(s, &server{})

	log.Printf("Starting gRPC server... on : %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	s.GracefulStop()

	log.Println("Server shutdown gracefully")
}
