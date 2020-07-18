package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model"
	"gf-server/library/global"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gdbadapter "github.com/flipped-aurora/gdb-adapter"
	"github.com/gogf/gf/frame/g"
	"strings"
)

// UpdateCasbin update casbin authority
// UpdateCasbin 更新casbin权限
func UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	ClearCasbin(0, authorityId)
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			ID:          0,
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		addFlag := AddCasbin(cm)
		if addFlag == false {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

// AddCasbin add casbin authority
// AddCasbin 添加权限
func AddCasbin(cm model.CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy(cm.AuthorityId, cm.Path, cm.Method)
}

// UpdateCasbinApi update casbin apis
// UpdateCasbinApi API更新随动
func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	_, err := global.GFVA_DB.Table("casbin_rule").Data(g.Map{"v1": newPath, "v2": newMethod}).Where(g.Map{"v1": oldPath, "v2": oldMethod}).Update()
	return err
}

// @title    GetPolicyPathByAuthorityId
// @description   get policy path by authorityId, 获取权限列表
func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin Clear matching permissions
// ClearCasbin 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)

}

// Casbin store to DB,
// Casbin 持久化到数据库  引入自定义规则
func Casbin() *casbin.Enforcer {
	a, err := gdbadapter.NewAdapterByDB(global.GFVA_DB)
	if err != nil {
		panic(err)
	}
	e := casbin.NewEnforcer(g.Cfg().GetString("casbin.ModelPath"), a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// ParamsMatch customized rule
// ParamsMatch 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0] // 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc customized function
// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return ParamsMatch(name1, name2), nil
}
