package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Aum-Patel1234/researchq_backend/models"
	"github.com/Aum-Patel1234/researchq_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie(utils.CookieName)
		if err != nil || tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.JsonResponse("AUTH_ERROR", "Missing or invalid authentication cookie", 0, false))
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			jwt.MapClaims{},
			func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
			jwt.WithValidMethods([]string{utils.JwtAlgo.Alg()}),
		)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.JsonResponse("JWT_INVALID", "Failed to parse or validate JWT token", err.Error(), false))
			return
		}

		// 3️⃣Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.JsonResponse("JWT_INVALID", "Invalid token claims", 0, false))
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.JsonResponse("JWT_EXPIRED", "Token has expired, please log in again", 0, false))
			return
		}

		userID, ok := claims["sub"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.JsonResponse("JWT_SUB_ERROR", "Invalid user ID in token", 0, false))
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.JsonResponse("USER_NOT_FOUND", "User not found in the database", 0, false))
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", err.Error(), 0, false))
			}
			return
		}

		// Set user in context and continue
		ctx.Set("user", user)
		ctx.Next()
	}
}
