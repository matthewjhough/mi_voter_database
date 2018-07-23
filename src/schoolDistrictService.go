package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type ISchoolDistrictService interface {
    CreateSchoolDistrict(SchoolDistrict) SchoolDistrict
    UpdateSchoolDistrict(SchoolDistrict) SchoolDistrict
    GetSchoolDistrict(uint64) (SchoolDistrict, error)
    EnsureSchoolDistricts([]SchoolDistrict)
}

type SchoolDistrictService struct {
    db *gorm.DB
}
func NewSchoolDistrictService(db *gorm.DB) *SchoolDistrictService {
    return &SchoolDistrictService{db: db}
}
func (p *SchoolDistrictService) CreateSchoolDistrict(school SchoolDistrict) SchoolDistrict {
    p.db.Create(&school)
    return school
}
func (p *SchoolDistrictService) UpdateSchoolDistrict(school SchoolDistrict) SchoolDistrict {
    p.db.Save(&school)
    return school
}
func (p *SchoolDistrictService) GetSchoolDistrict(code uint64) (SchoolDistrict, error) {
    var school SchoolDistrict
    err := p.db.Where(&SchoolDistrict{Code: code}).First(&school).Error
    return school, err
}
func (p *SchoolDistrictService) EnsureSchoolDistricts(schools []SchoolDistrict) {
    p.db.AutoMigrate(&SchoolDistrict{})

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

