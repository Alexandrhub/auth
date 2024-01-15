package auth

import (
	"github.com/alexandrhub/auth/internal/client/db"
	"github.com/alexandrhub/auth/internal/repository"
	"github.com/alexandrhub/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(authRepository repository.AuthRepository, txManager db.TxManager) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}

// NewMockService для тестирования
func NewMockService(deps ...any) service.AuthService {
	srv := &serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.AuthRepository:
			srv.authRepository = s
		}
	}

	return srv
}

func NewMockTxService(deps ...any) *serv {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case serv:
			srv = s
		}
	}

	return &srv
}
