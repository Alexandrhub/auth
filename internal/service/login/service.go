package login

import (
	"github.com/alexandrhub/auth/internal/client/db"
	"github.com/alexandrhub/auth/internal/repository"
	"github.com/alexandrhub/auth/internal/service"
)

type serverAuth struct {
	loginRepository repository.LoginRepository
	txManager       db.TxManager
}

func NewService(loginRepository repository.LoginRepository, txManager db.TxManager) service.LoginService {
	return &serverAuth{loginRepository: loginRepository, txManager: txManager}
}
