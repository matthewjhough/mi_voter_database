/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package models
import (
    "strconv"
)

type VoterHistory struct {
    ID uint `gorm:"primary_key"`
    VoterId uint64
    CountyCode uint
    JurisdictionCode uint
    SchoolCode uint
    ElectionCode uint64
    AbsenteeInd string  `gorm:"size:1"`
}
func GetVoterHistoryCSVHeader() []string {
    var ret []string

    ret = append(ret, "voter_id")
    ret = append(ret, "county_code")
    ret = append(ret, "jurisdiction_code")
    ret = append(ret, "school_code")
    ret = append(ret, "election_code")
    ret = append(ret, "absentee_ind")

    return ret
}
func (v *VoterHistory) ToSlice() []string {
    var ret []string

    ret = append(ret, strconv.FormatUint(v.VoterId, 10))
    ret = append(ret, strconv.FormatUint(uint64(v.CountyCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.JurisdictionCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.SchoolCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.ElectionCode), 10))
    ret = append(ret, v.AbsenteeInd)

    return ret
}
