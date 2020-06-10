package auth

import (
	"fmt"
	userModel "gf-admin/app/model/system/user"
	"gf-admin/library/orm"
	"gf-admin/library/response"
	"time"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

var (
	// GfJWTMiddleware The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {

	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             g.Cfg().GetBytes("app.JwtSecret"),
		Timeout:         g.Cfg().GetDuration("app.JwtExpiresin") * time.Minute,
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
	response.Res(r).Unauthorized(message)
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.Res(r).Success(token)
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	u := r.Get("JWT_PAYLOAD")
	fmt.Printf("user is     %s \n", u)
	response.Res(r).Success(token)
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
// Delete 登录
// @Summary 登录
// @Description 登录
// @Tags 系统
// @accept json
// @Produce  json
// @Param data body user.LoginReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/login [post]
// @Security ApiKeyAuth
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var req *userModel.LoginReq
	if err := r.Parse(&req); err != nil {
		return nil, err
	}
	var user userModel.Entity
	db := orm.Instance()

	if has, err := db.Where("username = ?", req.Username).Get(&user); err != nil {
		return nil, err
	} else if !has {
		return nil, gerror.New("用户名或密码错误")
	}

	password := req.Password + user.Salt

	if gmd5.MustEncryptString(password) != user.Password {
		return nil, gerror.New("用户名或密码错误")
	}

	return g.Map{
		"username": user.Username,
		"id":       user.ID,
	}, nil

}

// MiddlewareAuth is the HOOK function implements JWT logistics.
func MiddlewareAuth(r *ghttp.Request) {
	GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
