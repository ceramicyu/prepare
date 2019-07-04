package afu

/*
CREATE TABLE `order_huli_staff_relation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_huli_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单ID',
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目id',
  `user_patient_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户id',
  `department_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '科室id',
  `user_staff_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '员工id',
  `start_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '结束时间',
  `original_price` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '订单原价',
  `day_price` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '每日客户应付费用',
  `daily_wage` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '每日员工应得报酬',
  `is_dissatisfied` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否因为不满意，1 为是，0为不是，默认为0',
  `dissatisfied_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '不满意的原因',
  `has_wage` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有工资，1 为有，-1 为没有',
  `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '状态：1有效，-1无效',
  PRIMARY KEY (`id`),
  KEY `user_staff_id` (`user_staff_id`) USING BTREE,
  KEY `order_info_id` (`order_huli_id`) USING BTREE,
  KEY `start_time` (`start_time`) USING BTREE,
  KEY `end_time` (`end_time`) USING BTREE,
  KEY `project_id` (`project_id`)
) ENGINE=MyISAM AUTO_INCREMENT=112872 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='专护员工与单对应关系';
 */

type OrderHuliStaffRelationModel struct {
	OrderHuliId int32 `json:"order_huli_id" db:"order_huli_id"`
	UserStaffId int32 `json:"user_staff_id" db:"user_staff_id"`

}

