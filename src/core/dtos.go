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
}

type QueryFilter struct {
    Field string
    Value string
    Operator string
}

type QueryRequest struct {
    Limit uint
    Offset uint
    Filters []QueryFilter
}
