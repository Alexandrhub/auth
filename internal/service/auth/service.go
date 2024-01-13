package auth

import (
	"github.com/alexandrhub/auth/internal/db"
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
