package controllers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Aum-Patel1234/researchq_backend/models"
	"github.com/Aum-Patel1234/researchq_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetSubjects(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var subjects []models.Subject

		if err := db.Find(&subjects).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse(err.Error(), "Subject fetch failed", 0, false))
			return
		}

		ctx.JSON(http.StatusOK, utils.JsonResponse("", "Subjects Fetched Successfully!", subjects, true))
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
		if len(hmacSampleSecret) == 0 {
			log.Println("⚠️ JWT_SECRET not set in environment")
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("CONFIG_ERROR", "server misconfigured", 0, false))
			return
		}

		var req models.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.JsonResponse("BIND_ERROR", "invalid request body", 0, false))
			return
		}

		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusUnauthorized, utils.JsonResponse("INVALID_CREDENTIALS", "user not found", 0, false))
			} else {
				ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", err.Error(), 0, false))
			}
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			ctx.JSON(http.StatusUnauthorized, utils.JsonResponse("INVALID_CREDENTIALS", "email or password is incorrect", 0, false))
			return
		}

		now := time.Now().UTC()
		if err := db.Model(&user).UpdateColumn("last_login", now).Error; err != nil {
			// log.Printf("failed to update last_login: %v", err)
		}
		user.LastLogin = now

		// send jwt token

		claims := jwt.MapClaims{
			"sub": user.ID,
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			// log.Printf("failed to sign jwt for user %v: %v", user.ID, err)
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("JWT_ERROR", "failed to generate token", 0, false))
			return
		}

		payload := gin.H{
			"id":    user.ID,
			"token": tokenString,
		}
		ctx.JSON(http.StatusOK, utils.JsonResponse("", "User Logged in Successfully!", payload, true))
	}
}
