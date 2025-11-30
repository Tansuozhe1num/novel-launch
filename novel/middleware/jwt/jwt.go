package jwt

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"novel-launch/novel/Biu"
	redismw "novel-launch/novel/middleware/redis"

	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UID   uint64 `json:"uid"`
	UName string `json:"uname"`
	jwtv5.RegisteredClaims
}

func secret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "novel_secret_123"
	}
	return []byte(s)
}

func GenerateToken(uid uint64, uname string) (string, error) {
	claims := Claims{
		UID:   uid,
		UName: uname,
		RegisteredClaims: jwtv5.RegisteredClaims{
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
		},
	}
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString(secret())
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwtv5.ParseWithClaims(tokenStr, &Claims{}, func(token *jwtv5.Token) (interface{}, error) {
		return secret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			Biu.Failed(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		parts := strings.SplitN(h, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			Biu.Failed(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		claims, err := ParseToken(parts[1])
		if err != nil {
			Biu.Failed(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		if r := redismw.Get(); r != nil {
			v, _ := r.Get(context.Background(), "user:token:"+strconv.FormatUint(claims.UID, 10)).Result()
			if v == "" || v != parts[1] {
				Biu.Failed(c, http.StatusUnauthorized, "unauthorized")
				c.Abort()
				return
			}
		}
		c.Set("uid", claims.UID)
		c.Set("uname", claims.UName)
		c.Next()
	}
}
