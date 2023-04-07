package middleware

import (
	"go-template/common"
	"go-template/global"
	service "go-template/services"
	"strconv"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		// TODO: 判断不需要登录验证的接口(写法待优化，现阶段使用golang-set来判断是否不需要登录认证中)
		var noJWT = []interface{}{
			"/api/v1/login",
		}
		check := mapset.NewSetFromSlice(noJWT)
		if check.Contains(c.Request.URL.Path) {
			c.Next()
			return
		}
		if tokenStr == "" {
			common.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(global.TokenType)+1:]

		// 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil || service.IsInBlacklist(tokenStr) {
			common.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.CustomClaims)

		// 发布者校验
		if claims.Issuer != GuardName {
			common.TokenFail(c)
			c.Abort()
			return
		}

		// 续签
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				id, _ := strconv.Atoi(claims.Id)
				user, err := service.GetUserInfo(GuardName, id)
				if err != nil {
					global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _ := service.CreateToken(GuardName, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = service.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
