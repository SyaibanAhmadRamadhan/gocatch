package gMYSQL

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/gdb/gsql"
)

func TestMysqlDockerTest(t *testing.T) {
	dockerTest := gdb.InitDockerTest()
	defer dockerTest.CleanUp()

	mysqlDockerTestConf := MysqlDockerTestConf{}

	var db *sqlx.DB
	var err error

	dockerTest.NewContainer(mysqlDockerTestConf.ImageVersion(dockerTest, true, ""), func(res *dockertest.Resource) error {
		time.Sleep(10 * time.Second)

		db, err = mysqlDockerTestConf.ConnectSqlx(res)
		gcommon.PanicIfError(err)
		return nil
	})

	asd := gsql.NewSqlxCommander(db)

	_ = asd.BeginTxRun(context.Background(), nil, func(tx gsql.Commander) error {
		_, err := tx.ExecContext(context.Background(),
			"CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, username VARCHAR ( 50 ) NOT NULL, password VARCHAR ( 50 ) NOT NULL, email VARCHAR ( 255 ) NOT NULL, created_on TIMESTAMP NOT NULL, last_login TIMESTAMP);")
		gcommon.PanicIfError(err)
		return nil
	})

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			_ = asd.BeginTxRun(context.Background(), nil, func(tx gsql.Commander) error {
				_, err := tx.ExecContext(context.Background(), "INSERT INTO users (username, password, email, created_on, last_login) VALUES ('test', 'test', '', NOW(), NOW());")
				gcommon.PanicIfError(err)
				return nil
			})
			wg.Done()
		}()
	}

	wg.Wait()

	rows, err := asd.QueryxContext(context.Background(), "SELECT * FROM users;")
	fmt.Println(rows)
	gcommon.PanicIfError(err)

	rows, err = asd.QueryxContext(context.Background(), "SELECT COUNT(*) FROM users;")
	gcommon.PanicIfError(err)
	var count int

	for rows.Next() {
		err = rows.Scan(&count)
		gcommon.PanicIfError(err)

		break
	}
	fmt.Println(count)

}
