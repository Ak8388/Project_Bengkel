package auth

import "github.com/dgrijalva/jwt-go"

var Form = "http://localhost:8080"
var Secret = "klmdmosiweei9weiepqw023"

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenDetail struct {
	AcToken string `json:"token_detail"`
	AcUuid  string `json:"ac_uid"`
	AtExp   int64  `json:"exp"`
}

type JwtClaims struct {
	jwt.StandardClaims
	UuId  int      `json:"uid"`
	Group []string `json:"group"`
}
