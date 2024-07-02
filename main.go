package main

import (
	core "mpb/src"
	"time"
)



func main() {
	chnlMap:= make(map[string]chan string)
	store:=core.GetInstance(chnlMap)
	client:= core.GetClient(store)
	go func(){
		client.SUBSCRIBE("hello")
	}()

	time.Sleep(2*time.Second)
	core.Pub()
}