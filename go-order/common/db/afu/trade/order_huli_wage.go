package trade

import (
	"fmt"
)

/**
CREATE TABLE `order_huli_wage` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目id',
  `department_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '科室id',
  `user_manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '主管id',
  `order_huli_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单id',
  `order_huli_staff_relation_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '专护员工与单对应关系id',
  `user_patient_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户id',
  `user_patient_name` varchar(50) NOT NULL DEFAULT '' COMMENT '客户姓名',
  `user_staff_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '大护床位id',
  `user_staff_name` varchar(50) NOT NULL DEFAULT '' COMMENT '员工姓名',
  `price_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '定价id（sku统计用）',
  `actual_wage` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '每日员工实际应得工资（含服务之星）（order_huli_staff_relation.daily_wage）',
  `wage_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '最终工资(actual_wage_amount - abs(dissatisfied_decrease_staff_amount))',
  `actual_wage_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '合计工资（actual_wage*天数）',
  `dissatisfied_decrease_staff_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '不满意扣员工工资(存负数)',
  `has_wage` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否有工资（如果客户不满意就没有），1有，-1没有',
  `month` int(6) unsigned NOT NULL DEFAULT '0' COMMENT '月份',
  `not_hundred_percent_commission_second` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '非100%提成的时间(秒)',
  `original_price` decimal(6,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '原价',
  `start_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '截至时间',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_huli_id`) USING BTREE,
  KEY `relation_id` (`order_huli_staff_relation_id`) USING BTREE,
  KEY `department_month` (`department_id`,`month`) USING BTREE,
  KEY `project_month` (`project_id`,`month`)
) ENGINE=InnoDB AUTO_INCREMENT=5675442 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='专护员工工资';


*/
type OrderHuliWageModel struct {
	ID            int32 `json:"id" db:"id"`
	ProjectId     int32 `json:"project_id" db:"project_id"`
}

func (afu *AfuBbModel) DeleteOrderHuliWage(projectId int32,month int32) {

	result,err := afu.Db.Exec("DELETE FROM `order_huli_wage` WHERE (`month`= $1  AND  project_id = $2)",month,projectId)
	fmt.Println("DeleteOrderHuliWage 信息", result,err)

}
