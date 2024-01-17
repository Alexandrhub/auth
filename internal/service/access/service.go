package access

import (
	"github.com/alexandrhub/auth/internal/client/db"
	"github.com/alexandrhub/auth/internal/repository"
	"github.com/alexandrhub/auth/internal/service"
)

type serverAccess struct {
	accessRepository repository.AccessRepository
	txManager        db.TxManager
}

func NewService(accessRepository repository.AccessRepository, txManager db.TxManager) service.AccessService {
	return &serverAccess{accessRepository: accessRepository, txManager: txManager}
}
