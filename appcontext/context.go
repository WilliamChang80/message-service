package appcontext

import (
	"github.com/gocql/gocql"
	"github.com/williamChang80/message-service/config"
)

type appContext struct {
	db *gocql.Session
}

var context *appContext

func Init() {
	config.Init()

	if context == nil {
		context = &appContext{
			db: config.InitDb(),
		}
	}
}

func GetDb() *gocql.Session {
	return context.db
}
