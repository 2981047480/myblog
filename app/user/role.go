package user

type Role int

const (
	visitor Role = iota
	member
	admin
)
