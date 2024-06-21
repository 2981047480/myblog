package common

type UPDATE_MODE int

const (
	UPDATE_MODE_POST UPDATE_MODE = iota
	UPDATE_MODE_PATCH
)
