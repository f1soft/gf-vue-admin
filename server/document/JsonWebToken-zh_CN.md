##  Json Web Token

## 1.Simple configuration

| Name              | Type          | Example                                            | Description                                                  |
| :---------------- | ------------- | :------------------------------------------------- | :----------------------------------------------------------- |
| Realm             | string        | "test zhou"                                        | 显示给用户的领域名称,必须的                                  |
| SigningAlgorithm  | string        | "HS256"                                            | 签名算法-可能的值为HS256，HS384，HS512，可选，默认值为HS256。 |
| Key               | []byte        | []byte("secret key")                               | 用于签名的密钥,必须的.                                       |
| Timeout           | time.Duration | time.Minute * 5                                    | jwt令牌有效的持续时间。可选，默认为一小时。                  |
| MaxRefresh        | time.Duration | time.Minute * 5                                    | 该字段允许客户端刷新令牌，直到MaxRefresh通过。<br/>请注意，客户端可以在MaxRefresh的最后时刻刷新其令牌。<br/>这意味着令牌的最大有效时间跨度为TokenTime + MaxRefresh。<br/>可选，默认为0，表示不可刷新。 |
| IdentityKey       | string        | "id"                                               | 设置身份密钥                                                 |
| TokenLookup       | string        | "header: Authorization, query: token, cookie: jwt" | TokenLookup是"<source>:<name>"形式的字符串，用于从请求中提取令牌,可选的。<br/>默认值"header:Authorization". <br/>可能的值：<br />- "header:<name>"<br />- "query:<name>" <br />- "cookie:<name>" |
| TokenHeadName     | string        | "Bearer"                                           | TokenHeadName是标题中的字符串。默认值为"Bearer"              |
| PrivKeyFile       | string        |                                                    | 非对称算法的私钥文件                                         |
| PubKeyFile        | string        |                                                    | 非对称算法的公钥文件                                         |
| SendCookie        | bool          |                                                    | (可选)将令牌作为Cookie返回                                   |
| SecureCookie      | bool          |                                                    | 允许不安全的Cookie通过HTTP进行开发                           |
| CookieHTTPOnly    | bool          |                                                    | 允许访问客户端的Cookie进行开发                               |
| CookieDomain      | string        |                                                    | 允许更改Cookie域以进行开发                                   |
| SendAuthorization | bool          |                                                    | SendAuthorization允许为每个请求返回授权标头                  |
| DisabledAbort     | bool          |                                                    | 禁用上下文的abort()                                          |
| CookieName        | string        |                                                    | CookieName允许更改Cookie名称以进行开发                       |

|      配置名      | 类型                                          |                        示例                        |                           中文描述                           |
| :--------------: | --------------------------------------------- | :------------------------------------------------: | :----------------------------------------------------------: |
|      Realm       | string                                        |                    "test zhou"                     |            要显示给用户的领域名称,必须传递的参数             |
| SigningAlgorithm | string                                        |                                                    | 签名算法-可能的值为HS256、HS384、HS512，可选，默认值为HS256  |
|       Key        | []byte                                        |                []byte("secret key")                |                                                              |
|     Timeout      | time.Duration                                 |                  time.Minute * 5                   |                        token过期时间                         |
|    MaxRefresh    | time.Duration                                 |                  time.Minute * 5                   |                                                              |
|   IdentityKey    | string                                        |                        "id"                        |                       身份验证的key值                        |
|   TokenLookup    | string                                        | "header: Authorization, query: token, cookie: jwt" |         token检索模式，用于提取token-> Authorization         |
|  TokenHeadName   | string                                        |                      "Bearer"                      | token在请求头时的名称，默认值为Bearer.客户端在header中传入Authorization 对一个值是Bearer + 空格 + token |
|     TimeFunc     | func() time.Time                              |                      time.Now                      |              测试或服务器在其他时区可设置该属性              |
|  Authenticator   | func(r *ghttp.Request) (interface{}, error)   |                   Authenticator                    |           根据登录信息对用户进行身份验证的回调函数           |
|   Authorizator   | func(data interface{}, r *ghttp.Request) bool |                                                    | 回调函数，该函数应执行已验证用户的授权。仅在身份验证成功后调用。成功时必须返回true，失败时返回false。可选，默认为成功 |
|  LoginResponse   |                                               |                   LoginResponse                    |     完成登录后返回的信息，用户可自定义返回数据，默认返回     |
| RefreshResponse  |                                               |                auth.RefreshResponse                |    刷新token后返回的信息，用户可自定义返回数据，默认返回     |
|   Unauthorized   |                                               |                 auth.Unauthorized                  |                     处理不进行授权的逻辑                     |
| IdentityHandler  |                                               |                auth.IdentityHandler                |                    解析并设置用户身份信息                    |
|   PayloadFunc    |                                               |                  auth.PayloadFunc                  |                     登录期间的回调的函数                     |

