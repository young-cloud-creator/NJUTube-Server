package main

import (
	"github.com/gin-gonic/gin"
	"goto2023/repository"
	"log"
)

func main() {
	// init the database
	err := repository.InitDB()
	if err != nil {
		log.Fatal(err) // cannot connect to the database
	}

	// setup routes
	router := gin.Default()
	initRouter(router)

	// listen and serve on 0.0.0.0:8080
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err) // cannot start server
	}
}
