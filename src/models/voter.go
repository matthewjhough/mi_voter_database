package models
import (
    "time"
)

type Voter struct {
    ID uint `gorm:"primary_key"`
    LastName string  `gorm:"size:35"`
    FirstName string  `gorm:"size:20"`
    MiddleName string  `gorm:"size:20"`
    NameSuffix string  `gorm:"size:3"`
    BirthYear string  `gorm:"size:4"`
    Gender string       `gorm:"size:3"`
    DateOfRegistration time.Time `gorm:"type:datetime"`
    HouseNumberCharacter string       `gorm:"size:1"`
    ResidenceStreetNumber string       `gorm:"size:7"`
    HouseSuffix string       `gorm:"size:4"`
    AddressPreDirection string       `gorm:"size:2"`
    StreetName string       `gorm:"size:30"`
    StreetType string       `gorm:"size:6"`
    SuffixDirection string       `gorm:"size:2"`
    ResidenceRxtension string       `gorm:"size:13"`
    City string       `gorm:"size:35"`
    State string       `gorm:"size:2"`
    Zip string       `gorm:"size:5"`
    PhoneNumber string       `gorm:"size:50"`
    MailAddress1 string       `gorm:"size:50"`
    MailAddress2 string       `gorm:"size:50"`
    MailAddress3 string       `gorm:"size:50"`
    MailAddress4 string       `gorm:"size:50"`
    MailAddress5 string       `gorm:"size:50"`
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