## 2. 回调函数配置

### 2.1 Authenticator

- Type:==func(r *ghttp.Request) (interface{}, error)==
- Description: 
	- 根据登录信息对用户进行身份验证的回调函数。
	- 必须返回用户数据作为用户标识符，它将存储在索赔数组中。必须的。
	- 检查错误(e)以确定适当的错误信息.

### 2.2 Authorizator

- Type:==func(data interface{}, r *ghttp.Request) bool==
- Description: 
	- 回调函数，该函数应执行经过身份验证的用户的授权。
	- 仅在身份验证成功后调用. 成功时必须返回true，失败时必须返回false.
	- 可选，默认success.

### 2.3 PayloadFunc

- Type:==func(data interface{}) MapClaims==
- Description: 
	- 在登录期间将调用的回调函数。
	- 使用此功能可以向网络令牌添加其他有效负载数据。
	- 然后在请求期间通过c.Get(" JWT_PAYLOAD")使数据可用。
	- 请注意，有效负载未加密。
	- jwt.io上提到的属性不能用作地图的键。
	- 可选，默认情况下不会设置其他数据。

### 2.4 Unauthorized

- Type:==func(*ghttp.Request, int, string)==
- Description: 
	- 用户可以定义自己的未经授权的功能。

### 2.5 LoginResponse

- Type:==func(*ghttp.Request, int, string, time.Time)==
- Description: 
	- 用户可以定义自己的LoginResponse函数。

### 2.6 RefreshResponse

- Type:==func(*ghttp.Request, int, string, time.Time)==
- Description: 
	- 用户可以定义自己的 RefreshResponse 函数.

### 2.7 IdentityHandler

- Type:==func(*ghttp.Request) interface{}==
- Description: 
	- 设置身份处理程序功能

## 3.示例代码

```go
// 示例代码
package auth

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
	"net/http"
	"time"
)

var (
	// 底层JWT中间件.
	GfJWTMiddleware *jwt.GfJWTMiddleware
	// 自定义登录参数验证规则.
	ValidationRules = g.Map {
		"username": "required",
		"password": "required",
	}
)

// 初始化函数，
// 重写此函数以自定义您自己的JWT设置。
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

// PayloadFunc 是将在登录期间调用的回调函数。
// 使用此函数可以向webtoken添加额外的有效负载数据。
// 然后，在通过c.Get（“JWT_PAYLOAD”）请求时，可以使用这些数据。
// 请注意，有效负载没有加密。
// 上面提到的属性智威汤逊不能用作地图的键。
// 可选，默认情况下不设置其他数据。
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

// IdentityHandler 设置JWT的标识。
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// Unauthorized 用于定义自定义的未授权回调函数。
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  message,
	})
	r.ExitAll()
}

// LoginResponse 用于定义自定义登录成功回调函数。
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// RefreshResponse 用于获取新令牌，无论当前令牌是否过期。
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}


// Authenticator 用于验证登录参数。
// 它必须返回用户数据作为用户标识符，并将其存储在Claim Array中。
// 检查错误（e），以确定适当的错误消息。
func Authenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetMap()
	if e := gvalid.CheckMap(data, ValidationRules); e != nil {
		return "", jwt.ErrFailedAuthentication
	}
	if data["username"] == "admin" && data["password"] == "admin" {
		return g.Map {
			"username": data["username"],
			"id":       data["username"],
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

```


