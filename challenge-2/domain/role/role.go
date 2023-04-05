package role

import "time"

// Role is a struct that contains the role information
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
