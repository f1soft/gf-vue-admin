package response

import (
	"gf-server/app/model/admins"
)

type LoginResponse struct {
	User      *admins.Entity `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
