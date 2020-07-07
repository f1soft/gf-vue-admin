package service

import (
	"errors"
	"gf-server/app/model/user"
	"gf-server/global"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// Login 用户登录
func Login(u *user.Entity) (userReturn *user.Entity, err error){
	var r gdb.Record
	r, err = global.GFVA_DB.Table("sys_user").Where(g.Map{"username = ?": u.Username, "password = ?": u.Username}).One()
	userInter := (*user.Entity)(nil) // 这种方式只有在查询到数据的时候才会执行初始化及内存分配,这里传递的参数类型其实是**User
	err = gconv.Struct(r, &userInter)
	if userReturn.CompareHashAndPassword(u.Password) { // 检查密码是否正确
		return userReturn, nil
	}
	return u, errors.New("未知用户")
}

// Register 注册
func Register(u *user.Entity) (userReturn *user.Entity, err error) {
	var (
		r gdb.Record
		userInter user.Entity
	)
	// 判断用户名是否注册
	r, err = global.GFVA_DB.Table("sys_user").Where(&u).One()
	err = gconv.Struct(r, &userInter)
	if userInter.Id != 0 {
			return &userInter, errors.New("用户名已注册")
	}
	// notRegister为false表明读取到了 不能注册
	//if !notRegister {
	//	return errors.New("用户名已注册"), userInter
	//} else {
	//	// 否则 附加uuid 密码md5简单加密 注册
	//	u.Password = utils.MD5V([]byte(u.Password))
	//	u.UUID = uuid.NewV4()
	//	err = global.GVA_DB.Create(&u).Error
	//}
	//return err, u
}