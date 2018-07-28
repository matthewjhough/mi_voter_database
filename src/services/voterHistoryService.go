package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type IVoterHistoryService interface {
    CreateVoterHistory(core.VoterHistory) core.VoterHistory
    UpdateVoterHistory(core.VoterHistory) core.VoterHistory
    GetVoterHistory(uint64) (core.VoterHistory, error)
    EnsureVoterHistoryTable()
    EnsureVoterHistory(core.VoterHistory)
}

type VoterHistoryService struct {
    db *gorm.DB
}
func NewVoterHistoryService(db *gorm.DB) *VoterHistoryService {
    return &VoterHistoryService{db: db}
}
func (p *VoterHistoryService) CreateVoterHistory(history core.VoterHistory) core.VoterHistory {
    p.db.Create(&history)
    return history
}
func (p *VoterHistoryService) UpdateVoterHistory(history core.VoterHistory) core.VoterHistory {
    p.db.Save(&history)
    return history
}
func (p *VoterHistoryService) GetVoterHistory(electionCode uint64) (core.VoterHistory, error) {
    var history core.VoterHistory
    err := p.db.Where(&core.VoterHistory{ElectionCode: electionCode}).First(&history).Error
    return history, err
}
func (p *VoterHistoryService) EnsureVoterHistoryTable() {
    p.db.AutoMigrate(&core.VoterHistory{})
}
func (p *VoterHistoryService) EnsureVoterHistory(history core.VoterHistory) {
    existing, err := p.GetVoterHistory(history.ElectionCode)

    if err != nil {
        p.CreateVoterHistory(history)
    } else {
        history.ID = existing.ID
        p.UpdateVoterHistory(history)
    }
}

