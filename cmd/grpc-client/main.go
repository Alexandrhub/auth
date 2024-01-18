package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/alexandrhub/auth/internal/model"
	"github.com/alexandrhub/auth/pkg/access_v1"
)

var accessToken = flag.String("a", "", "access token")

const (
	servicePort = 50051
	authPrefix  = "Bearer"
	authHeader  = "Authorization"
)

func main() {
	flag.Parse()

	ctx := context.Background()
	md := metadata.New(
		map[string]string{
			authHeader: authPrefix + " " + *accessToken,
		},
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("failed to dial:", err)
	}

	cl := access_v1.NewAccessV1Client(conn)

	_, err = cl.Check(
		ctx, &access_v1.CheckRequest{
			EndpointAddress: model.ExamplePath,
		},
	)
	if err != nil {
		log.Fatal("failed to check:", err)
	}

	log.Println("access granted")
}
