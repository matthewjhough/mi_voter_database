package models
import (
    "github.com/jinzhu/gorm"
)

type Jurisdiction struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
    CountyCode uint
}
