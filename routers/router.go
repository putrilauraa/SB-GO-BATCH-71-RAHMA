package routers

import (
	"github.com/gin-gonic/gin"
	"SB-GO-BATCH-71-RAHMA/controllers"
)

func SetupRouter(bioskopCtrl *controllers.BioskopController) *gin.Engine {
	r := gin.Default()

	bioskopAPI := r.Group("/bioskop")
	{
		bioskopAPI.POST("", bioskopCtrl.CreateBioskop)
		bioskopAPI.GET("", bioskopCtrl.GetBioskops)
		bioskopAPI.GET("/:id", bioskopCtrl.GetBioskopByID)
		bioskopAPI.PUT("/:id", bioskopCtrl.UpdateBioskop)
		bioskopAPI.DELETE("/:id", bioskopCtrl.DeleteBioskop)
	}

	return r
}