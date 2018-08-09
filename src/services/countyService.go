package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/models"
)

type ICountyService interface {
    CreateCounty(models.County) models.County
    UpdateCounty(models.County) models.County
    GetCounty(uint) (models.County, error)
    EnsureCounties([]models.County)
}

type CountyService struct {
    db *gorm.DB
}
func NewCountyService(db *gorm.DB) *CountyService {
    return &CountyService{db: db}
}
func (p *CountyService) CreateCounty(county models.County) models.County {
    p.db.Create(&county)
    return county
}
func (p *CountyService) UpdateCounty(county models.County) models.County {
    p.db.Save(&county)
    return county
}
func (p *CountyService) GetCounty(code uint) (models.County, error) {
    var county models.County
    err := p.db.Where(&models.County{Code: code}).First(&county).Error
    return county, err
}
func (p *CountyService) EnsureCounties(counties []models.County) {
    p.db.AutoMigrate(&models.County{})
    p.db.Model(&models.County{}).AddUniqueIndex("idx_county_code", "code")

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

