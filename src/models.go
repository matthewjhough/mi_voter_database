package main
import (
    "github.com/jinzhu/gorm"
)


//Turf
type Turf struct {
    gorm.Model

    TownshipId   uint
    TurfNumber   string  `gorm:"size:32"`
    TurfStatusId uint
    FormId       *uint
}
