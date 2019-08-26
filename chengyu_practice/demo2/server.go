package main

import (
	"encoding/json"
	"fmt"

	"io"
	"io/ioutil"
	"net/http"
)

/*
docker build -t server9000 .

docker run --rm -it -d -p 3000:3000  -v /data:/data  server
 */
func main(){
	Addr:=":9090"
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		resp,err:=http.Get("server.com")
		if err!=nil{
			io.WriteString(writer,fmt.Sprintf("%v",err))
			return
		}
		b,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			io.WriteString(writer,fmt.Sprintf("%v",err))
			return
		}
		io.WriteString(writer,string(b))
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(">>>>>>")
		type Output struct {
			Code int `json:"code"`
			Msg string `json:"msg"`
			Data interface{} `json:"data"`
		}
		obj:=Output{
			Code:102,
			Msg:"成功",
			Data:[]string{"sagsh","dagd","dgahdh"},
		}
		output,_:=json.Marshal(obj)
		io.WriteString(writer,string(output))
	})
	fmt.Println(">>>>>>侦听",Addr)
	err:=http.ListenAndServe(Addr,nil)
	fmt.Println(err)
}



