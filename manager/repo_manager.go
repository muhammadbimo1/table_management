package manager

import "table_management/repository"

type RepoManager interface {
	TableTransactionRepo() repository.TableTransactionRepository
	TableRepo() repository.TableRepository
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) TableTransactionRepo() repository.TableTransactionRepository {
	return repository.NewTableTransactionRepository(r.infra.SqlDb())
}

func (r *repoManager) TableRepo() repository.TableRepository {
	return repository.NewTableRepository(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
