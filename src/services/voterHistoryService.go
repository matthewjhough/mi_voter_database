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

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/models"
)

type IVoterHistoryService interface {
    CreateVoterHistory(models.VoterHistory) models.VoterHistory
    UpdateVoterHistory(models.VoterHistory) models.VoterHistory
    GetVoterHistory(uint64) (models.VoterHistory, error)
    GetVoterHistories(skaioskit.QueryRequest) ([]models.VoterHistory, uint64, error)
    GetVoterHistoryCount() uint64
    EnsureVoterHistoryTable()
    EnsureVoterHistory(models.VoterHistory)
}

type VoterHistoryService struct {
    db *gorm.DB
}
func NewVoterHistoryService(db *gorm.DB) *VoterHistoryService {
    return &VoterHistoryService{db: db}
}
func (p *VoterHistoryService) CreateVoterHistory(history models.VoterHistory) models.VoterHistory {
    p.db.Create(&history)
    return history
}
func (p *VoterHistoryService) UpdateVoterHistory(history models.VoterHistory) models.VoterHistory {
    p.db.Save(&history)
    return history
}
func (p *VoterHistoryService) GetVoterHistory(electionCode uint64) (models.VoterHistory, error) {
    var history models.VoterHistory
    err := p.db.Where(&models.VoterHistory{ElectionCode: electionCode}).First(&history).Error
    return history, err
}
func (p *VoterHistoryService) GetVoterHistories(query skaioskit.QueryRequest) ([]models.VoterHistory, uint64, error) {
    var count uint64
    var voterHistories []models.VoterHistory
    voterHistory := models.VoterHistory{}

    skaioskit.BuildQueryWithoutPagination(p.db, query, &models.VoterHistory{}).Count(&count)
    err := skaioskit.BuildQuery(p.db, query, &voterHistory).Find(&voterHistories).Error
    return voterHistories, count, err
}
func (p *VoterHistoryService) GetVoterHistoryCount() uint64 {
    var count uint64
    p.db.Model(&models.VoterHistory{}).Count(&count)
    return count
}
func (p *VoterHistoryService) EnsureVoterHistoryTable() {
    p.db.AutoMigrate(&models.VoterHistory{})
    p.db.Model(&models.VoterHistory{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.VoterHistory{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.VoterHistory{}).AddForeignKey("school_code", "school_districts(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.VoterHistory{}).AddForeignKey("election_code", "elections(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.VoterHistory{}).AddForeignKey("voter_id", "voters(voter_id)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.VoterHistory{}).AddUniqueIndex("idx_voter_history_voter_id_election_code", "voter_id", "election_code")
}
func (p *VoterHistoryService) EnsureVoterHistory(history models.VoterHistory) {
    existing, err := p.GetVoterHistory(history.ElectionCode)

    if err != nil {
        p.CreateVoterHistory(history)
    } else {
        history.ID = existing.ID
        p.UpdateVoterHistory(history)
    }
}

