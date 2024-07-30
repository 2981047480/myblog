package user

type Role int

const (
	Visitor Role = iota
	Member
	Admin
)
