package example

// var user1 = map[string]any{
// 	"username":     "rama1",
// 	"password":     "pwrama1",
// 	"phone_number": "hprama1",
// }
//
// var user2 = map[string]any{
// 	"username":     "rama2",
// 	"password":     "pwrama2",
// 	"phone_number": "hprama2",
// }
//
// var pool *pgxpool.Pool
// var rdbms JOsql.RDBMS
//
// func initConn() {
// 	confPG := JOLpg.PostgresConf{
// 		User:     "root",
// 		Password: "root",
// 		Host:     "localhost",
// 		Port:     5432,
// 		DB:       "jolly",
// 		SSL:      "disable",
// 	}
// 	pool = JOLpg.PgxNewConnection(confPG, true)
//
// 	rdbms = JOLpg.NewRDBMSpgx(pool)
// }
//
// // func create(user *User) {
// // 	_ = rdbms.BeginRun(context.Background(), func(rdbms JOLpg.RDBMS) error {
// // 		value := user.FieldAndValue().JoinKey(string(JOsql.Pgx))
// // 		column := user.FieldAndValue().JoinKey()
// // 		sql := "INSERT INTO public.user (" + column + ") VALUES (" + value + ")"
// // 		fmt.Println(sql)
// // 		_, err := rdbms.Write(context.Background(), sql, pgx.NamedArgs(user.FieldAndValue()))
// // 		jolly.PanicIF(err)
// //
// // 		return nil
// // 	})
// // }
// //
// // func update(user *User) {
// // 	_ = rdbms.BeginRun(context.Background(), func(rdbms JOLpg.RDBMS) error {
// // 		args, value := user.FNamedArgs.ToQuery(true, JOsql.Pgx)
// // 		sql := "UPDATE public.user SET " + user.FieldArgForUpdate(JOsql.Pgx) + " " + args
// //
// // 		fmt.Println(sql)
// // 		fieldValue := user.FieldAndValue()
// // 		fieldValue.Merge(value)
// //
// // 		res, err := rdbms.Write(context.Background(), sql, pgx.NamedArgs(fieldValue))
// // 		jolly.PanicIF(err)
// // 		fmt.Println(res)
// //
// // 		return nil
// // 	})
// // }
//
// func setValueForCreate(user *User) {
// 	user.SetUsername("ibann")
// 	user.SetPassword("rama_password")
// }
//
// // func setWhereCondition(user *User) {
// // 	user.SetArgFieldPassword("rama_password", JOsql.Equals, JOsql.And)
// // }
//
// func selectField() {
// 	newUser := NewUser()
// 	newUser.RQFieldSet(
// 		newUser.FieldPassword(),
// 		newUser.FieldUsername(),
// 	)
//
// 	// newUser.FNamedArgsSet(
// 	// 	newUser.FNamedArgsRoleID("rama", JOsql.Equals, JOsql.And),
// 	// )
//
// 	newUser.SetID("rama")
// 	fmt.Println(newUser.RQField)
// 	fmt.Println(newUser.WCField)
// 	fmt.Println(newUser.FNamedArgs.ToQuery(true, JOsql.Pgx))
//
// }
//
// func TestUser(t *testing.T) {
// 	// selectField()
//
// 	newUser := NewUser()
//
// 	newUser.FNamedArgsSet(
// 		newUser.Where(newUser.FieldID(), JOsql.In, "name, rama"),
// 		newUser.OrWhere(newUser.FieldUsername(), JOsql.In, "name, rama"),
// 		newUser.Where(newUser.FieldID(), JOsql.NotIn, "name, rama"),
// 		newUser.OrWhere(newUser.FieldUsername(), JOsql.NotIn, "name, rama"),
// 	)
// 	newUser.FNamedArgsSet(
// 		newUser.Where(newUser.FieldID(), JOsql.IsNull, "name, rama"),
// 		newUser.OrWhere(newUser.FieldUsername(), JOsql.IsNotNull, "name, rama"),
// 		newUser.Where(newUser.FieldID(), JOsql.Like, "name, rama"),
// 		newUser.OrWhere(newUser.FieldUsername(), JOsql.NotLike, "name, rama"),
// 		newUser.Where(newUser.FieldID(), JOsql.Where, "name, rama"),
// 		newUser.OrWhere(newUser.FieldUsername(), JOsql.Where, "name, rama"),
// 		newUser.Where(JOstr.EmptyString, JOsql.FullTextSearch, "name, rama"),
// 	)
// 	t.Log(newUser.FNamedArgs.ToQuery(true, JOsql.Pgx))
// 	// initConn()
//
// 	// setValueForCreate(newUser)
// 	// t.Log(newUser.RQField)
// 	// setWhereCondition(newUser)
// 	//
// 	// // create(newUser)
// 	// update(newUser)
// 	// rows, err := rdbms.QueryAll(context.Background(), "SELECT phone_number, password, username FROM public.user")
// 	// jolly.PanicIF(err)
// 	//
// 	// var userModels []User
// 	// for _, row := range rows {
// 	// 	userModel := NewUserWithOutPtr()
// 	// 	err = userModel.Scan(row)
// 	// 	jolly.PanicIF(err)
// 	//
// 	// 	userModels = append(userModels, userModel)
// 	// }
// 	// fmt.Println(userModels)
// 	//
// 	// row, err := rdbms.QueryOne(context.Background(), "SELECT phone_number, password, username FROM public.user WHERE username = 'tidak ada'")
// 	// jolly.PanicIF(err)
// 	//
// 	// userModel := NewUserWithOutPtr()
// 	// err = userModel.Scan(row)
// 	// jolly.PanicIF(err)
// 	// fmt.Println(userModel)
// 	//
// 	// exist, err := rdbms.CheckOne(context.Background(), "SELECT EXISTS(SELECT 1 FROM public.user)")
// 	// jolly.PanicIF(err)
// 	// fmt.Println(exist)
// }
