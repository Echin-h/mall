package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = "mall"

// 这里使用了长短jwt的用法
func GenerateToken(id uint, username string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	rtExpireTime := nowTime.Add(7 * 24 * time.Hour)
	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}
	// 加密获得token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "mall",
	}).SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaim, err := ParseToken(aToken)
	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return
	}
	if accessClaim.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}
	if refreshClaim.ExpiresAt > time.Now().Unix() {
		return GenerateToken(refreshClaim.ID, refreshClaim.Username)
	}
	return "", "", errors.New("身份过期，重新登录")
}
