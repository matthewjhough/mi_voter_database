package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type IVillageService interface {
    CreateVillage(core.Village) core.Village
    UpdateVillage(core.Village) core.Village
    GetVillage(uint) (core.Village, error)
    EnsureVillages([]core.Village)
}

type VillageService struct {
    db *gorm.DB
}
func NewVillageService(db *gorm.DB) *VillageService {
    return &VillageService{db: db}
}
func (p *VillageService) CreateVillage(village core.Village) core.Village {
    p.db.Create(&village)
    return village
}
func (p *VillageService) UpdateVillage(village core.Village) core.Village {
    p.db.Save(&village)
    return village
}
func (p *VillageService) GetVillage(code uint) (core.Village, error) {
    var village core.Village
    err := p.db.Where(&core.Village{Code: code}).First(&village).Error
    return village, err
}
func (p *VillageService) EnsureVillages(villages []core.Village) {
    p.db.AutoMigrate(&core.Village{})

    for _, village := range villages {
        existing, err := p.GetVillage(village.Code)

        if err != nil {
            p.CreateVillage(village)
        } else {
            existing.Name = village.Name
            p.UpdateVillage(existing)
        }
    }
}
