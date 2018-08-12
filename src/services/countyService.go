/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/models"
)

type ICountyService interface {
    CreateCounty(models.County) models.County
    UpdateCounty(models.County) models.County
    GetCounties(skaioskit.QueryRequest) ([]models.County, uint64, error)
    GetCounty(uint) (models.County, error)
    EnsureCountyTable()
    EnsureCounty(models.County)
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
func (p *CountyService) GetCounties(query skaioskit.QueryRequest) ([]models.County, uint64, error) {
    var count uint64
    var counties []models.County

    skaioskit.BuildQueryWithoutPagination(p.db, query, &models.County{}).Count(&count)
    err := skaioskit.BuildQuery(p.db, query, &models.County{}).Find(&counties).Error
    return counties, count, err
}
func (p *CountyService) EnsureCountyTable() {
    p.db.AutoMigrate(&models.County{})
    p.db.Model(&models.County{}).AddUniqueIndex("idx_county_code", "code")
}
func (p *CountyService) EnsureCounty(county models.County) {
    existing, err := p.GetCounty(county.Code)
    if err != nil {
        p.CreateCounty(county)
    } else {
        existing.Name = county.Name
        p.UpdateCounty(existing)
    }
}
