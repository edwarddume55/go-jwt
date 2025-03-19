package routes

import(
	"github.com/gin-gonic/gin"
	controller "go-jwt-project/controllers"

)

func AuthRoutes(incommingRoutes *gin.Engine){
	incommingRoutes.POST("users/signup", controller.Signup())
	incommingRoutes.POST("users/login", controller.Login())
}