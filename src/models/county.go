package models
import (
    "github.com/jinzhu/gorm"
)

type County struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
}
