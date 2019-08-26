package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)
var(
	c string
	h string
	d string
)
func init() {
	flag.StringVar(&c, "c", "", "命令")
	flag.StringVar(&h, "h", "", "数据库host")
	flag.StringVar(&d, "d", "", "数据库dbName")
}
func main(){

	flag.Parse()

	//执行phinx命令
	Phinx()

	if strings.Index(c,"create")<0{
		//执行Model脚本
		ExecModel()
	}



}
func Phinx(){
	phinxFile := "/data/www/fm/phinx.yml"
	os.MkdirAll("/data/www/fm/db/migrations/"+h+"/"+d, os.ModePerm)
	os.MkdirAll("/data/www/fm/db/seeds/"+h+"/"+d, os.ModePerm)
	fStr:=CreateFile(h,d)
	ioutil.WriteFile(phinxFile,[]byte(fStr),os.ModePerm)
	fmt.Println(fStr)

	s:=fmt.Sprintf(	"docker exec  php-nginx /bin/bash -c \"cd /data/www/fm&& php vendor/robmorgan/phinx/bin/phinx  %s \" ",c)
	cmd := exec.Command(`/bin/bash`,"-c",s)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", s, err.Error())
		fmt.Println(string(output))
		os.Exit(127)
	}
	outStr:=fmt.Sprintf("\033[%s;%dm"+string(output)+"\033[0m", strings.Join([]string{"0"}, ";"), 32)
	fmt.Println("ok\n", outStr)
}

func ExecModel(){
	modelFile := fmt.Sprintf(`/data/www/cron/model/%v.sh`,d)
	simModelFile := fmt.Sprintf(`/data/www/cron/model/sim/%v.sh`,d)


	fmt.Println(">>>执行",d,"model")
	CreateModel(modelFile)
	fmt.Println(">>>执行 sim ",d,"model")
	CreateModel(simModelFile)
}

func CreateModel(modelFile string){
	//执行model

	f,err:=ioutil.ReadFile(modelFile)
	if err != nil {
		fmt.Println("文件读取错误  ",modelFile ,err,string(f))
		return
	}
	cmd := exec.Command(`/bin/bash`,"-c",string(f))
	output, err :=cmd.Output()
	if err!= nil{
		fmt.Printf("Execute Shell:%s failed with error:%s", modelFile, err.Error())
		fmt.Println(string(output))
		return
	}else{
		fmt.Println(string(output),">>>执行成功")
	}
}

func CreateFile(host,dataBase string)string{
	return fmt.Sprintf(
`paths:
    migrations: ./db/migrations/%v/%v
    seeds: ./db/seeds/%v/%v

environments:
    default_migration_table: phinxlog
    default_database: production
    production:
        adapter: mysql
        host: mysql-master-%v
        name: '%v'
        user: root
        pass: '123456'
        port: 3306
        charset: utf8

version_order: creation`,host,dataBase,host,dataBase,host,dataBase)
}