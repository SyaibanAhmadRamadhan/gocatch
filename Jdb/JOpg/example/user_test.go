package example

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/jolly"
	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOpg"
	"github.com/SyaibanAhmadRamadhan/jolly/Jsql"
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

var pool *pgxpool.Pool
var rdbms JOpg.RDBMS

func initConn() {
	confPG := JOpg.PostgresConf{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     5432,
		DB:       "jolly",
		SSL:      "disable",
	}
	pool = JOpg.PgxNewConnection(confPG, true)

	rdbms = JOpg.NewRDBMSpgx(pool)
}

func create(user *User) {
	_ = rdbms.BeginRun(context.Background(), func(rdbms JOpg.RDBMS) error {
		value := user.FieldAndValue().JoinKey(string(Jsql.Pgx))
		column := user.FieldAndValue().JoinKey()
		sql := "INSERT INTO public.user (" + column + ") VALUES (" + value + ")"
		fmt.Println(sql)
		_, err := rdbms.Write(context.Background(), sql, pgx.NamedArgs(user.FieldAndValue()))
		jolly.PanicIF(err)

		return nil
	})
}

func update(user *User) {
	_ = rdbms.BeginRun(context.Background(), func(rdbms JOpg.RDBMS) error {
		args, value := user.QFilterNamedArgs.ToQuery(true, Jsql.Pgx)
		sql := "UPDATE public.user SET " + user.FieldArgForUpdate(Jsql.Pgx) + " " + args

		fmt.Println(sql)
		fieldValue := user.FieldAndValue()
		fieldValue.Merge(value)

		res, err := rdbms.Write(context.Background(), sql, pgx.NamedArgs(fieldValue))
		jolly.PanicIF(err)
		fmt.Println(res)

		return nil
	})
}

func setValueForCreate(user *User) {
	user.SetUsername("ibann")
	user.SetPassword("rama_password")
}

func setWhereCondition(user *User) {
	user.SetArgFieldPassword("rama_password", Jsql.Equals, Jsql.And)
}

func TestUser(t *testing.T) {
	initConn()

	newUser := NewUser()
	newUser.SetColumn(
		newUser.FieldPassword(),
	)
	setValueForCreate(newUser)
	t.Log(newUser.QColumnFields)
	setWhereCondition(newUser)

	// create(newUser)
	update(newUser)
	// rows, err := rdbms.QueryAll(context.Background(), "SELECT phone_number, password, username FROM public.user")
	// jolly.PanicIF(err)
	//
	// var userModels []User
	// for _, row := range rows {
	// 	userModel := NewUserWithOutPtr()
	// 	err = userModel.Scan(row)
	// 	jolly.PanicIF(err)
	//
	// 	userModels = append(userModels, userModel)
	// }
	// fmt.Println(userModels)
	//
	// row, err := rdbms.QueryOne(context.Background(), "SELECT phone_number, password, username FROM public.user WHERE username = 'tidak ada'")
	// jolly.PanicIF(err)
	//
	// userModel := NewUserWithOutPtr()
	// err = userModel.Scan(row)
	// jolly.PanicIF(err)
	// fmt.Println(userModel)
	//
	// exist, err := rdbms.CheckOne(context.Background(), "SELECT EXISTS(SELECT 1 FROM public.user)")
	// jolly.PanicIF(err)
	// fmt.Println(exist)
}
