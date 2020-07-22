package response

import (
	"server/app/model/admins"
	"server/app/model/authorities"

	"github.com/gogf/gf/os/gtime"
)

// Admin response Structure
type Admin struct {
	Id          uint                    `orm:"id,primary"   json:"id"`           // 自增ID
	CreateAt    *gtime.Time             `orm:"create_at"    json:"create_at"`    // 创建时间
	UpdateAt    *gtime.Time             `orm:"update_at"    json:"update_at"`    // 更新时间
	DeleteAt    *gtime.Time             `orm:"delete_at"    json:"delete_at"`    // 删除时间
	Uuid        string                  `orm:"uuid"         json:"uuid"`         // 用户唯一标识UUID
	Nickname    string                  `orm:"nickname"     json:"nickname"`     // 用户昵称
	HeaderImg   string                  `orm:"header_img"   json:"header_img"`   // 用户头像
	AuthorityId string                  `orm:"authority_id" json:"authority_id"` // 用户角色ID
	Username    string                  `orm:"username"     json:"username"`     // 用户名
	Password    string                  `orm:"password"     json:"-"`            // 用户登录密码
	Authority   authorities.Authorities `json:"authority"`
}

// AdminLogin response Structure
type AdminLogin struct {
	User      *Admin `json:"user"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

// AdminResponse response Structure
type AdminResponse struct {
	Admin *admins.Entity `json:"user"`
}
