package usecase

import (
	"errors"
	"table_management/constant"
	"table_management/dto"
	"table_management/entity"
	"table_management/repository"

	"gorm.io/gorm"
)

type CustomerTableUseCase interface {
	GetTodayListCustomerTable() ([]dto.TableAvailability, error)
	TableCheckIn(request dto.CheckInRequest) (*entity.CustomerTableTransaction, error)
	TableCheckOut(billNo string) error
}

type customerTableUsecase struct {
	tableTransactionRepo repository.TableTransactionRepository
	tableRepo            repository.TableRepository
}

func NewCustomerTableUseCase(repo repository.TableTransactionRepository, tablerepo repository.TableRepository) CustomerTableUseCase {
	return &customerTableUsecase{
		tableTransactionRepo: repo,
		tableRepo:            tablerepo,
	}
}

func (c *customerTableUsecase) GetTodayListCustomerTable() ([]dto.TableAvailability, error) {
	return c.tableTransactionRepo.GetByBusinessDate()
}

func (c *customerTableUsecase) TableCheckIn(request dto.CheckInRequest) (*entity.CustomerTableTransaction, error) {
	_, err := c.tableRepo.CountByTableId(request.TableId)
	if err != nil {
		return nil, errors.New("table not found")
	}

	_, err = c.tableTransactionRepo.GetByBillNo(request.BillNo)
	if err != nil {
		return nil, errors.New("bill no duplicate")
	}

	tbl, err := c.tableTransactionRepo.CountByTableIdAndStatus(request.TableId, constant.TableOccupied)
	if err != nil {
		return nil, errors.New("check in error")
	}
	if tbl == 0 {
		return c.tableTransactionRepo.CreateOne(entity.CustomerTableTransaction{
			BillNo:          request.BillNo,
			CustomerTableID: request.TableId,
			Model:           gorm.Model{},
		})
	} else {
		return nil, errors.New("table occupied")
	}
}

func (c *customerTableUsecase) TableCheckOut(billNo string) error {
	return c.tableTransactionRepo.Delete(billNo)
}
