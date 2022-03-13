// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNamePersonField = "chii_person_fields"

// PersonField mapped from table <chii_person_fields>
type PersonField struct {
	OwnerType string `gorm:"column:prsn_cat;type:enum('prsn','crt');primaryKey" json:"prsn_cat"`
	OwnerID   uint32 `gorm:"column:prsn_id;type:int(8) unsigned;primaryKey;index:prsn_id,priority:1" json:"prsn_id"`
	Gender    uint8  `gorm:"column:gender;type:tinyint(4) unsigned;not null" json:"gender"`
	Bloodtype uint8  `gorm:"column:bloodtype;type:tinyint(4) unsigned;not null" json:"bloodtype"`
	BirthYear uint16 `gorm:"column:birth_year;type:year(4);not null" json:"birth_year"`
	BirthMon  uint8  `gorm:"column:birth_mon;type:tinyint(2) unsigned;not null" json:"birth_mon"`
	BirthDay  uint8  `gorm:"column:birth_day;type:tinyint(2) unsigned;not null" json:"birth_day"`
}

// TableName PersonField's table name
func (*PersonField) TableName() string {
	return TableNamePersonField
}
