package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmorion/sqlx"
)
var db_sso *sqlx.DB
/**
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '项目名称',
  `hospital_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '医院id',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市id',
  `company_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '公司id',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT 'DicProject::$TYPE_OPTIONS',
  `note` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `start_time` int(10) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` int(10) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `create_manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `edit_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `edit_manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '状态：1有效，-1无效',
  `is_third` tinyint(3) NOT NULL DEFAULT '-1' COMMENT '是否第三方经营 -1:否 1:是',
 */
type ProjectList struct {
	ID int32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	HospitalId int32 `json:"hospital_id" db:"hospital_id"`
	CityId int32 `json:"city_id" db:"city_id"`
	CompanyId int32 `json:"company_id" db:"company_id"`
	Type int8 `json:"type" db:"type"`
	Note string `json:"note" db:"note"`
	StartTime int32 `json:"start_time" db:"start_time"`
	EndTime int32 `json:"end_time" db:"end_time"`
	CreateTime int32 `json:"create_time" db:"create_time"`
	CreateManagerId int32 `json:"create_manager_id" db:"create_manager_id"`
	EditTime int32 `json:"edit_time" db:"edit_time"`
	EditManagerId int32 `json:"edit_manager_id" db:"edit_manager_id"`
	Status int8 `json:"status" db:"status"`
	IsThird int8 `json:"is_third" db:"is_third"`
}
func init()  {
	var err error
	if db_sso,err=sqlx.Open("mysql","root:123456@tcp(192.168.0.199:3330)/oa?charset=utf8&parseTime=True&loc=Local");err!=nil{
		fmt.Println("数据库链接错误",err)
	}
}
func GetProjectList()([]ProjectList){


	proj:=[]ProjectList{}

  err:= db_sso.Select(&proj,"select * from project limit 10")

 fmt.Println(err,proj)
  return proj
}