package core

type GetVotersResponse struct {
    Voters []Voter
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
