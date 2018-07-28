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
    Date time.Time `gorm:"type:datetime"`
}

//Voter
type Voter struct {
    //gorm.Model
    ID uint `gorm:"primary_key"`

    LastName string  `gorm:"size:35"`
    FirstName string  `gorm:"size:20"`
    MiddleName string  `gorm:"size:20"`
    NameSuffix string  `gorm:"size:3"`
    BirthYear string  `gorm:"size:4"`
    Gender string       `gorm:"size:3"`
    DateOfRegistration time.Time `gorm:"type:datetime"`
    Address string       `gorm:"size:511"`
    VoterId uint64
    CountyCode uint
    JurisdictionCode uint
    Ward string `gorm:"size:6"`
    SchoolCode uint
    StateHouse uint
    StateSenate uint
    UsCongress uint
    CountyCommissioner uint
    VillageCode uint
    VillagePrecinct string `gorm:"size:6"`
    SchoolPrecinct string `gorm:"size:6"`
    PermanentAbsenteeInd string  `gorm:"size:1"`
    StatusType string        `gorm:"size:2"`
    UOCAVAStatus string        `gorm:"size:1"`
}

//VoterHistory
type VoterHistory struct {
    ID uint `gorm:"primary_key"`
    VoterId uint64
    CountyCode uint
    JurisdictionCode uint
    SchoolCode uint
    ElectionCode uint64
    AbsenteeInd string  `gorm:"size:1"`
}
