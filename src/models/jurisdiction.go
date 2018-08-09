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
    "github.com/jinzhu/gorm"
)

type Jurisdiction struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
    CountyCode uint
}
