package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type IElectionService interface {
    CreateElection(core.Election) core.Election
    UpdateElection(core.Election) core.Election
    GetElection(uint64) (core.Election, error)
    EnsureElections([]core.Election)
}

type ElectionService struct {
    db *gorm.DB
}
func NewElectionService(db *gorm.DB) *ElectionService {
    return &ElectionService{db: db}
}
func (p *ElectionService) CreateElection(election core.Election) core.Election {
    p.db.Create(&election)
    return election
}
func (p *ElectionService) UpdateElection(election core.Election) core.Election {
    p.db.Save(&election)
    return election
}
func (p *ElectionService) GetElection(code uint64) (core.Election, error) {
    var election core.Election
    err := p.db.Where(&core.Election{Code: code}).First(&election).Error
    return election, err
}
func (p *ElectionService) EnsureElections(elections []core.Election) {
    p.db.AutoMigrate(&core.Election{})
    p.db.Model(&core.Election{}).AddUniqueIndex("idx_election_code", "code")

    for _, election := range elections {
        existing, err := p.GetElection(election.Code)

        if err != nil {
            p.CreateElection(election)
        } else {
            existing.Name = election.Name
            p.UpdateElection(existing)
        }
    }
}

