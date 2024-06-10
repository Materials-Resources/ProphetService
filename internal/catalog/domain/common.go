package domain

type DeleteMode int32

const (
	DeleteModeUnspecified DeleteMode = 0
	DeleteModeSoft        DeleteMode = 1
	DeleteModeHard        DeleteMode = 2
)
