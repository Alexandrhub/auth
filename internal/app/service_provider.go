package app

import (
	"context"
	"log"

	"github.com/alexandrhub/auth/internal/auth"
	"github.com/alexandrhub/auth/internal/closer"
	"github.com/alexandrhub/auth/internal/config"
	"github.com/alexandrhub/auth/internal/db"
	"github.com/alexandrhub/auth/internal/db/pg"
	"github.com/alexandrhub/auth/internal/db/transaction"
	"github.com/alexandrhub/auth/internal/repository"
	authRepository "github.com/alexandrhub/auth/internal/repository/auth"
	"github.com/alexandrhub/auth/internal/service"
	authService "github.com/alexandrhub/auth/internal/service/auth"
)

type serviceProvider struct {
	Config config.Config

	dbClient  db.Client
	txManager db.TxManager
	// TODO: add other services

	authService service.AuthService

	authRepository repository.AuthRepository

	authImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx), s.TxManager(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.NewDBClient(ctx, s.Config.DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}
