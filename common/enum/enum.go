package enum

// UserRole ...
type UserRole int

const (
	Unknow UserRole = iota + 1
	SuperAdmin
	Admin
	User
)
