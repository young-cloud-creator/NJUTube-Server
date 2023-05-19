package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"nju-tube/repository"
	"nju-tube/service"
	"os"
)

// create public directories
func initPublicPath() {
	err := os.Mkdir(service.PublicDir, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}

	err = os.Mkdir(service.VideoDir, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}

	err = os.Mkdir(service.CoverDir, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	// init the public dir
	initPublicPath()

	// init the database
	err := repository.InitDB()
	if err != nil {
		log.Fatal(err) // cannot connect to the database
	}

	// setup routes
	router := gin.Default()
	initRouter(router)

	addr := "0.0.0.0:8080"
	server := http.Server{Addr: addr, Handler: router}
	ln, _ := net.Listen("tcp4", addr)
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	_ = server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})

	/* will run on IPv6 address if server has
	// listen and serve on 0.0.0.0:8080
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err) // cannot start server
	}
	*/
}
