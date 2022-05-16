package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/williamChang80/message-service/appcontext"
	"github.com/williamChang80/message-service/config"
)

func Migrate() {
	appcontext.Init()
	db := appcontext.GetDb()

	driver, err := cassandra.WithInstance(db, &cassandra.Config{
		KeyspaceName: config.GetDbKeyspace(),
	})
	if err != nil {
		log.Fatalf("Migration failed with message: %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://migration", "messageservice", driver)

	if err != nil {
		log.Fatalf("Migration failed with message: %s", err.Error())
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("Error with message: %s", err.Error())
		log.Println("Rolling back...")
		m.Down()
	}

	defer db.Close()

	log.Print("migration success")
}
