package service

import (
	"gf-server/app/model/jwts"
	"gf-server/library/global"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"
)

func JsonInBlacklist(jwtList *jwts.Entity) (err error) {
	_, err = jwts.Insert(&jwtList)
	return err
}

// IsBlacklist check if the Jwt is in the blacklist or not
// IsBlacklist 判断JWT是否在黑名单内部
func IsBlacklist(jwt string) bool {
	return !jwts.RecordNotFound("jwt", jwt)
}

// GetRedisJWT Get user info in redis
// GetRedisJWT 获取用户在Redis的token
func GetRedisJWT(userUUID string) (string, error) {
	conn := global.GFVA_REDIS.Conn()
	defer conn.Close()
	r, err := conn.Do("GET", gconv.String(g.Cfg().Get("redis.Login_Prefix"))+userUUID)
	return gconv.String(r), err
}

// SetRedisJWT set jwt into the Redis
// SetRedisJWT 保存jwt到Redis
func SetRedisJWT(userUUID string, jwt string) (err error) {
	_, err = global.GFVA_REDIS.Do("SETEX", gconv.String(g.Cfg().Get("redis.Login_Prefix"))+userUUID, time.Hour*24, jwt)
	return err
}

// GetToken 根据uuid获取
// 抑制了错误, 但是不建议使用
func GetToken(userUUID string) string {
	conn := global.GFVA_REDIS.Conn()
	defer conn.Close()
	r, _ := conn.Do("GET", gconv.String(g.Cfg().Get("redis.Login_Prefix"))+userUUID)
	return gconv.String(r)
}

// ValidatorRedisToken 鉴权jwt
func ValidatorRedisToken(userUUID string, oldToken string) bool {
	redisPrefix := gconv.String(g.Cfg().Get("redis.Login_Prefix"))
	key := redisPrefix + userUUID
	if GetToken(key) != oldToken {
		return false
	}
	return true
}
