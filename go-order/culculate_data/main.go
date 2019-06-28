package main

import (
	"fmt"
	"github.com/ceramicyu/prepare/go-order/common/db/bb/bb"
	city2 "github.com/ceramicyu/prepare/go-order/common/db/sso/city"
	"github.com/jmoiron/sqlx"
)

var DbPool=make(map[string]*sqlx.DB)
func init()  {

	if db_bb_bb,err:=sqlx.Open("mysql","root:123456@tcp(192.168.0.199:3320)/bb?charset=utf8&parseTime=True&loc=Local");err!=nil{
		fmt.Println("数据库链接错误",err)
	}else{
		DbPool["db_bb_bb"]=db_bb_bb
	}
	if db_sso_city,err:=sqlx.Open("mysql","root:123456@tcp(192.168.0.199:3330)/city?charset=utf8&parseTime=True&loc=Local");err!=nil{
		fmt.Println("数据库链接错误",err)
	}else{
		DbPool["db_sso_city"]=db_sso_city
	}
}
func main(){

	//获取项目列表 projectList
	//projectList:=GetProjectList()
	//for _,projectInfo :=range projectList{
	//	fmt.Println(projectInfo)
	//}

	b:=bb.BbBbModel{
		Db:DbPool["db_bb_bb"],
	}
   b.GetPriceHomeInfo()


	city:=city2.SsoCityModel{
		Db:DbPool["db_sso_city"],
	}
	city.GetCityInfo()
}
