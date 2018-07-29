package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type ICountyService interface {
    CreateCounty(core.County) core.County
    UpdateCounty(core.County) core.County
    GetCounty(uint) (core.County, error)
    EnsureCounties([]core.County)
}

type CountyService struct {
    db *gorm.DB
}
func NewCountyService(db *gorm.DB) *CountyService {
    return &CountyService{db: db}
}
func (p *CountyService) CreateCounty(county core.County) core.County {
    p.db.Create(&county)
    return county
}
func (p *CountyService) UpdateCounty(county core.County) core.County {
    p.db.Save(&county)
    return county
}
func (p *CountyService) GetCounty(code uint) (core.County, error) {
    var county core.County
    err := p.db.Where(&core.County{Code: code}).First(&county).Error
    return county, err
}
func (p *CountyService) EnsureCounties(counties []core.County) {
    p.db.AutoMigrate(&core.County{})
    p.db.Model(&core.County{}).AddUniqueIndex("idx_county_code", "code")

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

