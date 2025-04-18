package main


import (
	"os"
	"github.com/gin-gonic/gin"
	routes "go-jwt-project/routes"
)

func main(){
	port = os.GETENV("PORT")

	if port==""{
		port="8000"
	}
	router := gin.new()
	router.User(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success": "access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})
	router.Run{":" +port}
}