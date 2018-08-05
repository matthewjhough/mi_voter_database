package providers

import (
    "fmt"
    "os"
    "bytes"
    "bufio"
    "strconv"
    "strings"

    "github.com/metakeule/fmtdate"
    "github.com/jinzhu/gorm"

    "skaioskit/core"
    "skaioskit/services"
)

type MichiganByteWidthDataProvider struct {
    db *gorm.DB
}
func NewMichiganByteWidthDataProvider(db *gorm.DB) *MichiganByteWidthDataProvider {
    return &MichiganByteWidthDataProvider{db: db}
}
func (p *MichiganByteWidthDataProvider) EnsureCounties(service services.ICountyService) {
    counties := []core.County{}

    file, err := os.Open("/data/countycd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        code, err := strconv.ParseUint(strings.TrimLeft(line[0:2], "0"), 0, 32)
        if err != nil {
            panic(err)
        }

        county := core.County{Code: uint(code), Name: line[2:]}
        counties = append(counties, county)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureCounties(counties)
}
func (p *MichiganByteWidthDataProvider) EnsureJurisdictions(service services.IJurisdictionService) {
    counties := []core.Jurisdiction{}

    file, err := os.Open("/data/jurisdcd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        countyCode, err := strconv.ParseUint(strings.TrimLeft(line[0:2], "0"), 0, 32)
        code, err := strconv.ParseUint(strings.TrimLeft(line[2:7], "0"), 0, 32)
        if err != nil {
            panic(err)
        }

        county := core.Jurisdiction{Code: uint(code), CountyCode: uint(countyCode), Name: line[7:]}
        counties = append(counties, county)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureJurisdictions(counties)
}
func (p *MichiganByteWidthDataProvider) EnsureSchools(service services.ISchoolDistrictService) {
    schools := []core.SchoolDistrict{}

    file, err := os.Open("/data/schoolcd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        countyCode, err := strconv.ParseUint(strings.TrimLeft(line[0:2], "0"), 0, 32)
        jurisdictionCode, err := strconv.ParseUint(strings.TrimLeft(line[2:7], "0"), 0, 32)
        code, err := strconv.ParseUint(strings.TrimLeft(line[7:12], "0"), 0, 32)
        if err != nil {
            panic(err)
        }

        school := core.SchoolDistrict{CountyCode: uint(countyCode), JurisdictionCode: uint(jurisdictionCode), Code: uint(code), Name: line[12:]}
        schools = append(schools, school)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureSchoolDistricts(schools)
}
func (p *MichiganByteWidthDataProvider) EnsureVillages(service services.IVillageService) {
    villages := []core.Village{}

    file, err := os.Open("/data/villagecd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        id, err := strconv.ParseUint(strings.TrimLeft(line[0:13], "0"), 0, 64)
        countyCode, err := strconv.ParseUint(strings.TrimLeft(line[13:15], "0"), 0, 64)
        jurisdictionCode, err := strconv.ParseUint(strings.TrimLeft(line[15:20], "0"), 0, 64)
        code, err := strconv.ParseUint(strings.Trim(strings.TrimLeft(line[20:25], "0"), " "), 0, 64)
        if err != nil {
            panic(err)
        }

        village := core.Village{Code: uint(code), CountyCode: uint(countyCode), JurisdictionCode: uint(jurisdictionCode), VillageId: id, Name: line[25:]}
        villages = append(villages, village)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureVillages(villages)
}
func (p *MichiganByteWidthDataProvider) EnsureElections(service services.IElectionService) {
    elections := []core.Election{}

    file, err := os.Open("/data/electionscd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        code, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[0:13], " "), "0"), 0, 64)
        date, err := fmtdate.Parse("MMDDYYYY", line[13:21])
        if err != nil {
            panic(err)
        }

        election := core.Election{Code: code, Date: date, Name: line[21:]}
        elections = append(elections, election)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureElections(elections)
}
func (p *MichiganByteWidthDataProvider) EnsureVoters(service services.IVoterService) {
    service.EnsureVoterTable()

    file, err := os.Open("/data/entire_state_v.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    onDuplicateClause := " ON DUPLICATE KEY UPDATE last_name=VALUES(last_name), first_name=VALUES(first_name)"
    insertClause := "INSERT INTO voters(voter_id, last_name, first_name, middle_name, name_suffix, birth_year, gender, date_of_registration, house_number_character, residence_street_number, house_suffix, address_pre_direction, street_name, street_type, suffix_direction, residence_rxtension, city, state, zip, mail_address1, mail_address2, mail_address3, mail_address4, mail_address5, county_code, jurisdiction_code, ward, school_code, state_house, state_senate, us_congress, county_commissioner, village_code, village_precinct, school_precinct, permanent_absentee_ind, status_type, uocava_status) VALUES "
    var buffer bytes.Buffer
    buffer.WriteString(insertClause)
    vals := []interface{}{}

    counter := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        voterId, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[448:461], " "), "0"), 0, 64)
        if err != nil {
            fmt.Println("voter id: " + err.Error())
            //TODO Log
            continue
        }

        dateOfReg, err := fmtdate.Parse("MMDDYYYY", line[83:91])
        if err != nil {
            fmt.Println("date of reg: " + err.Error())
            //TODO Log
            continue
        }
        countyCode, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[461:463], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("county code: " + err.Error())
            //TODO Log
            continue
        }
        jurisdictionCode, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[463:468], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("jurisdiction code: " + err.Error())
            //TODO Log
            continue
        }
        schoolCode, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[474:479], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("school code: " + err.Error())
            //TODO Log
            continue
        }
        stateHouse, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[479:484], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("stateHouse: " + err.Error())
            //TODO Log
            continue
        }
        stateSenate, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[484:489], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("stateSenate: " + err.Error())
            //TODO Log
            continue
        }
        usCongress, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[489:494], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("usCongress: " + err.Error())
            //TODO Log
            continue
        }
        countyCommissioner, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[494:499], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("countyCommissioner: " + err.Error())
            //TODO Log
            continue
        }
        var villageCode *uint
        villageCodeTmp, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[499:504], " "), "0"), 0, 32)
        if err != nil {
            villageCode = nil
        } else {
            villageCodeTmp32 := uint(villageCodeTmp)
            villageCode = &villageCodeTmp32
        }

        counter++
        buffer.WriteString("(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?), ")
        vals = append(
            vals,
            voterId, //VoterId: 
            strings.Trim(line[0:35], " "), //LastName: 
            strings.Trim(line[35:55], " "), //FirstName: 
            strings.Trim(line[55:75], " "), //MiddleName: 
            strings.Trim(line[75:78], " "), //NameSuffix: 
            strings.Trim(line[78:82], " "),  //BirthYear: 
            strings.Trim(line[82:83], " "),  //Gender:
            dateOfReg.Format("2006-01-02 15:04:05"),
            strings.Trim(line[91:92], " "),  //house_number_character
            strings.Trim(line[92:99], " "),  //residence_street_number
            strings.Trim(line[99:103], " "),  //house_suffix
            strings.Trim(line[103:105], " "),  //address_pre_direction
            strings.Trim(line[105:135], " "),  //street_name
            strings.Trim(line[135:141], " "),  //street_type
            strings.Trim(line[141:143], " "),  //suffix_direction
            strings.Trim(line[143:156], " "),  //residence_rxtension
            strings.Trim(line[156:191], " "),  //City:
            strings.Trim(line[191:193], " "),  //State:
            strings.Trim(line[193:198], " "),  //Zip:
            strings.Trim(line[198:248], " "),  //Address 1:
            strings.Trim(line[248:298], " "),  //Address 2:
            strings.Trim(line[298:348], " "),  //Address 3:
            strings.Trim(line[348:398], " "),  //Address 4:
            strings.Trim(line[398:448], " "),  //Address 5:
            uint(countyCode),
            uint(jurisdictionCode),
            strings.Trim(line[468:474], " "), //Ward:
            uint(schoolCode),
            uint(stateHouse),
            uint(stateSenate),
            uint(usCongress),
            uint(countyCommissioner),
            villageCode,
            strings.Trim(line[504:510], " "), //VillagePrecinct
            strings.Trim(line[510:516], " "), //SchoolPrecinct
            strings.Trim(line[516:517], " "), //PermanentAbsenteeInd
            strings.Trim(line[517:519], " "), //StatusType
            strings.Trim(line[519:520], " "), //UOCAVAStatus
        )

        if counter > 1500 {
            sqlStr := buffer.String()
            //trim the last ,
            sqlStr = sqlStr[0:len(sqlStr)-2]

            //prepare the statement
            stmt, err := p.db.DB().Prepare(sqlStr + onDuplicateClause)
            if err != nil {
                panic(err)
            }
            _, err = stmt.Exec(vals...)
            if err != nil {
                panic(err)
            }
            stmt.Close()

            counter = 0
            vals = []interface{}{}
            buffer.Reset()
            buffer.WriteString(insertClause)
        }
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    if counter > 0 {
        sqlStr := buffer.String()
        //trim the last ,
        sqlStr = sqlStr[0:len(sqlStr)-2]
        //prepare the statement
        stmt, err := p.db.DB().Prepare(sqlStr + onDuplicateClause)
        if err != nil {
            panic(err)
        }
        _, err = stmt.Exec(vals...)
        if err != nil {
            panic(err)
        }
        stmt.Close()
    }
}
func (p *MichiganByteWidthDataProvider) EnsureVoterHistories(service services.IVoterHistoryService) {
    service.EnsureVoterHistoryTable()

    file, err := os.Open("/data/entire_state_h.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    onDuplicateClause := " ON DUPLICATE KEY UPDATE county_code=VALUES(county_code), jurisdiction_code=VALUES(jurisdiction_code), school_code=VALUES(school_code), absentee_ind=VALUES(absentee_ind)"
    insertSqlClause := "INSERT INTO voter_histories(voter_id, election_code, county_code, jurisdiction_code, school_code, absentee_ind) VALUES "
    var buffer bytes.Buffer
    buffer.WriteString(insertSqlClause)
    vals := []interface{}{}

    counter := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        voterId, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[0:13], " "), "0"), 0, 64)
        if err != nil {
            fmt.Println("voter code: " + err.Error())
            //TODO Log
            continue
        }
        code, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[25:38], " "), "0"), 0, 64)
        if err != nil {
            fmt.Println("election code: " + err.Error())
            //TODO Log
            continue
        }
        countyCode, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[13:15], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("county code: " + err.Error())
            //TODO Log
            continue
        }
        jurisdictionCode, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[15:20], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("jurisdiction code: " + err.Error())
            //TODO Log
            continue
        }
        schoolCode, err := strconv.ParseUint(strings.TrimLeft(strings.Trim(line[20:25], " "), "0"), 0, 32)
        if err != nil {
            fmt.Println("school code: " + err.Error())
            //TODO Log
            continue
        }

        counter++
        buffer.WriteString("(?, ?, ?, ?, ?, ?), ")
        vals = append(
            vals,
            voterId,
            code,
            countyCode,
            jurisdictionCode,
            schoolCode,
            line[38:39],
        )

        if counter > 6000 {
            sqlStr := buffer.String()
            //trim the last ,
            sqlStr = sqlStr[0:len(sqlStr)-2]

            //prepare the statement
            stmt, err := p.db.DB().Prepare(sqlStr + onDuplicateClause)
            if err != nil {
                panic(err)
            }
            _, err = stmt.Exec(vals...)
            if err != nil {
                panic(err)
            }
            stmt.Close()

            counter = 0
            vals = []interface{}{}
            buffer.Reset()
            buffer.WriteString(insertSqlClause)
        }
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    if counter > 0 {
        sqlStr := buffer.String()
        //trim the last ,
        sqlStr = sqlStr[0:len(sqlStr)-2]
        //prepare the statement
        stmt, err := p.db.DB().Prepare(sqlStr + onDuplicateClause)
        if err != nil {
            panic(err)
        }
        _, err = stmt.Exec(vals...)
        if err != nil {
            panic(err)
        }
        stmt.Close()
    }
}
