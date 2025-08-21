package middleware

import (
    "net/http"
    "strings"
    "cms-api/database"
    "cms-api/models"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("SUPERSECRETKEY")

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return JwtSecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        var user models.User
        database.DB.First(&user, claims["user_id"])
        c.Set("user", user)

        c.Next()
    }
}
