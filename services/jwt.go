package service

import (
	"context"
	"errors"
	"go-template/global"
	"go-template/models"
	"go-template/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 所有需要颁发 token 的用户模型必须实现这个接口

type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// CreateToken 生成 Token
func CreateToken(GuardName string, user *models.User) (TokenOutPut, error) {
	var tokenData TokenOutPut
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl,
				Id:        strconv.Itoa(user.ID),
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))
	if err != nil {
		return tokenData, err
	}

	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		global.TokenType,
	}
	return tokenData, nil
}

// 获取黑名单key
func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + utils.MD5(tokenStr)
}

// JoinBlackList token 加入黑名单
func JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	err = global.App.Redis.SetNX(context.Background(), getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

// IsInBlacklist token 是否在黑名单中
func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(context.Background(), getBlackListKey(tokenStr)).Result()
	if err != nil {
		return false
	}
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true
}

// 根据token ，查询用户表数据
func GetUserInfo(GuardName string, id int) (user *models.User, err error) {
	switch GuardName {
	case global.AppGuardName:
		return GetUser(id)
	default:
		err = errors.New("guard " + GuardName + " does not exist")
	}
	return
}
