package app

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/alexandrhub/auth/internal/api/auth"
	"github.com/alexandrhub/auth/internal/client/db/pg"
	"github.com/alexandrhub/auth/internal/client/db/transaction"
	authRepository "github.com/alexandrhub/auth/internal/repository/auth"
	authService "github.com/alexandrhub/auth/internal/service/auth"
	pb "github.com/alexandrhub/auth/pkg/user_v1"
)

func RunWithoutDI(ctx context.Context) (*App, error) {
	a := &App{}
	a.serviceProvider = newServiceProvider()
	// a.serviceProvider.Config = config.MustConfig()

	dbClient, err := pg.NewDBClient(ctx, a.serviceProvider.pgConfig.DSN())
	if err != nil {
		log.Fatal("failed to create db client: ", err)
	}
	a.serviceProvider.dbClient = dbClient
	a.serviceProvider.txManager = transaction.NewTransactionManager(dbClient.DB())

	a.serviceProvider.authRepository = authRepository.NewRepository(dbClient)
	a.serviceProvider.authService = authService.NewService(
		a.serviceProvider.authRepository,
		a.serviceProvider.txManager,
	)

	a.serviceProvider.authImpl = auth.NewImplementation(a.serviceProvider.authService)

	a.grpcServer = grpc.NewServer()

	reflection.Register(a.grpcServer)

	pb.RegisterUserV1Server(a.grpcServer, a.serviceProvider.authImpl)

	return a, nil
}
