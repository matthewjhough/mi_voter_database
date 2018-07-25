package core
import (
    "github.com/jinzhu/gorm"
)


//School District
type SchoolDistrict struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint64
}

//County
type County struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
}

//Jurisdiction
type Jurisdiction struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
}

//Election
type Election struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint64
}

//Voter
type Voter struct {
    //gorm.Model
    ID uint `gorm:"primary_key"`

    LastName string  `gorm:"size:35"`
    FirstName string  `gorm:"size:20"`
    MiddleName string  `gorm:"size:20"`
    NameSuffix string  `gorm:"size:3"`
    Gender string       `gorm:"size:3"`

    VoterId uint64
}

//VoterHistory
type VoterHistory struct {
    ID uint `gorm:"primary_key"`
    VoterId uint64
    ElectionCode uint64
}
