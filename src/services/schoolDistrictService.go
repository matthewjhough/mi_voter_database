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

    "skaioskit/core"
    "skaioskit/models"
)

type ISchoolDistrictService interface {
    CreateSchoolDistrict(models.SchoolDistrict) models.SchoolDistrict
    UpdateSchoolDistrict(models.SchoolDistrict) models.SchoolDistrict
    GetSchoolDistrict(uint) (models.SchoolDistrict, error)
    GetSchoolDistricts(core.QueryRequest) ([]models.SchoolDistrict, error)
    EnsureSchoolDistricts([]models.SchoolDistrict)
}

type SchoolDistrictService struct {
    db *gorm.DB
}
func NewSchoolDistrictService(db *gorm.DB) *SchoolDistrictService {
    return &SchoolDistrictService{db: db}
}
func (p *SchoolDistrictService) CreateSchoolDistrict(school models.SchoolDistrict) models.SchoolDistrict {
    p.db.Create(&school)
    return school
}
func (p *SchoolDistrictService) UpdateSchoolDistrict(school models.SchoolDistrict) models.SchoolDistrict {
    p.db.Save(&school)
    return school
}
func (p *SchoolDistrictService) GetSchoolDistrict(code uint) (models.SchoolDistrict, error) {
    var school models.SchoolDistrict
    err := p.db.Where(&models.SchoolDistrict{Code: code}).First(&school).Error
    return school, err
}
func (p *SchoolDistrictService) GetSchoolDistricts(query core.QueryRequest) ([]models.SchoolDistrict, error) {
    var schoolDistricts []models.SchoolDistrict
    schoolDistrict := models.SchoolDistrict{}

    err := core.BuildQuery(p.db, query, &schoolDistrict).Find(&schoolDistricts).Error
    return schoolDistricts, err
}
func (p *SchoolDistrictService) EnsureSchoolDistricts(schools []models.SchoolDistrict) {
    p.db.AutoMigrate(&models.SchoolDistrict{})
    p.db.Model(&models.SchoolDistrict{}).AddUniqueIndex("idx_school_district_code", "code")
    p.db.Model(&models.SchoolDistrict{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.SchoolDistrict{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")

    for _, school := range schools {
        existing, err := p.GetSchoolDistrict(school.Code)

        if err != nil {
            p.CreateSchoolDistrict(school)
        } else {
            existing.Name = school.Name
            p.UpdateSchoolDistrict(existing)
        }
    }
}

