package models

type Address struct {
	BaseModel
	Cep          *string `gorm:"column:tx_cep"`
	Street       *string `gorm:"column:tx_street"`
	Neighborhood *string `gorm:"column:tx_neighborhood"`
	City         *string `gorm:"column:tx_city"`
	State        *string `gorm:"column:tx_state"`
}
