// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crm

const VersionNo = "20140725"

/* 卖家设置的等级优惠信息 */
type GradePromotion struct {
	CurGrade          string `json:"cur_grade"`
	CurGradeName      string `json:"cur_grade_name"`
	Discount          int    `json:"discount"`
	NextGrade         string `json:"next_grade"`
	NextGradeName     string `json:"next_grade_name"`
	NextUpgradeAmount int    `json:"next_upgrade_amount"`
	NextUpgradeCount  int    `json:"next_upgrade_count"`
}

/* 描述分组的数据结构 */
type Group struct {
	GroupCreate string `json:"group_create"`
	GroupId     int    `json:"group_id"`
	GroupModify string `json:"group_modify"`
	GroupName   string `json:"group_name"`
	MemberCount int    `json:"member_count"`
	Status      string `json:"status"`
}

/* 表示会员关系的基本信息字段，用于会员列表的基本查询 */
type BasicMember struct {
	BizOrderId       int     `json:"biz_order_id"`
	BuyerId          int     `json:"buyer_id"`
	BuyerNick        string  `json:"buyer_nick"`
	CloseTradeAmount float64 `json:"close_trade_amount,string"`
	CloseTradeCount  int     `json:"close_trade_count"`
	Grade            int     `json:"grade"`
	GroupIds         string  `json:"group_ids"`
	ItemNum          int     `json:"item_num"`
	LastTradeTime    string  `json:"last_trade_time"`
	RelationSource   int     `json:"relation_source"`
	Status           string  `json:"status"`
	TradeAmount      float64 `json:"trade_amount,string"`
	TradeCount       int     `json:"trade_count"`
}

/* 会员信息对象 */
type CrmMember struct {
	AvgPrice         float64 `json:"avg_price,string"`
	BuyerId          int     `json:"buyer_id"`
	BuyerNick        string  `json:"buyer_nick"`
	City             string  `json:"city"`
	CloseTradeAmount float64 `json:"close_trade_amount,string"`
	CloseTradeCount  int     `json:"close_trade_count"`
	Grade            int     `json:"grade"`
	GroupIds         string  `json:"group_ids"`
	ItemCloseCount   int     `json:"item_close_count"`
	ItemNum          int     `json:"item_num"`
	LastTradeTime    string  `json:"last_trade_time"`
	Province         int     `json:"province"`
	RelationSource   int     `json:"relation_source"`
	Status           string  `json:"status"`
	TradeAmount      float64 `json:"trade_amount,string"`
	TradeCount       int     `json:"trade_count"`
}

/* tmall权益 */
type GradeEquity struct {
	CurGrade     string `json:"cur_grade"`
	CurGradeName string `json:"cur_grade_name"`
	ExcludeArea  string `json:"exclude_area"`
	Point        int    `json:"point"`
	Postage      bool   `json:"postage"`
}
