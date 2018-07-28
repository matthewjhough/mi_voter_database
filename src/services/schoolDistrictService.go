package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type ISchoolDistrictService interface {
    CreateSchoolDistrict(core.SchoolDistrict) core.SchoolDistrict
    UpdateSchoolDistrict(core.SchoolDistrict) core.SchoolDistrict
    GetSchoolDistrict(uint) (core.SchoolDistrict, error)
    EnsureSchoolDistricts([]core.SchoolDistrict)
}

type SchoolDistrictService struct {
    db *gorm.DB
}
func NewSchoolDistrictService(db *gorm.DB) *SchoolDistrictService {
    return &SchoolDistrictService{db: db}
}
func (p *SchoolDistrictService) CreateSchoolDistrict(school core.SchoolDistrict) core.SchoolDistrict {
    p.db.Create(&school)
    return school
}
func (p *SchoolDistrictService) UpdateSchoolDistrict(school core.SchoolDistrict) core.SchoolDistrict {
    p.db.Save(&school)
    return school
}
func (p *SchoolDistrictService) GetSchoolDistrict(code uint) (core.SchoolDistrict, error) {
    var school core.SchoolDistrict
    err := p.db.Where(&core.SchoolDistrict{Code: code}).First(&school).Error
    return school, err
}
func (p *SchoolDistrictService) EnsureSchoolDistricts(schools []core.SchoolDistrict) {
    p.db.AutoMigrate(&core.SchoolDistrict{})

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

