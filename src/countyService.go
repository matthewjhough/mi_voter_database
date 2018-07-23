package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type ICountyService interface {
    CreateCounty(County) County
    UpdateCounty(County) County
    GetCounty(uint) (County, error)
    EnsureCounties([]County)
}

type CountyService struct {
    db *gorm.DB
}
func NewCountyService(db *gorm.DB) *CountyService {
    return &CountyService{db: db}
}
func (p *CountyService) CreateCounty(county County) County {
    p.db.Create(&county)
    return county
}
func (p *CountyService) UpdateCounty(county County) County {
    p.db.Save(&county)
    return county
}
func (p *CountyService) GetCounty(code uint) (County, error) {
    var county County
    err := p.db.Where(&County{Code: code}).First(&county).Error
    return county, err
}
func (p *CountyService) EnsureCounties(counties []County) {
    p.db.AutoMigrate(&County{})

    for _, county := range counties {
        existing, err := p.GetCounty(county.Code)

        if err != nil {
            p.CreateCounty(county)
        } else {
            existing.Name = county.Name
            p.UpdateCounty(existing)
        }
    }
}

