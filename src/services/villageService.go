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

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/models"
)

type IVillageService interface {
    CreateVillage(models.Village) models.Village
    UpdateVillage(models.Village) models.Village
    GetVillage(uint) (models.Village, error)
    GetVillages(skaioskit.QueryRequest) ([]models.Village, uint64, error)
    EnsureVillageTable()
    EnsureVillage(models.Village)
}

type VillageService struct {
    db *gorm.DB
}
func NewVillageService(db *gorm.DB) *VillageService {
    return &VillageService{db: db}
}
func (p *VillageService) CreateVillage(village models.Village) models.Village {
    p.db.Create(&village)
    return village
}
func (p *VillageService) UpdateVillage(village models.Village) models.Village {
    p.db.Save(&village)
    return village
}
func (p *VillageService) GetVillage(code uint) (models.Village, error) {
    var village models.Village
    err := p.db.Where(&models.Village{Code: code}).First(&village).Error
    return village, err
}
func (p *VillageService) GetVillages(query skaioskit.QueryRequest) ([]models.Village, uint64, error) {
    var count uint64
    var villages []models.Village
    village := models.Village{}

    skaioskit.BuildQueryWithoutPagination(p.db, query, &models.Village{}).Count(&count)
    err := skaioskit.BuildQuery(p.db, query, &village).Find(&villages).Error
    return villages, count, err
}
func (p *VillageService) EnsureVillageTable() {
    p.db.AutoMigrate(&models.Village{})
    p.db.Model(&models.Village{}).AddUniqueIndex("idx_village_code", "code")
    p.db.Model(&models.Village{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Village{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
}
func (p *VillageService) EnsureVillage(village models.Village) {
    existing, err := p.GetVillage(village.Code)

    if err != nil {
        p.CreateVillage(village)
    } else {
        existing.Name = village.Name
        p.UpdateVillage(existing)
    }
}
