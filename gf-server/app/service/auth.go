package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/api/response"
	"gf-server/app/model/admins"
	"gf-server/app/model/jwts"
	"gf-server/library/global"
	"time"

	"github.com/gomodule/redigo/redis"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var (
	// The underlying JWT middleware.
	// 底层的JWT中间件
	GfJWTMiddleware *jwt.GfJWTMiddleware
)

// Initialization function,
// 初始化函数

// rewrite this function to customized your own JWT settings.
// 重写此函数以自定义您自己的JWT设置。
func init() {
	signingKey := g.Cfg().GetString("jwt.SigningKey")
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           signingKey,
		Key:             []byte(signingKey),
		Timeout:         time.Hour * 24,     // 1 天
		MaxRefresh:      time.Hour * 24 * 7, // 刷新的token设置为一星期
		IdentityKey:     "id",
		TokenLookup:     "header: x-token, query: token, cookie: jwt",
		TokenHeadName:   "",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
	if err != nil {
		global.GFVA_LOG.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.

// PayloadFunc是将在登录期间调用的回调函数。
// 使用此功能可以向网络令牌添加其他有效载荷数据。
// 然后在请求期间通过c.Get("JWT_PAYLOAD")使数据可用。
// 请注意，有效负载未加密。
// jwt.io上提到的属性不能用作地图的键。
// 可选，默认情况下不会设置其他数据。
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler sets the identity for JWT.
// IdentityHandler 设置JWT的身份。
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// Unauthorized is used to define customized Unauthorized callback function.
// Unauthorized 用于定义自定义的未经授权的回调函数。
func Unauthorized(r *ghttp.Request, code int, message string) {
	global.FailWithDetailed(r, code, g.Map{"reload": true}, "未登录或非法访问"+message)
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
// LoginResponse 用于定义自定义的登录成功回调函数
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	userReturn := (*admins.Entity)(nil)
	if err := gconv.Struct(r.GetParam("user_info"), &userReturn); err != nil {
		global.FailWithMessage(r, "登录失败")
		r.Exit()
	}
	if !g.Cfg().GetBool("system.UseMultipoint") {
		global.OkDetailed(r, response.LoginResponse{User: userReturn, Token: token, ExpiresAt: expire.Unix() * 1000}, "登录成功!")
		r.Exit()
	}
	redisJwt, err := GetRedisJWT(userReturn.Uuid)
	if err == redis.ErrNil {
		if err := SetRedisJWT(userReturn.Uuid, token); err != nil {
			global.Result(r, code, g.Map{}, "设置登录状态失败")
			r.Exit()
		}
	} else if err != nil {
		global.Result(r, code, g.Map{}, err.Error())
	} else {
		if err = JsonInBlacklist(&jwts.Entity{Jwt: redisJwt}); err != nil {
			global.Result(r, code, g.Map{}, "jwt作废失败")
			r.Exit()
		}
		if err := SetRedisJWT(userReturn.Uuid, token); err != nil {
			global.Result(r, code, g.Map{}, "设置登录状态失败")
			r.Exit()
		}
	}
	global.OkDetailed(r, response.LoginResponse{User: userReturn, Token: token, ExpiresAt: expire.Unix() * 1000}, "登录成功!")
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
// RefreshResponse 用于获取新令牌，无论当前令牌是否过期。
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {

	// TODO Redis判断jwt是否有效
	//_ = r.Response.WriteJson(g.Map{
	//	"code":   http.StatusOK,
	//	"token":  token,
	//	"expire": expire.Format(time.RFC3339),
	//})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
// Authenticator 用于验证登录参数。
// 它必须返回用户数据作为用户标识符，并将其存储在Claim Array中。
// 检查错误（e），以确定适当的错误消息。
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var L request.LoginRequest
	if err := r.Parse(&L); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if !store.Verify(L.CaptchaId, L.Captcha, true) { // 验证码校对
		return nil, errors.New("验证码错误")
	}
	userReturn, err := admins.FindOne(g.Map{"username": L.Username})
	if err != nil {
		global.GFVA_LOG.Error(err)
		return nil, jwt.ErrFailedAuthentication
	}
	if userReturn.CompareHashAndPassword(L.Password) { // 检查密码是否正确
		r.SetParam("user_info", userReturn) // 设置参数保存到请求中
		return gconv.Map(&userReturn), nil
	}
	return nil, jwt.ErrFailedAuthentication

}
