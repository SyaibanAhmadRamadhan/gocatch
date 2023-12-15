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

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gsql"
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

	createTable := `CREATE TABLE IF NOT EXISTS users 
			(id serial PRIMARY KEY, username VARCHAR ( 50 ) NOT NULL, password VARCHAR ( 50 ) NOT NULL, 
    		email VARCHAR ( 255 ) NOT NULL, created_on TIMESTAMP NOT NULL, last_login TIMESTAMP);`
	ctx := context.Background()
	pgxCommander := NewPgxCommander(pool)
	sqlxComannder := gsql.NewSqlxCommander(db)
	txPgx := NewTxPgx(pool)
	txSqlx := gdb.NewTxSqlx(db)

	err := txPgx.DoTransaction(ctx, nil, func(c context.Context) error {
		_, err := pgxCommander.Exec(c, createTable)
		return err
	})
	gcommon.PanicIfError(err)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			err = txPgx.DoTransaction(ctx, nil, func(c context.Context) error {
				_, err := pgxCommander.Exec(c, `INSERT INTO users (username, password, email, created_on, last_login) 
																		VALUES ('test', 'test', '', NOW(), NOW());`)
				return err
			})
			if err != nil {
				fmt.Println(err)
			}

			err = txSqlx.DoTransaction(ctx, nil, func(c context.Context) error {
				_, err := sqlxComannder.ExecContext(c, `INSERT INTO users (username, password, email, created_on, last_login) 
																		VALUES ('test', 'test', '', NOW(), NOW());`)
				return err
			})
			if err != nil {
				fmt.Println(err)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	row := pgxCommander.QueryRow(context.Background(), "SELECT COUNT(*) FROM users;")
	var count int
	err = row.Scan(&count)
	fmt.Println(count)

	rowsqlx, err := sqlxComannder.QueryxContext(context.Background(), "SELECT COUNT(*) FROM users;")
	var countSqlx int
	for rowsqlx.Next() {
		err = rowsqlx.Scan(&countSqlx)
	}
	fmt.Println(countSqlx)
}
