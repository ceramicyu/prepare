package main

import (
	"fmt"
	"github.com/ceramicyu/prepare/go-order/common"
	afu2 "github.com/ceramicyu/prepare/go-order/common/db/afu/afu"
	"github.com/ceramicyu/prepare/go-order/common/db/afu/trade"
	"github.com/ceramicyu/prepare/go-order/common/db/bb/bb"
	"github.com/ceramicyu/prepare/go-order/common/db/sso/city"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var DbPool = make(map[string][]*sqlx.DB)

const SOURCE_TYPE_PAY = 3
const STATUS_CONFIRM = 3
const BOOL_YES = 1

type DSN struct {
	Key      string
	User     string
	PassWord string
	Addr     string
	Port     int32
	DbName   string
}

func init() {
	configs := []DSN{
		{
			Key:      "db_bb_bb",
			User:     "root",
			PassWord: "123456",
			Addr:     "192.168.0.199",
			Port:     3320,
			DbName:   "trade",
		},
		{
			Key:      "db_sso_city",
			User:     "root",
			PassWord: "123456",
			Addr:     "192.168.0.199",
			Port:     3330,
			DbName:   "city",
		},
	}
	for _, v := range configs {
		epool := make([]*sqlx.DB, 0)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			v.User, v.PassWord, v.Addr, v.Port, v.DbName)
		for {
			if db, err := sqlx.Open("mysql", dsn); err != nil {
				fmt.Println("数据库链接错误", err)
			} else {
				epool = append(epool, db)
			}
			if len(epool) > 100 {
				break
			}
		}
		DbPool[v.Key] = epool
	}
}

//清除专护数据
func ClearZhWageData(projectId int32, startTime int64) {
	afu := trade.AfuBbModel{
		Db: DbPool["db_afu_trade"][0],
	}

	tm := time.Unix(int64(startTime), 0)
	month := int32(tm.Year()*100 + common.Month(tm.Month()))
	afu.DeleteOrderHuliWage(projectId, month)
	date := fmt.Sprintf("%04d-%02d-%02d", tm.Year(), common.Month(tm.Month()), tm.Day())
	where := fmt.Sprintf(" project_id = %d AND date >=%s AND source_type = %d", projectId, date, SOURCE_TYPE_PAY)
	afu.DeleteStaffAccountList(where)

}

//重新计算客户余额
func updateZhWageData(projectId int32, startTime int32) {
	//重新计算客户余额
	afu := trade.AfuBbModel{
		Db: DbPool["db_afu_trade"][0],
	}

	//重新计算专护订单的提成
	sql := fmt.Sprintf(`UPDATE afu.order_huli AS oh SET oh.staff_total_amount = IFNULL((
SELECT SUM(ohfw.wage_amount) AS total FROM trade.order_huli_wage AS ohfw WHERE ohfw.order_huli_id = oh.id
), 0.00) WHERE   ( oh.end_time = 0 or oh.end_time >= %v ) and oh.project_id = %v ;`, startTime, projectId)
	result, err := afu.Db.Exec(sql)
	fmt.Println("重新计算", result, err)
}

func calculateByProjectId(ProjectId int32) {
	// 未冻结的月份的第一天
	noFrozenFirstDate := 201708
	afu := trade.AfuBbModel{
		Db: DbPool["db_afu_wage"][0],
	}
	type WageFrozen struct {
		Month int
	}
	months := []WageFrozen{}
	err=afu.Db.Select(&months, "SELECT month FROM wage_frozen WHERE status = $1 order by `month` desc ", STATUS_CONFIRM)
	if len(months) > 0 {
		noFrozenFirstDate = months[0].Month
	}
	t, _ := time.Parse("2006-01-02 15:04:05", "2017-04-25 09:14:00")
	startTime := int(t.Unix())
	userStaffs := []afu2.OrderHuliStaffRelationModel{}
	err=afu.Db.Select(&userStaffs, "SELECT order_huli_id from order_huli_staff_relation where project_id = $1 and status = $2 and ( end_time = 0 or end_time >= $3)",
		ProjectId, BOOL_YES, startTime)
	noFrozenFirstDate = noFrozenFirstDate
	for _, userStaff := range userStaffs {
		userStaff.UserStaffId = userStaff.UserStaffId
		projectConfig := []afu2.ProjectConfigModel{}
		err=afu.Db.Select(&projectConfig, "SELECT * FROM project_config where project_id = $1;", ProjectId)
		commissionDays := 0
		if len(projectConfig) > 0 {
			commissionDays = projectConfig[0].CommissionDays
		}
		commissionDays = commissionDays

	}


}

func calculateByStaffId() {

}

//计算专户单员工的工资
func CalcZhWage() {

	//清除专护数据
	ClearZhWageData(1, 0)

	//计算专户单员工的工资
	/*  OrderHuliWage::calculateByProjectId($projectInfo['id'], false);

	    updateZhWageData($projectInfo['id'], $startTime);
	*/
	updateZhWageData(1, 0)
}
func main() {
	tm := time.Unix(int64(1499683613), 0)
	date := fmt.Sprintf("%04d-%02d-%02d", tm.Year(), common.Month(tm.Month()), tm.Day())

	fmt.Println(date)
	TestDb()
}

func TestDb() {
	b := bb.BbBbModel{
		Db: DbPool["db_bb_bb"][2],
	}
	b.GetPriceHomeInfo()

	c := city.NewSsoCityModel(DbPool["db_sso_city"][5])

    fmt.Println("+++++++++",c.City)
	c.City.GetCityInfo()
}
