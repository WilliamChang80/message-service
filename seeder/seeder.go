package seeder

import (
	"database/sql"
	"fmt"

	"github.com/bxcodec/faker/v3"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

func Run() {
	cs := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", cs)
	if err != nil {
		panic(err)
	}

	i := 0
	for i < 10000000 {

		sqlStatement := fmt.Sprintf(`INSERT INTO messages (created_at, updated_at, content, sent_at, sender_id, room_id)
VALUES (now(), now(), '%s', now(), %d, %d)`, faker.Sentence(), 999000000, 999000000)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
		i++
	}
	fmt.Println("Success insert data!")
	defer db.Close()
}
