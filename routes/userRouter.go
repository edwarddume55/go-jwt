package routes

import(
	"go-jwt-project/middleware"
	"github.com/gin-gonic/gin"
	controller "go-jwt-project/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine){
	incommingRoutes.User(middleware.Authenitcate())
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUser())
}