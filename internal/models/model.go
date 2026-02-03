package models

import "time"

type BaseModel struct {
	CreatedDate  time.Time `gorm:"column:created_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	ModifiedDate time.Time `gorm:"column:modified_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified_date"`
}

type Country struct {
	CountryID   string `gorm:"primaryKey;column:country_id;type:char(2)" json:"country_id"`
	CountryName string `gorm:"column:country_name;type:varchar(40)" json:"country_name"`
	RegionID    uint   `gorm:"column:region_id" json:"region_id"`
	Region      Region `gorm:"foreignKey:RegionID;references:RegionID" json:"region"`
	BaseModel          //embedded field
}

func (Country) TableName() string { return "hr.countries" } //using schema hr

type  Region struct {
	RegionID   uint      `gorm:"primaryKey;column:region_id" json:"region_id"`
	RegionName string    `gorm:"column:region_name;type:varchar(25)" json:"region_name"`
	Countries  []Country `gorm:"foreignKey:RegionID;references:RegionID" json:"countries"`
	BaseModel
}

func (Region) TableName() string { return "hr.regions" }
