/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package providers

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"

    "github.com/metakeule/fmtdate"

    "skaioskit/models"
)

type MichiganByteWidthDataProvider struct {
}
func NewMichiganByteWidthDataProvider() *MichiganByteWidthDataProvider {
    return &MichiganByteWidthDataProvider{}
}
func (p *MichiganByteWidthDataProvider) ParseCounties() <-chan models.County {
    chnl := make(chan models.County)
    go func() {
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

            chnl <- models.County{Code: uint(code), Name: strings.Trim(line[2:], " ")}
        }

        if err := scanner.Err(); err != nil {
            panic(err)
        }

        close(chnl)
    }()
    return chnl
}
func (p *MichiganByteWidthDataProvider) ParseJurisdictions() <-chan models.Jurisdiction {
    chnl := make(chan models.Jurisdiction)
    go func() {
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

            chnl <- models.Jurisdiction{Code: uint(code), CountyCode: uint(countyCode), Name: strings.Trim(line[7:], " ")}
        }

        if err := scanner.Err(); err != nil {
            panic(err)
        }

        close(chnl)
    }()
    return chnl
}
func (p *MichiganByteWidthDataProvider) ParseSchools() <-chan models.SchoolDistrict {
    chnl := make(chan models.SchoolDistrict)
    go func() {
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

            chnl <- models.SchoolDistrict{CountyCode: uint(countyCode), JurisdictionCode: uint(jurisdictionCode), Code: uint(code), Name: strings.Trim(line[12:], " ")}
        }

        if err := scanner.Err(); err != nil {
            panic(err)
        }
        close(chnl)
    }()
    return chnl
}
func (p *MichiganByteWidthDataProvider) ParseVillages() <-chan models.Village {
    chnl := make(chan models.Village)
    go func() {
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

            chnl <- models.Village{Code: uint(code), CountyCode: uint(countyCode), JurisdictionCode: uint(jurisdictionCode), VillageId: id, Name: strings.Trim(line[25:], " ")}
        }

        if err := scanner.Err(); err != nil {
            panic(err)
        }

        close(chnl)
    }()
    return chnl
}
func (p *MichiganByteWidthDataProvider) ParseElections() <-chan models.Election {
    chnl := make(chan models.Election)

    go func() {
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

            chnl <- models.Election{Code: code, Date: date, Name: strings.Trim(line[21:], " ")}
        }

        if err := scanner.Err(); err != nil {
            panic(err)
        }

        close(chnl)
    }()
    return chnl
}
func (p *MichiganByteWidthDataProvider) ParseVoters() <-chan models.Voter {
    chnl := make(chan models.Voter)
    go func() {
        file, err := os.Open("/data/entire_state_v.lst")
        if err != nil {
            panic(err)
        }
        defer file.Close()

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

            chnl <- models.Voter{
                VoterId: voterId,
                LastName: strings.Trim(line[0:35], " "),
                FirstName: strings.Trim(line[35:55], " "),
                MiddleName: strings.Trim(line[55:75], " "),
                NameSuffix: strings.Trim(line[75:78], " "),
                BirthYear: strings.Trim(line[78:82], " "),
                Gender: strings.Trim(line[82:83], " "),
                DateOfRegistration: dateOfReg,
                HouseNumberCharacter: strings.Trim(line[91:92], " "),
                ResidenceStreetNumber: strings.Trim(line[92:99], " "),
                HouseSuffix: strings.Trim(line[99:103], " "),
                AddressPreDirection: strings.Trim(line[103:105], " "),
                StreetName: strings.Trim(line[105:135], " "),
                StreetType: strings.Trim(line[135:141], " "),
                SuffixDirection: strings.Trim(line[141:143], " "),
                ResidenceRxtension: strings.Trim(line[143:156], " "),
                City: strings.Trim(line[156:191], " "),
                State: strings.Trim(line[191:193], " "),
                Zip: strings.Trim(line[193:198], " "),
                MailAddress1: strings.Trim(line[198:248], " "),
                MailAddress2: strings.Trim(line[248:298], " "),
                MailAddress3: strings.Trim(line[298:348], " "),
                MailAddress4: strings.Trim(line[348:398], " "),
                MailAddress5: strings.Trim(line[398:448], " "),
                CountyCode: uint(countyCode),
                JurisdictionCode: uint(jurisdictionCode),
                Ward: strings.Trim(line[468:474], " "),
                SchoolCode: uint(schoolCode),
                StateHouse: uint(stateHouse),
                StateSenate: uint(stateSenate),
                UsCongress: uint(usCongress),
                VillageCode: villageCode,
                CountyCommissioner: uint(countyCommissioner),
                VillagePrecinct: strings.Trim(line[504:510], " "),
                SchoolPrecinct: strings.Trim(line[510:516], " "),
                PermanentAbsenteeInd: strings.Trim(line[516:517], " "),
                StatusType: strings.Trim(line[517:519], " "),
                UOCAVAStatus: strings.Trim(line[519:520], " "),
            }
        }
        if err := scanner.Err(); err != nil {
            panic(err)
        }

        close(chnl)
    }()
    return chnl
}
func (p *MichiganByteWidthDataProvider) ParseVoterHistories() <-chan models.VoterHistory {
    chnl := make(chan models.VoterHistory)
    go func() {
        file, err := os.Open("/data/entire_state_h.lst")
        if err != nil {
            panic(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()
            voterId, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[0:13], " "), "0"), 0, 64)
            if err != nil {
                fmt.Println("voter code: " + err.Error())
                //TODO Log
                continue
            }
            electionCode, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[25:38], " "), "0"), 0, 64)
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

            chnl <- models.VoterHistory{
                VoterId: voterId,
                ElectionCode: electionCode,
                CountyCode: uint(countyCode),
                JurisdictionCode: uint(jurisdictionCode),
                SchoolCode: uint(schoolCode),
                AbsenteeInd: line[38:39],
            }
        }
        if err := scanner.Err(); err != nil {
            panic(err)
        }
        close(chnl)
    }()
    return chnl
}
