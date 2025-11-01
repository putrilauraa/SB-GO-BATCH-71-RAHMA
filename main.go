package main

import (
	"SB-GO-BATCH-71-RAHMA/controllers"
	"SB-GO-BATCH-71-RAHMA/database"
	"SB-GO-BATCH-71-RAHMA/routers"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Cannot load .env file")
	}

	db := database.ConnectDatabase()

	bioskopController := controllers.BioskopController{DB: db}

	r := routers.SetupRouter(&bioskopController)

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	log.Println("Server berjalan di port : " + port)
	
	r.Run(":" + port)
}