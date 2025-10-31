package routers

import (
	"github.com/gin-gonic/gin"
	"SB-GO-BATCH-71-RAHMA/controllers"
)

func SetupRouter(bioskopCtrl *controllers.BioskopController) *gin.Engine {
	r := gin.Default()

	r.POST("/bioskop", bioskopCtrl.CreateBioskop)

	return r
}