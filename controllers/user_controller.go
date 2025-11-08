package controllers

import (
	"net/http"

	"github.com/Aum-Patel1234/researchq_backend/models"
	"github.com/Aum-Patel1234/researchq_backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []models.User

		if err := db.Find(&users).Error; err != nil {
			res := utils.JsonResponse("", "Users not found", 0, false)
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		ctx.JSON(http.StatusOK, utils.JsonResponse("", "Users Fetched Succesfully!", users, true))
	}
}
