package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Aum-Patel1234/researchq_backend/models"
	"github.com/Aum-Patel1234/researchq_backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []models.User

		if err := db.Find(&users).Error; err != nil {
			res := utils.JsonResponse(err.Error(), "Users not found", 0, false)
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		ctx.JSON(http.StatusOK, utils.JsonResponse("", "Users Fetched Succesfully!", users, true))
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.CreateUserRequest

		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.JsonResponse("BIND_ERROR", "Invalid request data", 0, false))
			return
		}

		var exists models.User
		if err := db.Where("email = ?", req.Email).First(&exists).Error; err == nil {
			ctx.JSON(http.StatusConflict, utils.JsonResponse("EMAIL_EXISTS", "Email aldready registered", 0, false))
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse(err.Error(), "DB error", 0, false))
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("HASH_ERROR", "failed to hash password", 0, false))
			return
		}

		user := models.User{
			Email:    req.Email,
			Password: string(hashed),
			Name:     req.Name,
		}

		if err := db.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", err.Error(), 0, false))
			return
		}

		// fmt.Println(user)
		ctx.JSON(http.StatusCreated, utils.JsonResponse("", "User created Succesfully!", 1, true))
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.JsonResponse("INVALID_ID", "invalid user id", 0, false))
			return
		}

		var req models.UpdateUserRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.JsonResponse("BIND_ERROR", "Invalid request data", 0, false))
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, utils.JsonResponse("NOT_FOUND", "user not found", 0, false))
			} else {
				ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", err.Error(), 0, false))
			}
			return
		}

		if req.Email != nil && *req.Email != "" {
			user.Email = *req.Email
		}
		if req.Name != nil && *req.Name != "" {
			user.Name = *req.Name
		}

		if req.Password != nil && *req.Password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("HASH_ERROR", "failed to hash password", 0, false))
				return
			}
			user.Password = string(hash)
		}

		// IMPORTANT: update and save
		// db.Save updates all fields — even unchanged ones; consider using db.Updates for partial updates to improve efficiency.
		// db.Updates allows selective field updates; ideal for PATCH requests or when handling optional fields.
		if err := db.Save(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", err.Error(), 0, false))
			return
		}

		ctx.JSON(http.StatusOK, utils.JsonResponse("USER_UPDATED", "User updated successfully", 1, true))
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.JsonResponse("INVALID_ID", "invalid user id", 0, false))
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, utils.JsonResponse("USER_NOT_FOUND", "user not found", 0, false))
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", err.Error(), 0, false))
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.JsonResponse("DB_ERROR", "user deletion failed", 0, false))
			return
		}

		ctx.JSON(http.StatusOK, utils.JsonResponse("", "User Deleted Succesfully!", 1, false))
	}
}
