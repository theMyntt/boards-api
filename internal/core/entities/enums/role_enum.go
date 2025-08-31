package enums

import "fmt"

type RoleEnum int

const (
	User RoleEnum = iota
	Admin
)

var roleStrings = [...]string{
	"User",
	"Admin",
}

func (s StatusEnum) String() string {
	if int(s) < len(roleStrings) {
		return roleStrings[s]
	}
	return fmt.Sprintf("Status(%d)", s)
}
