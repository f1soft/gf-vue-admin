package response

import "gf-server/app/model/user"

type LoginResponse struct {
	User      *user.Entity `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
