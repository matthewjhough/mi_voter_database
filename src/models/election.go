package models
import (
    "time"

    "github.com/jinzhu/gorm"
)

type Election struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint64
    Date time.Time `gorm:"type:datetime"`
}
