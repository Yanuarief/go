package routers

import(
	"github.com/gin-gonic/gin"

	"tutorial2/controllers"
)

// Routers is...
func Routers(){
	router := gin.Default()

	// Router Controller1
	router.GET("/gets", controllers.GetData)
	router.GET("/gets/:id", controllers.GetDataByID)
	router.POST("/gets", controllers.InData)
	router.PATCH("/gets/:id", controllers.UpData)
	router.DELETE("/gets/:id", controllers.DelData)

	router.Run(":4000")
}