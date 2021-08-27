package model

type RelationType struct {
	Relation Relation `json:"relation"`
}

type Relation int

const(
	NotFollowing Relation = iota
	Following
	NotAccepted
	Blocked
	Blocking
)
