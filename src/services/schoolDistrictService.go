package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/models"
)

type ISchoolDistrictService interface {
    CreateSchoolDistrict(models.SchoolDistrict) models.SchoolDistrict
    UpdateSchoolDistrict(models.SchoolDistrict) models.SchoolDistrict
    GetSchoolDistrict(uint) (models.SchoolDistrict, error)
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

