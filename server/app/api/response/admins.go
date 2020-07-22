package response

import (
	"server/app/model/admins"
	"server/app/model/authorities"
)

// Admin response Structure
type Admin struct {
	admins.Entity
	Authority authorities.Authorities `json:"authority"`
}

// Login response Structure
type Login struct {
	User      *Admin `json:"user"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

// AdminResponse response Structure
type AdminResponse struct {
	Admin *admins.Entity `json:"user"`
}
