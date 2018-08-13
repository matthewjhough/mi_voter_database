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

    "github.com/jinzhu/gorm"
)

type SchoolDistrict struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
    JurisdictionCode uint
    CountyCode uint
}
func GetSchoolDistrictCSVHeader() []string {
    var ret []string

    ret = append(ret, "name")
    ret = append(ret, "school_code")
    ret = append(ret, "jurisdiction_code")
    ret = append(ret, "county_code")

    return ret
}
func (s *SchoolDistrict) ToSlice() []string {
    var ret []string

    ret = append(ret, s.Name)
    ret = append(ret, strconv.FormatUint(uint64(s.Code), 10))
    ret = append(ret, strconv.FormatUint(uint64(s.JurisdictionCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(s.CountyCode), 10))

    return ret
}
