package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type IElectionService interface {
    CreateElection(Election) Election
    UpdateElection(Election) Election
    GetElection(uint64) (Election, error)
    EnsureElections([]Election)
}

type ElectionService struct {
    db *gorm.DB
}
func NewElectionService(db *gorm.DB) *ElectionService {
    return &ElectionService{db: db}
}
func (p *ElectionService) CreateElection(election Election) Election {
    p.db.Create(&election)
    return election
}
func (p *ElectionService) UpdateElection(election Election) Election {
    p.db.Save(&election)
    return election
}
func (p *ElectionService) GetElection(code uint64) (Election, error) {
    var election Election
    err := p.db.Where(&Election{Code: code}).First(&election).Error
    return election, err
}
func (p *ElectionService) EnsureElections(elections []Election) {
    p.db.AutoMigrate(&Election{})

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

