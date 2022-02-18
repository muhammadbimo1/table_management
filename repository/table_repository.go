package repository

import (
	"errors"
	"table_management/entity"

	"gorm.io/gorm"
)

type TableRepository interface {
	CountByTableId(tableId string) (*entity.CustomerTable, error)
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{db: db}
}

func (t *tableRepository) CountByTableId(tableId string) (*entity.CustomerTable, error) {
	var tableResult entity.CustomerTable
	result := t.db.First(&tableResult, "id=?", tableId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("not found")
		}
		return nil, result.Error
	}
	return &tableResult, nil
}
