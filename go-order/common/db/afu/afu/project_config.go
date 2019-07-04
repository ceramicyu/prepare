package afu


/**
           'project_id' => 'int', // 项目id
           'petty_cash' => 'decimal', // 项目经理可申领的备用金数
           'staff_social_security' => 'decimal', // 员工社保个人部份金额
           'spend_limit' => 'decimal', // 每月开支上限
           'borrow_limit' => 'decimal', // 每月员工借款上限
           'status' => 'tinyint', // 状态：参考DicGlobal
           'business_type' => 'varchar', // 业务类型|1|2|
           'commission_days' => 'int', // 每月指定天数公司需提成 0为每天提成
           'create_time' => 'int', // 创建时间
           'create_manager_id' => 'int', // 创建人
           'edit_time' => 'int', // 修改时间
           'edit_manager_id' => 'int', // 修改人
 */

type ProjectConfigModel struct {
	CommissionDays int `json:"commission_days" db:"commission_days"`
}