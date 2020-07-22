package model

import "github.com/gogf/gf/os/gtime"

type BaseModel struct {
	ID       int         `json:"id" orm:"id"`
	CreateAt *gtime.Time `json:"create_at" orm:"create_at"`
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at"`
	DeleteAt *gtime.Time `json:"delete_at" orm:"delete_at"`
}
