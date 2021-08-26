package model

type Auth struct {
	Username string `json:"username"`
	Role int `json:"role"`
}


type Role int

const(
	Regular Role = iota
	Administrator
	Agent
)
