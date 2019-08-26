package main

import (
	"bufio"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
var DB *sqlx.DB
func init(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root","123456", "192.168.0.199", 3340, "gj")
	var(err error)
	DB,err=sqlx.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB.Ping()
}

func main() {

    logs:=[]string{}
	ReadLine("./log", func(s string) {
	if s != ""{
		logs=append(logs, s)
	}

	})
	list, err := getDirList("./sql")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("s",isInArray([]string{"s","a"},"s"))
	for _, v := range list {
		fmt.Println(v,isInArray(logs,v))
		if !isInArray(logs,v){
			//TODO 执行SQL
			i:=int(0)
	 for{
	 	i++

		 b,err:= ioutil.ReadFile(v+"/up.sql")
			r,err:=	DB.Exec(string(b))
			  fmt.Println(">>>>>>正在执行 ",v,"up.sql")
			  if err!=nil{
				  fmt.Println(">>>>>>正在执行 ",v,"up.sql",">>>> 错误")
				  fmt.Println(r,err)
				  continue
			  }
				logs=append(logs, v)
				fmt.Println(">>>>>>正在执行 ",v,"up.sql",">>>> 完成",i)
		 if i > 100000{
           break
		 }

	 }

		}

		}
     str:=""
     for _,v:=range logs{
     	str+=v+"\n"
	 }
     ioutil.WriteFile("./log",[]byte(str),0777)

}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}


func isInArray(arr []string, item string)bool{
	for _,v:=range  arr{
		if strings.Trim(v," ")== strings.Trim(item," "){
			return true
		}
	}
	return false
}
func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}
