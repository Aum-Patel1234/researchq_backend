package routes

import (
	"net/http"

	"github.com/Aum-Patel1234/researchq_backend/controllers"
	"github.com/Aum-Patel1234/researchq_backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, utils.JsonResponse("", "pong", 1, true))
	})

	v1.GET("/err", func(ctx *gin.Context) {
		response := utils.ResponseJson[map[string]string]{
			Error:   "Something went Wrong",
			Message: "This is a dummy error response for testing",
			Result:  map[string]string{"foo": "bar"},
			Success: false,
		}

		ctx.JSON(http.StatusBadRequest, response)
	})

	v1.GET("/users", controllers.GetAllUsers(db))
	// v1.POST("/users", controllers.CreateUser(db))
	v1.PUT("/users/:id", controllers.UpdateUser(db))
	v1.DELETE("/users/:id", controllers.DeleteUser(db))

	v1.POST("/signup", controllers.CreateUser(db))
	v1.POST("/login", controllers.Login(db))

	// v1.GET("/subjects", controllers.GetSubjects(db))

	return router
}
