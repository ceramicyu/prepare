package trade

import (
	"fmt"
)

/*
CREATE TABLE `staff_account_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目id',
  `staff_account_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '员工账户id',
  `date` date NOT NULL DEFAULT '1000-10-10' COMMENT '费用产生的日期',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `source_type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '来源类型，1 底薪，2 提成，3 报酬，4 奖励，5 罚款，6 借款，7 损坏物品',
  `source_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '来源id',
  `amount` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '金额',
  `params` varchar(2000) NOT NULL DEFAULT '' COMMENT '其它键值，josn格式存储',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_account_source` (`staff_account_id`,`source_type`,`source_id`),
  KEY `idx_date` (`date`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16418562 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='员工工资明细';


*/
type StaffAccountListModel struct {
	ID            int32 `json:"id" db:"id"`
	ProjectId     int32 `json:"project_id" db:"project_id"`
}
func (afu *AfuBbModel) DeleteStaffAccountList(where string) {

	result,err := afu.Db.Exec("DELETE FROM `staff_account_list` WHERE $1",where)
	fmt.Println("DeleteOrderHuliWage 信息", result,err)

}
