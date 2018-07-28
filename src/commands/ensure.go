package commands

import (
    "os"
    "bytes"
    "bufio"
    "strconv"
    "strings"

    "github.com/metakeule/fmtdate"
    "github.com/spf13/cobra"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
    "skaioskit/services"
)

func ensureSchools(service services.ISchoolDistrictService) {
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

func ensureVillages(service services.IVillageService) {
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

func ensureCounties(service services.ICountyService) {
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

func ensureJurisdictions(service services.IJurisdictionService) {
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

func ensureElections(service services.IElectionService) {
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

func ensureVoters(db *gorm.DB, service services.IVoterService) {
    service.EnsureVoterTable()

    file, err := os.Open("/data/entire_state_v.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var buffer bytes.Buffer
    buffer.WriteString("INSERT INTO voters(voter_id, last_name, first_name, middle_name, name_suffix, gender) VALUES ")
    vals := []interface{}{}

    counter := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        voterId, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[448:461], " "), "0"), 0, 64)

        if err != nil {
            //panic(err)
            //TODO Log
        } else {
            counter++
            buffer.WriteString("(?, ?, ?, ?, ?, ?), ")
            vals = append(
                vals,
                voterId, //VoterId: 
                strings.Trim(line[0:35], " "), //LastName: 
                strings.Trim(line[35:55], " "), //FirstName: 
                strings.Trim(line[55:75], " "), //MiddleName: 
                strings.Trim(line[75:78], " "), //NameSuffix: 
                strings.Trim(line[82:83], " "),  //Gender: 
            )

            if counter > 3000 {
                sqlStr := buffer.String()
                //trim the last ,
                sqlStr = sqlStr[0:len(sqlStr)-2]

                //prepare the statement
                stmt, err := db.DB().Prepare(sqlStr)
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
                buffer.WriteString("INSERT INTO voters(voter_id, last_name, first_name, middle_name, name_suffix, gender) VALUES ")
            }
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
        stmt, err := db.DB().Prepare(sqlStr)
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

func ensureVoterHistories(db *gorm.DB, service services.IVoterHistoryService) {
    service.EnsureVoterHistoryTable()

    file, err := os.Open("/data/entire_state_h.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var buffer bytes.Buffer
    buffer.WriteString("INSERT INTO voter_histories(voter_id, election_code) VALUES ")
    vals := []interface{}{}

    counter := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        voterId, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[0:13], " "), "0"), 0, 64)
        code, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[25:38], " "), "0"), 0, 64)

        if err != nil {
            //panic(err)
            //TODO Log
        } else {
            counter++
            buffer.WriteString("(?, ?), ")
            vals = append(
                vals,
                voterId,
                code,
            )

            if counter > 3000 {
                sqlStr := buffer.String()
                //trim the last ,
                sqlStr = sqlStr[0:len(sqlStr)-2]

                //prepare the statement
                stmt, err := db.DB().Prepare(sqlStr)
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
                buffer.WriteString("INSERT INTO voter_histories(voter_id, election_code) VALUES ")
            }
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
        stmt, err := db.DB().Prepare(sqlStr)
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

var ensureCmd = &cobra.Command{
    Use:   "ensure",
    Short: "imports the database",
    Long:  `ensures the database schema exists and has imported the voter data.`,
    Run: func(cmd *cobra.Command, args []string) {
        //setup db connection
        db, err := gorm.Open("mysql", os.Getenv("APP_MYSQL_CONN_STR"))
        if err != nil {
            panic(err)
        }
        defer db.Close()

        //setup services
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        villageService := services.NewVillageService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        //ensure db
        ensureCounties(countyService)
        ensureJurisdictions(jurisdictionService)
        ensureSchools(schoolService)
        ensureVillages(villageService)
        ensureElections(electionService)
        ensureVoters(db, voterService)
        ensureVoterHistories(db, voterHistoryService)
    },
}

//Entry
func init() {
    RootCmd.AddCommand(ensureCmd)
}
