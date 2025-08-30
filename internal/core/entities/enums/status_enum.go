package enums

type StatusEnum int

const (
	Pending StatusEnum = iota
	Running
	Late
	Done
	Cancelled
)
