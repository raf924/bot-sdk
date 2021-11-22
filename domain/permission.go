package domain

type Permission uint

const IsUnknown Permission = 0
const (
	IsVerified    Permission = 1
	NeedVerified  Permission = 1
	IsModerator   Permission = 3
	NeedModerator Permission = 2
	IsAdmin       Permission = 7
	NeedAdmin     Permission = 4
)

func (p Permission) Has(permission Permission) bool {
	return permission == 0 || p&permission != 0
}
