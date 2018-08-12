/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package services

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
    "skaioskit/models"
)

type IElectionService interface {
    CreateElection(models.Election) models.Election
    UpdateElection(models.Election) models.Election
    GetElection(uint64) (models.Election, error)
    GetElections(core.QueryRequest) ([]models.Election, uint64, error)
    EnsureElectionTable()
    EnsureElection(models.Election)
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
func (p *ElectionService) GetElections(query core.QueryRequest) ([]models.Election, uint64, error) {
    var count uint64
    var elections []models.Election
    election := models.Election{}

    core.BuildQueryWithoutPagination(p.db, query, &models.Election{}).Count(&count)
    err := core.BuildQuery(p.db, query, &election).Find(&elections).Error
    return elections, count, err
}
func (p *ElectionService) EnsureElectionTable() {
    p.db.AutoMigrate(&models.Election{})
    p.db.Model(&models.Election{}).AddUniqueIndex("idx_election_code", "code")
}
func (p *ElectionService) EnsureElection(election models.Election) {
    existing, err := p.GetElection(election.Code)

    if err != nil {
        p.CreateElection(election)
    } else {
        existing.Name = election.Name
        p.UpdateElection(existing)
    }
}

