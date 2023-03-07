package jwt

import (
	"project_bengkel/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func Token(id int) (*auth.TokenDetail, error) {
	td := &auth.TokenDetail{}
	td.AtExp = time.Now().Add(time.Hour).Unix()
	uuID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	td.AcUuid = uuID.String()

	acClaims := auth.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    auth.Form,
			ExpiresAt: td.AtExp,
			NotBefore: time.Now().Unix(),
			Id:        uuID.String(),
			IssuedAt:  time.Now().Unix(),
		},
		UuId: id,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		acClaims,
	)

	sigAc, ernew := token.SignedString([]byte(auth.Secret))

	if ernew != nil {
		return nil, ernew
	}

	td.AcToken = sigAc

	return td, nil
}
