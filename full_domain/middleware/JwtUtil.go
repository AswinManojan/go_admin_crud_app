package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtUtil struct{}

type Claims struct {
	Username string
	Role     string
	*jwt.StandardClaims
}

func (j *JwtUtil) CreateToken(username string, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte("121212"))
	if err != nil {
		fmt.Println("Error creating token")
	}
	return strToken, nil
}

func (j *JwtUtil) ValidateToken(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token not found",
			})
			return
		}

		tokenString = string([]byte(tokenString)[7:])
		claims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("121212"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token not valid",
			})
			return
		}
		if claims.Role != role || !parsedToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "user-admin mismatch-- unauthorized acces-- access denied",
			})
			return
		}
		c.Next()
	}
}

func NewJwtUtil() *JwtUtil {
	return &JwtUtil{}
}
