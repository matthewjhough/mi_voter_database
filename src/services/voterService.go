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

type IVoterService interface {
    CreateVoter(models.Voter) models.Voter
    UpdateVoter(models.Voter) models.Voter
    GetVoters(core.QueryRequest) ([]models.Voter, error)
    GetVoter(uint64) (models.Voter, error)
    GetVoterCount() uint64
    EnsureVoterTable()
    EnsureVoter(models.Voter)
}

type VoterService struct {
    db *gorm.DB
}
func NewVoterService(db *gorm.DB) *VoterService {
    return &VoterService{db: db}
}
func (p *VoterService) CreateVoter(voter models.Voter) models.Voter {
    p.db.Create(&voter)
    return voter
}
func (p *VoterService) UpdateVoter(voter models.Voter) models.Voter {
    p.db.Save(&voter)
    return voter
}
func (p *VoterService) GetVoters(query core.QueryRequest) ([]models.Voter, error) {
    var voters []models.Voter
    voter := models.Voter{}

    err := core.BuildQuery(p.db, query, &voter).Find(&voters).Error
    return voters, err
}
func (p *VoterService) GetVoter(voterId uint64) (models.Voter, error) {
    var voter models.Voter
    err := p.db.Where(&models.Voter{VoterId: voterId}).First(&voter).Error
    return voter, err
}
func (p *VoterService) GetVoterCount() uint64 {
    var count uint64
    p.db.Model(&models.Voter{}).Count(&count)
    return count
}
func (p *VoterService) EnsureVoterTable() {
    p.db.AutoMigrate(&models.Voter{})
    p.db.Model(&models.Voter{}).AddUniqueIndex("idx_voter_voter_id", "voter_id")
    p.db.Model(&models.Voter{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Voter{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Voter{}).AddForeignKey("school_code", "school_districts(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Voter{}).AddForeignKey("village_code", "villages(code)", "RESTRICT", "RESTRICT")
}
func (p *VoterService) EnsureVoter(voter models.Voter) {
    existing, err := p.GetVoter(voter.VoterId)

    if err != nil {
        p.CreateVoter(voter)
    } else {
        voter.ID = existing.ID
        p.UpdateVoter(voter)
    }
}
