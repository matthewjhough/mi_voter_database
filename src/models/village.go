package models
import (
    "github.com/jinzhu/gorm"
)

type Village struct {
    gorm.Model

    VillageId uint64
    Name string  `gorm:"size:255"`
    Code uint
    JurisdictionCode uint
    CountyCode uint
}
