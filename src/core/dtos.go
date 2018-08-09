package core

type GetVotersResponse struct {
    Voters []Voter
}

type GetVoterCount struct {
    Count uint64
}

type GetVoterHistoryCount struct {
    Count uint64
}

type GetAboutResponse struct {
    Version string
    BuildTime string
}
