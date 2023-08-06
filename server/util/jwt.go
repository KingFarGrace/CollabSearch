package util

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/cristalhq/jwt/v5"
)

// JWTValidate can parse a raw token and verify it.
func JWTValidate(rawToken, salt string) *jwt.Token {
	verifier, err1 := jwt.NewVerifierHS(jwt.HS256, []byte(salt))
	if err1 != nil {
		WarnLogger("JWTValidate()", err1.Error())
		return nil
	}
	token, err2 := jwt.Parse([]byte(rawToken), verifier)
	if err2 != nil {
		WarnLogger("JWTValidate()", err2.Error())
		return nil
	}
	return token
}

// GenerateJWT can generate a JWT.
func GenerateJWT(email, salt string) *jwt.Token {
	signer, err1 := jwt.NewSignerHS(jwt.HS256, []byte(salt))
	if err1 != nil {
		ErrorLogger(err1, "GenerateJWT()")
		return nil
	}
	builder := jwt.NewBuilder(signer)
	claims := &entity.TokenClaims{
		Email: email,
	}
	token, err2 := builder.Build(claims)
	if err2 != nil {
		ErrorLogger(err2, "GenerateJWT()")
		return nil
	}
	return token
}
