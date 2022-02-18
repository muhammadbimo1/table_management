package repository

import (
	"errors"
	"log"
	"table_management/constant"
	"table_management/dto"
	"table_management/entity"
	"table_management/util"

	"gorm.io/gorm"
)

type TableTransactionRepository interface {
	CreateOne(trx entity.CustomerTableTransaction) (*entity.CustomerTableTransaction, error)
	GetByBusinessDate() ([]dto.TableAvailability, error)
	CountByTableIdAndStatus(tableId string, status constant.TableStatus) (int64, error)
	Delete(billNo string) error
	GetByBillNo(billno string) (*entity.CustomerTableTransaction, error)
}

type tableTransactionRepository struct {
	db *gorm.DB
}

func NewTableTransactionRepository(db *gorm.DB) TableTransactionRepository {
	return &tableTransactionRepository{db: db}
}

func (t *tableTransactionRepository) CreateOne(trx entity.CustomerTableTransaction) (*entity.CustomerTableTransaction, error) {
	err := t.db.Create(&trx).Error
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &trx, nil
}

func (t *tableTransactionRepository) Delete(billNo string) error {
	err := t.db.Where("bill_no=?", billNo).Delete(&entity.CustomerTableTransaction{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (t *tableTransactionRepository) CountByTableIdAndStatus(tableId string, status constant.TableStatus) (int64, error) {
	var count int64
	var err error
	switch status {
	case constant.TableAllStatus:
		err = t.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id =?", tableId).Count(&count).Error
	case constant.TableVacant:
		err = t.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id =? and deleted_at is not null", tableId).Count(&count).Error
	case constant.TableOccupied:
		err = t.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id =? and deleted_at is null", tableId).Count(&count).Error
	default:
		return -1, errors.New("unknown table status")
	}
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return count, nil
}

func (t *tableTransactionRepository) GetByBusinessDate() ([]dto.TableAvailability, error) {
	var tableListResult []dto.TableAvailability
	sd, ed := util.GetTodayWithTime()
	err := t.db.Model(&entity.CustomerTableTransaction{}).
		Select("customer_table.id as table_id,count(customer_table_transaction.created_at) as is_occupied").Group("customer_table.id").Joins("right join customer_table on customer_table_transaction.customer_table_id = customer_table.id and customer_table_transaction.created_at between ? and ? and customer_table_transaction.deleted_at is null", sd, ed).Scan(&tableListResult).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return tableListResult, nil
}

func (t *tableTransactionRepository) GetByBillNo(billno string) (*entity.CustomerTableTransaction, error) {
	var tableResult entity.CustomerTableTransaction
	_ = t.db.First(&tableResult, "bill_no=?", billno)
	log.Println(tableResult)
	if tableResult != (entity.CustomerTableTransaction{}) {
		return nil, errors.New("duplicate id")
	}
	return &tableResult, nil
}
