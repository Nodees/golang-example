package models

type Group struct {
	BaseModel
	Name string `gorm:"unique;column:tx_name"`
}

type UserGroup struct {
	BaseModel
	GroupID int   `gorm:"column:id_group"`
	UserID  int   `gorm:"column:id_user"`
	Group   Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User    User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
