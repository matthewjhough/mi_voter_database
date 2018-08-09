package services

import (
    "reflect"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
)

type IVoterService interface {
    CreateVoter(core.Voter) core.Voter
    UpdateVoter(core.Voter) core.Voter
    GetVoters(core.QueryRequest) ([]core.Voter, error)
    GetVoter(uint64) (core.Voter, error)
    GetVoterCount() uint64
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
func (p *VoterService) GetVoters(query core.QueryRequest) ([]core.Voter, error) {
    var voters []core.Voter
    voter := core.Voter{}

    for _, filter := range query.Filters {
        field := reflect.ValueOf(&voter).Elem().FieldByName(filter.Field)
        if field.IsValid() {
            field.SetString(filter.Value)
        } else {
            panic("Unknown Field: " + filter.Field)
        }
    }

    err := p.db.Limit(query.Limit).Offset(query.Offset).Where(&voter).Find(&voters).Error
    return voters, err
}
func (p *VoterService) GetVoter(voterId uint64) (core.Voter, error) {
    var voter core.Voter
    err := p.db.Where(&core.Voter{VoterId: voterId}).First(&voter).Error
    return voter, err
}
func (p *VoterService) GetVoterCount() uint64 {
    var count uint64
    p.db.Model(&core.Voter{}).Count(&count)
    return count
}
func (p *VoterService) EnsureVoterTable() {
    p.db.AutoMigrate(&core.Voter{})
    p.db.Model(&core.Voter{}).AddUniqueIndex("idx_voter_voter_id", "voter_id")
    p.db.Model(&core.Voter{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&core.Voter{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&core.Voter{}).AddForeignKey("school_code", "school_districts(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&core.Voter{}).AddForeignKey("village_code", "villages(code)", "RESTRICT", "RESTRICT")
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
