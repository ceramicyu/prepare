package apiserver

import (
	"net"
	"net/http"
	"time"
)

type ApiServe struct {
	HttpServer *http.Server
}
var (
	G_ApiServer *ApiServe
)
func InitApiServer()(err error){

	var (
		mux *http.ServeMux
		listener  net.Listener
		httpServer *http.Server
	)
	//路由
	mux=http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("saaaaaaaa"))
	})

	if listener,err=net.Listen("tcp",":8080");err!=nil{
		return
	}

	httpServer=&http.Server{
		ReadHeaderTimeout:time.Duration(5*time.Second),
		WriteTimeout:time.Duration(5*time.Second),
		Handler:mux,
	}

	G_ApiServer=&ApiServe{
		HttpServer:httpServer,
	}
	go func() {
		defer func() {

		}()
		httpServer.Serve(listener)

	}()
	return
}