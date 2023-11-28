package example

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/SyaibanAhmadRamadhan/jolly"
	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOpg"
)

var user1 = map[string]any{
	"username":     "rama1",
	"password":     "pwrama1",
	"phone_number": "hprama1",
}

var user2 = map[string]any{
	"username":     "rama2",
	"password":     "pwrama2",
	"phone_number": "hprama2",
}

func TestUser(t *testing.T) {
	confPG := JOpg.PostgresConf{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     5432,
		DB:       "jolly",
		SSL:      "disable",
	}

	conn := JOpg.PgxNewConnection(confPG)

	rdbms := JOpg.NewRDBMSpgx(conn)

	_ = rdbms.BeginRun(context.Background(), func(rdbms JOpg.RDBMS) error {
		_, err := rdbms.Write(context.Background(), "INSERT INTO public.user (username, password, phone_number) VALUES (@username, @password, @phone_number)",
			pgx.NamedArgs(user1))
		jolly.PanicIF(err)

		_, err = rdbms.Write(context.Background(), "INSERT INTO public.user (username, password, phone_number) VALUES (@username, @password, @phone_number)",
			pgx.NamedArgs(user2))
		jolly.PanicIF(err)

		return errors.New("asd")
	})

	rows, err := rdbms.QueryAll(context.Background(), "SELECT phone_number, password, username FROM public.user")
	jolly.PanicIF(err)

	var userModels []User
	for _, row := range rows {
		userModel := NewUserWithOutPtr()
		err = userModel.Scan(row)
		jolly.PanicIF(err)

		userModels = append(userModels, userModel)
	}
	fmt.Println(userModels)

	row, err := rdbms.QueryOne(context.Background(), "SELECT phone_number, password, username FROM public.user WHERE username = 'tidak ada'")
	jolly.PanicIF(err)

	userModel := NewUserWithOutPtr()
	err = userModel.Scan(row)
	jolly.PanicIF(err)
	fmt.Println(userModel)

	exist, err := rdbms.CheckOne(context.Background(), "SELECT EXISTS(SELECT 1 FROM public.user)")
	jolly.PanicIF(err)
	fmt.Println(exist)
}
