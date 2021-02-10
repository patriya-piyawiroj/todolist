package models

var (
	StatusTypes = map[string]bool{
		Todo:       true,
		InProgress: true,
		Done:       true,
	}
)

const (
	Todo       = "Todo"
	InProgress = "In Progress"
	Done       = "Done"
)
