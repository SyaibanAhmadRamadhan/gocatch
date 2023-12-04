package gpostgre

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"

	"github.com/SyaibanAhmadRamadhan/gocatch/gdb/gsql"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func TestPostgresDockerTest(t *testing.T) {
	dockerTest := ginfra.InitDockerTest()
	defer dockerTest.CleanUp()

	postgresDockerTest := PostgresDockerTestConf{}

	var pool *pgxpool.Pool
	var db *sqlx.DB

	dockerTest.NewContainer(postgresDockerTest.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
		time.Sleep(2 * time.Second)
		conn, err := postgresDockerTest.ConnectPgx(res)
		pool = conn
		gcommon.PanicIfError(err)

		db, err = postgresDockerTest.ConnectSqlx(res)
		gcommon.PanicIfError(err)
		return nil
	})

	asd := NewRdbmsPgx(pool)
	asd2 := gsql.NewSqlxCommander(db)

	_ = asd.BeginRun(context.Background(), func(tx Commander) error {
		_, err := tx.Exec(context.Background(),
			"CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, username VARCHAR ( 50 ) NOT NULL, password VARCHAR ( 50 ) NOT NULL, email VARCHAR ( 255 ) NOT NULL, created_on TIMESTAMP NOT NULL, last_login TIMESTAMP);")
		gcommon.PanicIfError(err)
		return nil
	})

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			_ = asd.BeginRun(context.Background(), func(tx Commander) error {
				_, err := tx.Exec(context.Background(), "INSERT INTO users (username, password, email, created_on, last_login) VALUES ('test', 'test', '', NOW(), NOW());")
				gcommon.PanicIfError(err)
				return nil
			})

			_ = asd2.BeginTxRun(context.Background(), nil, func(tx gsql.Commander) error {
				_, err := tx.ExecContext(context.Background(), "INSERT INTO users (username, password, email, created_on, last_login) VALUES ('test', 'test', '', NOW(), NOW());")
				gcommon.PanicIfError(err)
				return nil
			})
			wg.Done()
		}()
	}

	wg.Wait()
	rows, err := asd.Query(context.Background(), "SELECT * FROM users;")
	gcommon.PanicIfError(err)
	fmt.Println(rows)

	rows2, err := asd2.QueryxContext(context.Background(), "SELECT * FROM users;")
	gcommon.PanicIfError(err)

	row := asd.QueryRow(context.Background(), "SELECT COUNT(*) FROM users;")
	var count int
	err = row.Scan(&count)
	fmt.Println(rows2)
	fmt.Println(count)

}
