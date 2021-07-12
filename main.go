package main

import (
	"RedisUUID/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/getRedisUUID",handler.GetRedisUUID)
	_=http.ListenAndServe(":8081",nil)

}


