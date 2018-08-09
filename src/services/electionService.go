package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/models"
)

type IElectionService interface {
    CreateElection(models.Election) models.Election
    UpdateElection(models.Election) models.Election
    GetElection(uint64) (models.Election, error)
    EnsureElections([]models.Election)
}

type ElectionService struct {
    db *gorm.DB
}
func NewElectionService(db *gorm.DB) *ElectionService {
    return &ElectionService{db: db}
}
func (p *ElectionService) CreateElection(election models.Election) models.Election {
    p.db.Create(&election)
    return election
}
func (p *ElectionService) UpdateElection(election models.Election) models.Election {
    p.db.Save(&election)
    return election
}
func (p *ElectionService) GetElection(code uint64) (models.Election, error) {
    var election models.Election
    err := p.db.Where(&models.Election{Code: code}).First(&election).Error
    return election, err
}
func (p *ElectionService) EnsureElections(elections []models.Election) {
    p.db.AutoMigrate(&models.Election{})
    p.db.Model(&models.Election{}).AddUniqueIndex("idx_election_code", "code")

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

