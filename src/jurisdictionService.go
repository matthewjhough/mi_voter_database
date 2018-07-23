package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type IJurisdictionService interface {
    CreateJurisdiction(Jurisdiction) Jurisdiction
    UpdateJurisdiction(Jurisdiction) Jurisdiction
    GetJurisdiction(uint) (Jurisdiction, error)
    EnsureJurisdictions([]Jurisdiction)
}

type JurisdictionService struct {
    db *gorm.DB
}
func NewJurisdictionService(db *gorm.DB) *JurisdictionService {
    return &JurisdictionService{db: db}
}
func (p *JurisdictionService) CreateJurisdiction(jurisdiction Jurisdiction) Jurisdiction {
    p.db.Create(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) UpdateJurisdiction(jurisdiction Jurisdiction) Jurisdiction {
    p.db.Save(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) GetJurisdiction(code uint) (Jurisdiction, error) {
    var jurisdiction Jurisdiction
    err := p.db.Where(&Jurisdiction{Code: code}).First(&jurisdiction).Error
    return jurisdiction, err
}
func (p *JurisdictionService) EnsureJurisdictions(jurisdictions []Jurisdiction) {
    p.db.AutoMigrate(&Jurisdiction{})

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

