package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type IJurisdictionService interface {
    CreateJurisdiction(core.Jurisdiction) core.Jurisdiction
    UpdateJurisdiction(core.Jurisdiction) core.Jurisdiction
    GetJurisdiction(uint) (core.Jurisdiction, error)
    EnsureJurisdictions([]core.Jurisdiction)
}

type JurisdictionService struct {
    db *gorm.DB
}
func NewJurisdictionService(db *gorm.DB) *JurisdictionService {
    return &JurisdictionService{db: db}
}
func (p *JurisdictionService) CreateJurisdiction(jurisdiction core.Jurisdiction) core.Jurisdiction {
    p.db.Create(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) UpdateJurisdiction(jurisdiction core.Jurisdiction) core.Jurisdiction {
    p.db.Save(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) GetJurisdiction(code uint) (core.Jurisdiction, error) {
    var jurisdiction core.Jurisdiction
    err := p.db.Where(&core.Jurisdiction{Code: code}).First(&jurisdiction).Error
    return jurisdiction, err
}
func (p *JurisdictionService) EnsureJurisdictions(jurisdictions []core.Jurisdiction) {
    p.db.AutoMigrate(&core.Jurisdiction{})
    p.db.Model(&core.Jurisdiction{}).AddUniqueIndex("idx_jurisdiction_code", "code")
    p.db.Model(&core.Jurisdiction{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")

    for _, jurisdiction := range jurisdictions {
        existing, err := p.GetJurisdiction(jurisdiction.Code)

        if err != nil {
            p.CreateJurisdiction(jurisdiction)
        } else {
            existing.Name = jurisdiction.Name
            p.UpdateJurisdiction(existing)
        }
    }
}

