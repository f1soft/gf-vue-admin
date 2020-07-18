package response

import "gf-server/app/model/authorities"

type AuthorityCopyResponse struct {
	Authority    authorities.Authorities `json:"authority"`
	OldAuthorityId string             `json:"oldAuthorityId"`
}
