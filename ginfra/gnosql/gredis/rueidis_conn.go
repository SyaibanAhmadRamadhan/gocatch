package gredis

import (
	"github.com/redis/rueidis"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func OpenConnRuedis(opt rueidis.ClientOption) rueidis.Client {
	conn, err := rueidis.NewClient(opt)
	gcommon.PanicIfError(err)

	return conn
}
