package main

import (
	"github.com/ceramicyu/prepare/easy_server/apiserver"
	"time"
)

func main(){

     apiserver.InitApiServer()
     for{
     	time.Sleep(time.Second)
	 }
}