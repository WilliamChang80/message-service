package main

import (
	"os"

	"github.com/williamChang80/message-service/migration"
	"github.com/williamChang80/message-service/seeder"
	"github.com/williamChang80/message-service/server"
)

func main() {
	args := os.Args[1]
	switch args {
	case "seed":
		seeder.Run()
	case "start":
		server.StartAPIServer()
	case "migrate":
		migration.Migrate()
	}
}
