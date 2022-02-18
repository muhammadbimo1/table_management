package entity

type CustomerTable struct {
	ID string `gorm:"size:3;primaryKey"`
}

func (m *CustomerTable) TableName() string {
	return "customer_table"
} //?
