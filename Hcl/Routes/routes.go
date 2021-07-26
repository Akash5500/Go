package Routes

import (
	"Hcl/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetHotels)
		grp1.POST("user", Controllers.CreateHotel)
		grp1.GET("user/:id", Controllers.GetHotelByID)
		grp1.PUT("user/:id", Controllers.UpdateHotel)
		grp1.DELETE("user/:id", Controllers.Deletehotel)
	}
	return r
}
