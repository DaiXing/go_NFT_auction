package web

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"my.nft.auction/src/util"
)

// 生成 JWT token
func jwtBuildToken(userId uint64, username string) string {
	// 最重要的是，过期时间，userId
	tokenInfo := JwtTokenInfo{
		UserId:   userId,   // 业务参数
		Username: username, // 业务参数
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute *
				time.Duration(util.Params.Web.TokenValidMiutes))), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		}}

	tmp := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenInfo)

	jwtKeys := []byte(util.Params.Web.JwtKey) // 加密的key

	token, err := tmp.SignedString(jwtKeys) // 加密。

	util.CheckError(err)
	util.Logger.Info("生成jwt ", "token", token)
	return token
}

// jwt 验证token
func jwtVerifyToken(token string) *JwtTokenInfo {
	if len(token) == 0 {
		panic("token empty")
	}
	jwtKeys := []byte(util.Params.Web.JwtKey) // 加密的key
	// 解析token，验证token是否合法。
	// 返回 *Token
	token2, err2 := jwt.ParseWithClaims(
		token,
		&JwtTokenInfo{},
		func(t *jwt.Token) (any, error) {
			return jwtKeys, nil
		})
	util.CheckError(err2)

	// 转为 TokenInfo
	tokenInfo, ok := token2.Claims.(*JwtTokenInfo)
	if !ok {
		panic("jwt token parse error ")
	}
	return tokenInfo
}
