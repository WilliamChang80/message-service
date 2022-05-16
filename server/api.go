package server

import (
	"log"
	"net/http"

	"github.com/williamChang80/message-service/appcontext"
	"github.com/williamChang80/message-service/config"
	"github.com/williamChang80/message-service/router"
)

func listenServer(apiServer *http.Server) {
	log.Printf("Started server at %v", apiServer.Addr)

	if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}

func StartAPIServer() {
	config.Init()
	appcontext.Init()

	defer appcontext.GetDb().Close()

	r := router.InitRouter()

	server := &http.Server{Addr: ":" + config.GetAppPort(), Handler: r}
	listenServer(server)
}
