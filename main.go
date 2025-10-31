package main

import (
	"log"
	
	"SB-GO-BATCH-71-RAHMA/controllers"
	"SB-GO-BATCH-71-RAHMA/database"
	"SB-GO-BATCH-71-RAHMA/routers"
)

func main() {
	db := database.ConnectDatabase()

	bioskopController := controllers.BioskopController{DB: db}

	r := routers.SetupRouter(&bioskopController)

	log.Println("Server running on port :8080")
	r.Run(":8080")
}