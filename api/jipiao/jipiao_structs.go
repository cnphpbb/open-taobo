// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jipiao

const VersionNo = "20140725"

/* 机票订单的详情 */
type TripOrder struct {
	BaseInfo    *TripBaseInfo `json:"base_info"`
	CorpInfo    *CorpInfo     `json:"corp_info"`
	Extra       string        `json:"extra"`
	FlightInfos struct {
		TripFlightInfo []*TripFlightInfo `json:"trip_flight_info"`
	} `json:"flight_infos"`
	Itinerary *Itinerary `json:"itinerary"`
}

/* 订单基础信息 */
type TripBaseInfo struct {
	AlipayTradeNo    string `json:"alipay_trade_no"`
	Commission       string `json:"commission"`
	CreateTime       string `json:"create_time"`
	Extra            string `json:"extra"`
	ForceInsure      bool   `json:"force_insure"`
	InsurePromotion  bool   `json:"insure_promotion"`
	ModifyTime       string `json:"modify_time"`
	OrderId          int    `json:"order_id"`
	PayLatestTime    string `json:"pay_latest_time"`
	PayStatus        int    `json:"pay_status"`
	RelationEmail    string `json:"relation_email"`
	RelationMobile   string `json:"relation_mobile"`
	RelationName     string `json:"relation_name"`
	RelationPhoneBak string `json:"relation_phone_bak"`
	Status           int    `json:"status"`
	TotalPrice       int    `json:"total_price"`
	TripType         int    `json:"trip_type"`
}

/* 国内机票订单行程购票数据结构录入【top订单优化】 */
type CorpInfo struct {
	ApplyName      string `json:"apply_name"`
	ApplyNo        string `json:"apply_no"`
	ApplyTime      string `json:"apply_time"`
	CorprationId   string `json:"corpration_id"`
	CostCenter     string `json:"cost_center"`
	CostCenterCode string `json:"cost_center_code"`
	CostOu         string `json:"cost_ou"`
	Extra          string `json:"extra"`
	FormNo         string `json:"form_no"`
	FormStatus     string `json:"form_status"`
	ReceiptsStatus string `json:"receipts_status"`
	TripPersonName string `json:"trip_person_name"`
	TripPersonNo   string `json:"trip_person_no"`
	WorkSpace      string `json:"work_space"`
}

/* 订单航班信息（包含航班乘机人） */
type TripFlightInfo struct {
	AirlineCode    string `json:"airline_code"`
	ArrAirportCode string `json:"arr_airport_code"`
	ArrCityCode    string `json:"arr_city_code"`
	ArrTime        string `json:"arr_time"`
	Carrier        string `json:"carrier"`
	DepAirportCode string `json:"dep_airport_code"`
	DepCityCode    string `json:"dep_city_code"`
	DepTime        string `json:"dep_time"`
	Extra          string `json:"extra"`
	FlightId       int    `json:"flight_id"`
	FlightNo       string `json:"flight_no"`
	FlightType     string `json:"flight_type"`
	Passengers     struct {
		TripFlightPassenger []*TripFlightPassenger `json:"trip_flight_passenger"`
	} `json:"passengers"`
	SegmentNumber int    `json:"segment_number"`
	SegmentType   int    `json:"segment_type"`
	SpecialRule   string `json:"special_rule"`
	TicketPrice   int    `json:"ticket_price"`
}

/* 乘机人信息 */
type TripFlightPassenger struct {
	Birthday             string `json:"birthday"`
	CabinClass           int    `json:"cabin_class"`
	CabinCode            string `json:"cabin_code"`
	CertNo               string `json:"cert_no"`
	CertType             int    `json:"cert_type"`
	Ei                   string `json:"ei"`
	Extra                string `json:"extra"`
	Fee                  int    `json:"fee"`
	ForceInsurePrice     int    `json:"force_insure_price"`
	InsurePromotionPrice int    `json:"insure_promotion_price"`
	Memo                 string `json:"memo"`
	Name                 string `json:"name"`
	PassengerType        int    `json:"passenger_type"`
	Pnr                  string `json:"pnr"`
	PolicyId             int    `json:"policy_id"`
	PolicyType           int    `json:"policy_type"`
	Price                int    `json:"price"`
	Tax                  int    `json:"tax"`
	TicketNo             string `json:"ticket_no"`
	TripCardNo           string `json:"trip_card_no"`
	Tuigaiqian           string `json:"tuigaiqian"`
}

