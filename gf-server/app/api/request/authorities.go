package request

type CreateAuthority struct {
	AuthorityId   string `json:"authority_id"`
	AuthorityName string `json:"authority_name"`
	ParentId      string `json:"parent_id"`
}

type UpdateAuthority struct {
	AuthorityId   string `json:"authority_id"`
	AuthorityName string `json:"authority_name"`
	ParentId      string `json:"parent_id"`
}
