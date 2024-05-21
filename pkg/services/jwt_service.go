package services

import (
	"Marketplace/pkg/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func GenerateAccessToken(user *models.User, secretKey string) (string, error) {

	payload := jwt.MapClaims{
		"sub":    user.Email,
		"userId": user.Id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseClaims(tokenString string) (jwt.MapClaims, error) {
	// надеюсь на то, что если есть префикс, то я удаою его
	// а если префикса нет, то строка останется, как есть
	tokenString, _ = strings.CutPrefix(tokenString, "Bearer ")
	token, _ := jwt.Parse(tokenString, TokenFunc)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrTokenRequiredClaimMissing
	}
}

func GetUserId(tokenString string) (int, error) {
	claims, err := ParseClaims(tokenString)

	if err != nil {
		return -1, jwt.ErrTokenInvalidClaims
	}

	if userIdFloat, ok := claims["userId"].(float64); ok {
		userId := int(userIdFloat)
		return userId, nil
	} else {
		fmt.Printf("Type of num: %T\n", claims["userId"])
		return 0, jwt.ErrTokenRequiredClaimMissing
	}
}

func TokenFunc(token *jwt.Token) (interface{}, error) {
	return []byte("439f262a5145e0c4194df51c44c4f5b253c2c69db67835b80e2ed3ccdc67270f73ce10c264c9c5a8798d00de2b3deca5465d2af1de7aa29a01dc7f2586dbab50b702f30038fe0d570a87e48db789095a68ec245efe009a7338e44088c7002b96b26e9dd7172afc41ce7fd19ec32caab790a10ed4348887aa6eb1f4dc9726341b11b3ba89c5f98c1907c2fac7676105bd930983f9e05e74618096b2eb10115f1d73ae4065315261ac8554a45fe007df194e3f7c95cdf0e204eb4dd4ff79e42371708be27cd268fe75ddbe939fc36cf09ecfc8acce5d7f67c6557daacda72b51edc412ad91ace6133af1bf8ec0cbbc088c97a5374328805a3ff9fa2a21737c6cc8"), nil
}
