package models

type Policy struct {
	BaseModel
	Method  []string `gorm:"type:text[];column:ls_method"`
	Path    string   `gorm:"column:tx_path"`
	GroupID int      `gorm:"column:id_group"`
	Group   Group    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
