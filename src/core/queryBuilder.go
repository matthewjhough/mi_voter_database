/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package core

import (
    "reflect"

    "github.com/jinzhu/gorm"
)

type QueryFilter struct {
    Field string
    Value string
}

type QueryRequest struct {
    Limit uint
    Offset uint
    Filters []QueryFilter
}

func buildQueryGetLimit(query QueryRequest) uint {
    //TODO: Make the max query limit configurable somewhere.
    maxLimit := uint(1000)
    if query.Limit > maxLimit {
        return maxLimit
    } else if query.Limit == 0 {
        return maxLimit
    } else {
        return query.Limit
    }
}

func buildQueryGetOffset(query QueryRequest) uint {
    return query.Offset
}

func BuildQuery(db *gorm.DB, query QueryRequest, obj interface{}) *gorm.DB {
    for _, filter := range query.Filters {
        field := reflect.ValueOf(obj).Elem().FieldByName(filter.Field)
        if field.IsValid() {
            field.SetString(filter.Value)
        } else {
            panic("Unknown Field: " + filter.Field)
        }
    }

    return db.Limit(buildQueryGetLimit(query)).Offset(buildQueryGetOffset(query)).Where(obj)
}
