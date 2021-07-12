package handler

import (
	"RedisUUID/pkg/Redis-UUID"
	"fmt"
	"net/http"
	"strconv"
)

func GetRedisUUID(w http.ResponseWriter,r *http.Request ) {
	uuid,err:=Redis_UUID.GetUUID()
	if err!= nil {
		fmt.Fprintln(w,err.Error())
	}
	fmt.Fprintln(w,strconv.Itoa(int(uuid)))
}
