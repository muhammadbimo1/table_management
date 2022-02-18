package manager

import (
	"log"
	"table_management/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Panic(err)
	}
	return &infra{
		db: resource,
	}
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	return db, nil
}
