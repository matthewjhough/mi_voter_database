package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type IVoterService interface {
    CreateVoter(Voter) Voter
    UpdateVoter(Voter) Voter
    GetVoter(uint64) (Voter, error)
    EnsureVoters([]Voter)
}

type VoterService struct {
    db *gorm.DB
}
func NewVoterService(db *gorm.DB) *VoterService {
    return &VoterService{db: db}
}
func (p *VoterService) CreateVoter(voter Voter) Voter {
    p.db.Create(&voter)
    return voter
}
func (p *VoterService) UpdateVoter(voter Voter) Voter {
    p.db.Save(&voter)
    return voter
}
func (p *VoterService) GetVoter(voterId uint64) (Voter, error) {
    var voter Voter
    err := p.db.Where(&Voter{VoterId: voterId}).First(&voter).Error
    return voter, err
}
func (p *VoterService) EnsureVoters(voters []Voter) {
    p.db.AutoMigrate(&Voter{})

    for _, voter := range voters {
        existing, err := p.GetVoter(voter.VoterId)

        if err != nil {
            p.CreateVoter(voter)
        } else {
            voter.ID = existing.ID
            p.UpdateVoter(voter)
        }
    }
}

