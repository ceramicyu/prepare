package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmorion/sqlx"
)

type WageFrozenInterface struct {

}

/*
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目id',
  `month` int(6) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `create_manager_id` int(10) unsigned NOT NULL DEFAULT '0',
  `confirm_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '确认时间',
  `confirm_manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '确认人',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1:项目请求冻结中 2:已项目冻结 3:已财务确认',
  `frozen_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '冻结时间',
 */
type WageFrozenModel struct {
	ID int32 `json:"id" db:"id"`
	ProjectId int32 `json:"project_id" db:"project_id"`
	Month int32 `json:"month" db:"month"`
	CreateTime int32 `json:"create_time" db:"create_time"`
	CreateManagerId int32 `json:"create_manager_id" db:"create_manager_id"`
	ConfirmTime int32 `json:"confirm_time" db:"confirm_time"`
	ConfirmManagerId int32 `json:"confirm_manager_id" db:"confirm_manager_id"`
	Status int8 `json:"status" db:"status"`
	FrozenTime int32 `json:"frozen_time" db:"frozen_time"`
}
var db_afu  *sqlx.DB

func init(){
	var err error
	if db_afu,err=sqlx.Open("mysql","root:123456@tcp(192.168.0.199:3340)/afu?charset=utf8&parseTime=True&loc=Local");err!=nil{
		fmt.Println("数据库链接错误",err)
	}

	fmt.Println("数据库链接成功",err)

}
// 最后确认冻结的月份
func(self *WageFrozenInterface) getLastConfirmFrozenMonth(projectId int32)int32{
	wage:=[]WageFrozenModel{}
	var sql=fmt.Sprintf("select * from wage_frozen where project_id=1 and status=3 order by month desc limit 1 ")
	err:=db_afu.Select(&wage,sql)
	fmt.Println("数据库查询错误",err)
	if len(wage)==0{
		projectId=201708
	}
return projectId
}
// 未冻结的月份的第一天
func(self *WageFrozenInterface) getNoConfirmFrozenFirstDate(projectId int) {
		lastFrozenMonth := self.getLastConfirmFrozenMonth(1)
		if lastFrozenMonth==0 {
		   lastFrozenMonth = 201708
		}

		//$data[$projectId] = date('Ym01', strtotime('+1month', strtotime($lastFrozenMonth . '01 00:00:00')))


	return
}