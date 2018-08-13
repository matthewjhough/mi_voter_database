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
    "time"

    "github.com/jinzhu/gorm"
)

type Election struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint64
    Date time.Time `gorm:"type:datetime"`
}
func GetElectionCSVHeader() []string {
    var ret []string

    ret = append(ret, "name")
    ret = append(ret, "election_code")
    ret = append(ret, "date")

    return ret
}
func (e *Election) ToSlice() []string {
    var ret []string

    ret = append(ret, e.Name)
    ret = append(ret, strconv.FormatUint(uint64(e.Code), 10))
    ret = append(ret, e.Date.Format("2006-01-02"))

    return ret
}
