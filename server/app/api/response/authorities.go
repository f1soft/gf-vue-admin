package response

import "server/app/model/authorities"

type AuthorityCopy struct {
	Authority      authorities.Authorities `json:"authority"`
	OldAuthorityId string                  `r:"oldAuthorityId" json:"oldAuthorityId"`
}

type Authority struct {
	Authority *authorities.Entity `json:"authority"`
}
type CopyAuthority struct {
	Authority *authorities.Authorities `json:"authority"`
}
