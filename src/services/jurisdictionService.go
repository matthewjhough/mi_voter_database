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

type IJurisdictionService interface {
    CreateJurisdiction(models.Jurisdiction) models.Jurisdiction
    UpdateJurisdiction(models.Jurisdiction) models.Jurisdiction
    GetJurisdiction(uint) (models.Jurisdiction, error)
    GetJurisdictions(skaioskit.QueryRequest) ([]models.Jurisdiction, uint64, error)
    EnsureJurisdictionTable()
    EnsureJurisdiction(models.Jurisdiction)
}

type JurisdictionService struct {
    db *gorm.DB
}
func NewJurisdictionService(db *gorm.DB) *JurisdictionService {
    return &JurisdictionService{db: db}
}
func (p *JurisdictionService) CreateJurisdiction(jurisdiction models.Jurisdiction) models.Jurisdiction {
    p.db.Create(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) UpdateJurisdiction(jurisdiction models.Jurisdiction) models.Jurisdiction {
    p.db.Save(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) GetJurisdiction(code uint) (models.Jurisdiction, error) {
    var jurisdiction models.Jurisdiction
    err := p.db.Where(&models.Jurisdiction{Code: code}).First(&jurisdiction).Error
    return jurisdiction, err
}
func (p *JurisdictionService) GetJurisdictions(query skaioskit.QueryRequest) ([]models.Jurisdiction, uint64, error) {
    var count uint64
    var jurisdictions []models.Jurisdiction
    jurisdiction := models.Jurisdiction{}

    skaioskit.BuildQueryWithoutPagination(p.db, query, &models.Jurisdiction{}).Count(&count)
    err := skaioskit.BuildQuery(p.db, query, &jurisdiction).Find(&jurisdictions).Error
    return jurisdictions, count, err
}
func (p *JurisdictionService) EnsureJurisdictionTable() {
    p.db.AutoMigrate(&models.Jurisdiction{})
    p.db.Model(&models.Jurisdiction{}).AddUniqueIndex("idx_jurisdiction_code", "code")
    p.db.Model(&models.Jurisdiction{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
}
func (p *JurisdictionService) EnsureJurisdiction(jurisdiction models.Jurisdiction) {
    existing, err := p.GetJurisdiction(jurisdiction.Code)

    if err != nil {
        p.CreateJurisdiction(jurisdiction)
    } else {
        existing.Name = jurisdiction.Name
        p.UpdateJurisdiction(existing)
    }
}

