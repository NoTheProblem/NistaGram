package model

type Role int

const(
	Regular Role = iota
	Administrator
	Agent
)
