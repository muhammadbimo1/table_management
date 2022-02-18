package manager

import "table_management/usecase"

type UseCaseManager interface {
	TableTransactionUseCase() usecase.CustomerTableUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) TableTransactionUseCase() usecase.CustomerTableUseCase {
	return usecase.NewCustomerTableUseCase(uc.repo.TableTransactionRepo(), uc.repo.TableRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repoManager,
	}
}
