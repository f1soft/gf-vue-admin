package response

import "server/app/model"

type AuthorityMenu struct {
	Menus []*model.AuthorityMenu `json:"menus"`
}

type BaseMenus struct {
	Menus []model.BaseMenu `json:"menus"`
}

type BaseMenu struct {
	Menu model.BaseMenu `json:"menu"`
}
