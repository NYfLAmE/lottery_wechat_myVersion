package model

// Prize 奖品表
type Prize struct {
	ID             int64  `json:"id" gorm:"id"`
	Name           string `json:"name" gorm:"name"`                       // 奖品名称
	Pic            string `json:"pic" gorm:"pic"`                         // 奖品图片
	Link           string `json:"link" gorm:"link"`                       // 奖品链接
	Type           int64  `json:"type" gorm:"type"`                       // 奖品类型，1-虚拟币，2-虚拟券，3-实物小奖，4-实物大奖
	Data           string `json:"data" gorm:"data"`                       // 奖品数据
	Total          int64  `json:"total" gorm:"total"`                     // 奖品数量，0 无限量，>0限量，<0无奖品
	Left           int64  `json:"left" gorm:"left"`                       // 剩余数量
	IsUse          int64  `json:"is_use" gorm:"is_use"`                   // 是否使用中，1-使用中，2-未使用
	Probability    int64  `json:"probability" gorm:"probability"`         // 中奖概率，万分之n
	ProbabilityMax int64  `json:"probability_max" gorm:"probability_max"` // 中奖概率上限
	ProbabilityMin int64  `json:"probability_min" gorm:"probability_min"` // 中奖概率下限
}

type PrizeReq struct { // http请求中的奖品 对应的结构体
	Name  string `json:"name" gorm:"name"`   // 奖品名称
	Type  int64  `json:"type" gorm:"type"`   // 奖品类型，1-虚拟币，2-虚拟券，3-实物小奖，4-实物大奖
	Total int64  `json:"total" gorm:"total"` // 奖品数量，0 无限量，>0限量，<0无奖品
}

type AddPrizeReq struct { // http请求发来的奖品列表 对应的结构体
	PrizeList []*PrizeReq `json:"prize_list"`
}

type GetPrizeRsp struct {
	TypeNum int64 `json:"prize_type_num"`
	Total   int64 `json:"prize_total"`
}

func (p *Prize) TableName() string {
	return "prize" // 返回表名
}
