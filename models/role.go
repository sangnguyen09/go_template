package models

type Role int

const (
	ADMIN Role =987393 // cho một số bất kì
	MEMBER Role = 1
)

func (r Role) String() string  {
	switch r {
	case ADMIN:
		return "ADMIN"
	case MEMBER:
		return "MEMBER"
	default:
		return "Unknown"
	}
}