/* 国内机票行程单数据结构定义【top订单优化】 */
type Itinerary struct {
	Address       string `json:"address"`
	AlipayTradeNo string `json:"alipay_trade_no"`
	CompanyCode   string `json:"company_code"`
	ExpressCode   string `json:"express_code"`
	Extra         string `json:"extra"`
	Id            int    `json:"id"`
	ItineraryNo   string `json:"itinerary_no"`
	Mobile        string `json:"mobile"`
	MobileBak     string `json:"mobile_bak"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	SendDate      string `json:"send_date"`
	Status        int    `json:"status"`
	Type          int    `json:"type"`
}

/* 代理商订单搜索接口返回数据对象【订单优化】 */
type SearchOrderResult struct {
	HasNext  bool  `json:"has_next"`
	OrderIds []int `json:"order_ids"`
	PageSize int   `json:"page_size"`
}

/* 国内机票订单数据结构【top订单优化】 */
type AtOrder struct {
	BaseInfo     *BaseInfo  `json:"base_info"`
	CorpInfo     *CorpInfo  `json:"corp_info"`
	Extra        string     `json:"extra"`
	Itinerary    *Itinerary `json:"itinerary"`
	SegmentInfos struct {
		SegmentInfo []*SegmentInfo `json:"segment_info"`
	} `json:"segment_infos"`
}

/* 国内机票，订单基本信息数据结构【top订单优化】 */
type BaseInfo struct {
	AccountNo          string `json:"account_no"`
	AlipayTradeNo      string `json:"alipay_trade_no"`
	BookWay            int    `json:"book_way"`
	Commission         string `json:"commission"`
	CommissionDiscount string `json:"commission_discount"`
	CreateTime         string `json:"create_time"`
	Extra              string `json:"extra"`
	ForceInsure        bool   `json:"force_insure"`
	InsurePromotion    bool   `json:"insure_promotion"`
	ModifyTime         string `json:"modify_time"`
	OrderId            int    `json:"order_id"`
	PayLatestTime      string `json:"pay_latest_time"`
	PayStatus          int    `json:"pay_status"`
	RelationEmail      string `json:"relation_email"`
	RelationMobile     string `json:"relation_mobile"`
	RelationName       string `json:"relation_name"`
	RelationPhoneBak   string `json:"relation_phone_bak"`
	RelativeOrderId    int    `json:"relative_order_id"`
	Status             int    `json:"status"`
	TotalPrice         int    `json:"total_price"`
	TripType           int    `json:"trip_type"`
}

/* 国内机票航段信息数据结构，【订单优化】 */
type SegmentInfo struct {
	AirlineCode               string `json:"airline_code"`
	ArrAirportCode            string `json:"arr_airport_code"`
	ArrCityCode               string `json:"arr_city_code"`
	ArrTime                   string `json:"arr_time"`
	BookStatus                int    `json:"book_status"`
	CabinClass                int    `json:"cabin_class"`
	CabinCode                 string `json:"cabin_code"`
	CabinId                   int    `json:"cabin_id"`
	Carrier                   string `json:"carrier"`
	ChildFee                  int    `json:"child_fee"`
	ChildInsurePromotionPrice int    `json:"child_insure_promotion_price"`
	ChildPrice                int    `json:"child_price"`
	ChildTax                  int    `json:"child_tax"`
	DepAirportCode            string `json:"dep_airport_code"`
	DepCityCode               string `json:"dep_city_code"`
	DepTime                   string `json:"dep_time"`
	Extra                     string `json:"extra"`
	Fee                       int    `json:"fee"`
	FlightId                  int    `json:"flight_id"`
	FlightNo                  string `json:"flight_no"`
	FlightType                string `json:"flight_type"`
	InsurePromotionPrice      int    `json:"insure_promotion_price"`
	Memo                      string `json:"memo"`
	Passengers                struct {
		Passerger []*Passerger `json:"passerger"`
	} `json:"passengers"`
	PolicyId    int    `json:"policy_id"`
	PolicyType  int    `json:"policy_type"`
	Price       int    `json:"price"`
	SegmentType int    `json:"segment_type"`
	SpecialRule string `json:"special_rule"`
	Tax         int    `json:"tax"`
	TicketPrice int    `json:"ticket_price"`
}

/* 国内机票乘机人信息数据结构【top订单优化】 */
type Passerger struct {
	Birthday             string `json:"birthday"`
	CertNo               string `json:"cert_no"`
	CertType             int    `json:"cert_type"`
	Ei                   string `json:"ei"`
	Extra                string `json:"extra"`
	ForceInsurePrice     int    `json:"force_insure_price"`
	InsurePromotionPrice int    `json:"insure_promotion_price"`
	Name                 string `json:"name"`
	PassengerType        int    `json:"passenger_type"`
	Pnr                  string `json:"pnr"`
	TicketNo             string `json:"ticket_no"`
	TripCardNo           string `json:"trip_card_no"`
	Tuigaiqian           string `json:"tuigaiqian"`
}
