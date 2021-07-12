package Redis_UUID

import (
	"context"
	"github.com/go-redis/redis"
)

var rdb *redis.Client
var ctx = context.Background()

func initClient() (err error) {
	rdb=redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB:0,
	})

	_,err=rdb.Ping(ctx).Result()
	if err!=nil {
		return err
	}
	return nil
}

func GetUUID() (int64,error) {
	var res int64
	err:=initClient()
	if err!=nil {
		return -1,err
	}
	v:=rdb.Incr(ctx,"test-1")
	res,err=v.Result()
	if err!= nil {
		return -1,err
	}
	return res,nil
}


