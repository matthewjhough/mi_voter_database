package core
import (
    "time"

    "github.com/jinzhu/gorm"
)


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
    CountyCode uint
}

//School District
type SchoolDistrict struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
    JurisdictionCode uint
    CountyCode uint
}

//Village
type Village struct {
    gorm.Model

    VillageId uint64
    Name string  `gorm:"size:255"`
    Code uint
    JurisdictionCode uint
    CountyCode uint
}

//Election
type Election struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint64
    Date time.Time
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
