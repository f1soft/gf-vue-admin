##  Json Web Token

## 1.Simple configuration

| Name              | Type          | Example                                            | Description                                                  |
| :---------------- | ------------- | :------------------------------------------------- | :----------------------------------------------------------- |
| Realm             | string        | "test zhou"                                        | Realm name to display to the user. Required.                 |
| SigningAlgorithm  | string        | "HS256"                                            | signing algorithm - possible values are HS256, HS384, HS512 ,Optional, default is HS256. |
| Key               | []byte        | []byte("secret key")                               | Secret key used for signing. Required.                       |
| Timeout           | time.Duration | time.Minute * 5                                    | Duration that a jwt token is valid. Optional, defaults to one hour. |
| MaxRefresh        | time.Duration | time.Minute * 5                                    | This field allows clients to refresh their token until MaxRefresh has passed. Note that clients can refresh their token in the last moment of MaxRefresh. This means that the maximum validity timespan for a token is TokenTime + MaxRefresh. Optional, defaults to 0 meaning not refreshable. |
| IdentityKey       | string        | "id"                                               | Set the identity key                                         |
| TokenLookup       | string        | "header: Authorization, query: token, cookie: jwt" | TokenLookup is a string in the form of "<source>:<name>" that is used to extract token from the request. // Optional. Default value "header:Authorization".  Possible values: <br />"header:<name>" <br />"query:<name>"<br />"cookie:<name>" |
| TokenHeadName     | string        | "Bearer"                                           | TokenHeadName is a string in the header. Default value is "Bearer" |
| PrivKeyFile       | string        |                                                    | Private key file for asymmetric algorithms                   |
| PubKeyFile        | string        |                                                    | Public key file for asymmetric algorithms                    |
| SendCookie        | bool          |                                                    | Optionally return the token as a cookie                      |
| SecureCookie      | bool          |                                                    | Allow insecure cookies for development over http             |
| CookieHTTPOnly    | bool          |                                                    | Allow cookies to be accessed client side for development     |
| CookieDomain      | string        |                                                    | Allow cookie domain change for development                   |
| SendAuthorization | bool          |                                                    | SendAuthorization allow return authorization header for every request |
| DisabledAbort     | bool          |                                                    | Disable abort() of context.                                  |
| CookieName        | string        |                                                    | CookieName allow cookie name change for development          |

## 2. Callback function

### 2.1 Authenticator

- Type:==func(r *ghttp.Request) (interface{}, error)==
- Description: 
	- Callback function that should perform the authentication of the user based on login info.
	- Must return user data as user identifier, it will be stored in Claim Array. Required.
	- Check error (e) to determine the appropriate error message.

### 2.2 Authorizator 

- Type:==func(data interface{}, r *ghttp.Request) bool==
- Description: 
	- Callback function that should perform the authorization of the authenticated user. 
	- Called only after an authentication success. Must return true on success, false on failure.
	- Optional, default to success.
- 

### 2.3 PayloadFunc

- Type:==func(data interface{}) MapClaims==
- Description: 
	- Callback function that will be called during login.
	- Using this function it is possible to add additional payload data to the webtoken.
	- The data is then made available during requests via c.Get("JWT_PAYLOAD").
	- Note that the payload is not encrypted.
	- The attributes mentioned on jwt.io can't be used as keys for the map.
	- Optional, by default no additional data will be set.

### 2.4 Unauthorized

- Type:==func(*ghttp.Request, int, string)==
- Description: 
	- User can define own Unauthorized func.

### 2.5 LoginResponse

- Type:==func(*ghttp.Request, int, string, time.Time)==
- Description: 
	- User can define own LoginResponse func.

### 2.6 RefreshResponse

- Type:==func(*ghttp.Request, int, string, time.Time)==
- Description: 
	- User can define own RefreshResponse func.

### 2.7 IdentityHandler

- Type:==func(*ghttp.Request) interface{}==
- Description: 
	- Set the identity handler function

## 3. Sample code

```go
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
	// The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
	// Customized login parameter validation rules.
	ValidationRules = g.Map {
		"username": "required",
		"password": "required",
	}
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
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

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
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
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  message,
	})
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
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


