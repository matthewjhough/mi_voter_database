package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type IVoterService interface {
    CreateVoter(core.Voter) core.Voter
    UpdateVoter(core.Voter) core.Voter
    GetVoter(uint64) (core.Voter, error)
    EnsureVoterTable()
    EnsureVoter(core.Voter)
}

type VoterService struct {
    db *gorm.DB
}
func NewVoterService(db *gorm.DB) *VoterService {
    return &VoterService{db: db}
}
func (p *VoterService) CreateVoter(voter core.Voter) core.Voter {
    p.db.Create(&voter)
    return voter
}
func (p *VoterService) UpdateVoter(voter core.Voter) core.Voter {
    p.db.Save(&voter)
    return voter
}
func (p *VoterService) GetVoter(voterId uint64) (core.Voter, error) {
    var voter core.Voter
    err := p.db.Where(&core.Voter{VoterId: voterId}).First(&voter).Error
    return voter, err
}
func (p *VoterService) EnsureVoterTable() {
    p.db.AutoMigrate(&core.Voter{})
}
func (p *VoterService) EnsureVoter(voter core.Voter) {
    existing, err := p.GetVoter(voter.VoterId)

    if err != nil {
        p.CreateVoter(voter)
    } else {
        voter.ID = existing.ID
        p.UpdateVoter(voter)
    }
}

