// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trade

const VersionNo = "20140725"

/* 交易订单的帐务信息详情 */
type TradeAmount struct {
	AlipayNo            string `json:"alipay_no"`
	BuyerCodFee         string `json:"buyer_cod_fee"`
	BuyerObtainPointFee int    `json:"buyer_obtain_point_fee"`
	CodFee              string `json:"cod_fee"`
	CommissionFee       string `json:"commission_fee"`
	Created             string `json:"created"`
	EndTime             string `json:"end_time"`
	ExpressAgencyFee    string `json:"express_agency_fee"`
	OrderAmounts        struct {
		OrderAmount []*OrderAmount `json:"order_amount"`
	} `json:"order_amounts"`
	PayTime          string `json:"pay_time"`
	Payment          string `json:"payment"`
	PostFee          string `json:"post_fee"`
	PromotionDetails struct {
		PromotionDetail []*PromotionDetail `json:"promotion_detail"`
	} `json:"promotion_details"`
	SellerCodFee string `json:"seller_cod_fee"`
	Tid          int    `json:"tid"`
	TotalFee     string `json:"total_fee"`
}

/* 子订单的帐务数据结构 */
type OrderAmount struct {
	AdjustFee         string `json:"adjust_fee"`
	DiscountFee       string `json:"discount_fee"`
	DivideOrderFee    string `json:"divide_order_fee"`
	Num               int    `json:"num"`
	NumIid            int    `json:"num_iid"`
	Oid               int    `json:"oid"`
	PartMjzDiscount   string `json:"part_mjz_discount"`
	Payment           string `json:"payment"`
	Price             string `json:"price"`
	PromotionName     string `json:"promotion_name"`
	SkuId             int    `json:"sku_id"`
	SkuPropertiesName string `json:"sku_properties_name"`
	Title             string `json:"title"`
}

/* 交易的优惠信息详情 */
type PromotionDetail struct {
	DiscountFee   float64 `json:"discount_fee,string"`
	GiftItemId    string  `json:"gift_item_id"`
	GiftItemName  string  `json:"gift_item_name"`
	GiftItemNum   string  `json:"gift_item_num"`
	Id            int     `json:"id"`
	PromotionDesc string  `json:"promotion_desc"`
	PromotionId   string  `json:"promotion_id"`
	PromotionName string  `json:"promotion_name"`
}

/* 交易结构 */
type Trade struct {
	AdjustFee           string `json:"adjust_fee"`
	AlipayId            int    `json:"alipay_id"`
	AlipayNo            string `json:"alipay_no"`
	AlipayPoint         int    `json:"alipay_point"`
	AlipayUrl           string `json:"alipay_url"`
	AlipayWarnMsg       string `json:"alipay_warn_msg"`
	AreaId              string `json:"area_id"`
	ArriveCutTime       string `json:"arrive_cut_time"`
	ArriveInterval      int    `json:"arrive_interval"`
	AsyncModified       string `json:"async_modified"`
	AvailableConfirmFee string `json:"available_confirm_fee"`
	BuyerAlipayNo       string `json:"buyer_alipay_no"`
	BuyerArea           string `json:"buyer_area"`
	BuyerCodFee         string `json:"buyer_cod_fee"`
	BuyerEmail          string `json:"buyer_email"`
	BuyerFlag           int    `json:"buyer_flag"`
	BuyerMemo           string `json:"buyer_memo"`
	BuyerMessage        string `json:"buyer_message"`
	BuyerNick           string `json:"buyer_nick"`
	BuyerObtainPointFee int    `json:"buyer_obtain_point_fee"`
	BuyerRate           bool   `json:"buyer_rate"`
	CanRate             bool   `json:"can_rate"`
	CodFee              string `json:"cod_fee"`
	CodStatus           string `json:"cod_status"`
	CommissionFee       string `json:"commission_fee"`
	ConsignInterval     int    `json:"consign_interval"`
	ConsignTime         string `json:"consign_time"`
	Created             string `json:"created"`
	CreditCardFee       string `json:"credit_card_fee"`
	DiscountFee         string `json:"discount_fee"`
	EndTime             string `json:"end_time"`
	EticketExt          string `json:"eticket_ext"`
	ExpressAgencyFee    string `json:"express_agency_fee"`
	HasBuyerMessage     bool   `json:"has_buyer_message"`
	HasPostFee          bool   `json:"has_post_fee"`
	HasYfx              bool   `json:"has_yfx"`
	Iid                 string `json:"iid"`
	InvoiceName         string `json:"invoice_name"`
	InvoiceType         string `json:"invoice_type"`
	Is3D                bool   `json:"is_3D"`
	IsBrandSale         bool   `json:"is_brand_sale"`
	IsDaixiao           bool   `json:"is_daixiao"`
	IsForceWlb          bool   `json:"is_force_wlb"`
	IsLgtype            bool   `json:"is_lgtype"`
	IsPartConsign       bool   `json:"is_part_consign"`
	IsWt                bool   `json:"is_wt"`
	LgAging             string `json:"lg_aging"`
	LgAgingType         string `json:"lg_aging_type"`
	MarkDesc            string `json:"mark_desc"`
	Modified            string `json:"modified"`
	Num                 int    `json:"num"`
	NumIid              int    `json:"num_iid"`
	NutFeature          string `json:"nut_feature"`
	O2o                 string `json:"o2o"`
	O2oDelivery         string `json:"o2o_delivery"`
	O2oGuideId          string `json:"o2o_guide_id"`
	O2oGuideName        string `json:"o2o_guide_name"`
	O2oShopId           string `json:"o2o_shop_id"`
	O2oShopName         string `json:"o2o_shop_name"`
	Orders              struct {
		Order []*Order `json:"order"`
	} `json:"orders"`
	PayTime          string `json:"pay_time"`
	Payment          string `json:"payment"`
	PccAf            int    `json:"pcc_af"`
	PicPath          string `json:"pic_path"`
	PointFee         int    `json:"point_fee"`
	PostFee          string `json:"post_fee"`
	Price            string `json:"price"`
	Promotion        string `json:"promotion"`
	PromotionDetails struct {
		PromotionDetail []*PromotionDetail `json:"promotion_detail"`
	} `json:"promotion_details"`
	RealPointFee     int    `json:"real_point_fee"`
	ReceivedPayment  string `json:"received_payment"`
	ReceiverAddress  string `json:"receiver_address"`
	ReceiverCity     string `json:"receiver_city"`
	ReceiverDistrict string `json:"receiver_district"`
	ReceiverMobile   string `json:"receiver_mobile"`
	ReceiverName     string `json:"receiver_name"`
	ReceiverPhone    string `json:"receiver_phone"`
	ReceiverState    string `json:"receiver_state"`
	ReceiverZip      string `json:"receiver_zip"`
	SellerAlipayNo   string `json:"seller_alipay_no"`
	SellerCanRate    bool   `json:"seller_can_rate"`
	SellerCodFee     string `json:"seller_cod_fee"`
	SellerEmail      string `json:"seller_email"`
	SellerFlag       int    `json:"seller_flag"`
	SellerMemo       string `json:"seller_memo"`
	SellerMobile     string `json:"seller_mobile"`
	SellerName       string `json:"seller_name"`
	SellerNick       string `json:"seller_nick"`
	SellerPhone      string `json:"seller_phone"`
	SellerRate       bool   `json:"seller_rate"`
	SendTime         string `json:"send_time"`
	ServiceOrders    struct {
		ServiceOrder []*ServiceOrder `json:"service_order"`
	} `json:"service_orders"`
	ServiceTags struct {
		LogisticsTag []*LogisticsTag `json:"logistics_tag"`
	} `json:"service_tags"`
	ShippingType      string `json:"shipping_type"`
	Snapshot          string `json:"snapshot"`
	SnapshotUrl       string `json:"snapshot_url"`
	Status            string `json:"status"`
	StepPaidFee       string `json:"step_paid_fee"`
	StepTradeStatus   string `json:"step_trade_status"`
	Tid               int    `json:"tid"`
	TimeoutActionTime string `json:"timeout_action_time"`
	Title             string `json:"title"`
	TotalFee          string `json:"total_fee"`
	TradeFrom         string `json:"trade_from"`
	TradeMemo         string `json:"trade_memo"`
	TradeSource       string `json:"trade_source"`
	Type              string `json:"type"`
	YfxFee            string `json:"yfx_fee"`
	YfxId             string `json:"yfx_id"`
	YfxType           string `json:"yfx_type"`
	ZeroPurchase      bool   `json:"zero_purchase"`
}

/* 订单结构 */
type Order struct {
	AdjustFee         string `json:"adjust_fee"`
	BindOid           int    `json:"bind_oid"`
	BuyerNick         string `json:"buyer_nick"`
	BuyerRate         bool   `json:"buyer_rate"`
	Cid               int    `json:"cid"`
	ConsignTime       string `json:"consign_time"`
	DiscountFee       string `json:"discount_fee"`
	DivideOrderFee    string `json:"divide_order_fee"`
	EndTime           string `json:"end_time"`
	Iid               string `json:"iid"`
	InvoiceNo         string `json:"invoice_no"`
	IsDaixiao         bool   `json:"is_daixiao"`
	IsOversold        bool   `json:"is_oversold"`
	IsServiceOrder    bool   `json:"is_service_order"`
	IsWww             bool   `json:"is_www"`
	ItemMealId        int    `json:"item_meal_id"`
	ItemMealName      string `json:"item_meal_name"`
	LogisticsCompany  string `json:"logistics_company"`
	Modified          string `json:"modified"`
	Num               int    `json:"num"`
	NumIid            int    `json:"num_iid"`
	Oid               int    `json:"oid"`
	OrderFrom         string `json:"order_from"`
	OuterIid          string `json:"outer_iid"`
	OuterSkuId        string `json:"outer_sku_id"`
	PartMjzDiscount   string `json:"part_mjz_discount"`
	Payment           string `json:"payment"`
	PicPath           string `json:"pic_path"`
	Price             string `json:"price"`
	RefundId          int    `json:"refund_id"`
	RefundStatus      string `json:"refund_status"`
	SellerNick        string `json:"seller_nick"`
	SellerRate        bool   `json:"seller_rate"`
	SellerType        string `json:"seller_type"`
	ShippingType      string `json:"shipping_type"`
	SkuId             string `json:"sku_id"`
	SkuPropertiesName string `json:"sku_properties_name"`
	Snapshot          string `json:"snapshot"`
	SnapshotUrl       string `json:"snapshot_url"`
	Status            string `json:"status"`
	StoreCode         string `json:"store_code"`
	TicketExpdateKey  string `json:"ticket_expdate_key"`
	TicketOuterId     string `json:"ticket_outer_id"`
	TimeoutActionTime string `json:"timeout_action_time"`
	Title             string `json:"title"`
	TotalFee          string `json:"total_fee"`
}

/* 商城虚拟服务子订单数据结构 */
type ServiceOrder struct {
	BuyerNick        string `json:"buyer_nick"`
	ItemOid          int    `json:"item_oid"`
	Num              int    `json:"num"`
	Oid              int    `json:"oid"`
	Payment          string `json:"payment"`
	PicPath          string `json:"pic_path"`
	Price            string `json:"price"`
	RefundId         int    `json:"refund_id"`
	SellerNick       string `json:"seller_nick"`
	ServiceDetailUrl string `json:"service_detail_url"`
	ServiceId        int    `json:"service_id"`
	Title            string `json:"title"`
	TotalFee         string `json:"total_fee"`
}

/* 物流的服务标签 */
type LogisticsTag struct {
	LogisticServiceTagList struct {
		LogisticServiceTag []*LogisticServiceTag `json:"logistic_service_tag"`
	} `json:"logistic_service_tag_list"`
	OrderId string `json:"order_id"`
}

/* 物流服务标签 */
type LogisticServiceTag struct {
	ServiceTag  string `json:"service_tag"`
	ServiceType string `json:"service_type"`
}

/* 确认收货费用结构 */
type TradeConfirmFee struct {
	ConfirmFee     float64 `json:"confirm_fee,string"`
	ConfirmPostFee float64 `json:"confirm_post_fee,string"`
	IsLastOrder    bool    `json:"is_last_order"`
}

/* 外卖订单查询结果 */
type ListResult struct {
	ResultList struct {
		TakeoutThirdOrder []*TakeoutThirdOrder `json:"takeout_third_order"`
	} `json:"result_list"`
	ResultSize int    `json:"result_size"`
	RetCode    string `json:"ret_code"`
	TotalCount int    `json:"total_count"`
}

/* 外卖订单对象 */
type TakeoutThirdOrder struct {
	Address         string  `json:"address"`
	AlipayStreamId  string  `json:"alipay_stream_id"`
	CreateTime      string  `json:"create_time"`
	DeliveryPay     float64 `json:"delivery_pay,string"`
	EndDeliverytime string  `json:"end_deliverytime"`
	From            int     `json:"from"`
	GoodsList       struct {
		OrderGoods []*OrderGoods `json:"order_goods"`
	} `json:"goods_list"`
	Id                int               `json:"id"`
	Note              string            `json:"note"`
	OrderStatus       int               `json:"order_status"`
	PayAccount        string            `json:"pay_account"`
	PeopleNumber      int               `json:"people_number"`
	StartDeliverytime string            `json:"start_deliverytime"`
	StoreContactphone string            `json:"store_contactphone"`
	StoreId           int               `json:"store_id"`
	StoreName         string            `json:"store_name"`
	TotalPay          float64           `json:"total_pay,string"`
	UserAddress       *UserThirdAddress `json:"user_address"`
	UserId            int               `json:"user_id"`
}

/* 物流订单商品 */
type OrderGoods struct {
	AuctionCode   string `json:"auction_code"`
	Count         int    `json:"count"`
	GoodsCode     string `json:"goods_code"`
	GoodsName     string `json:"goods_name"`
	GoodsPicId    string `json:"goods_pic_id"`
	GoodsQuantity int    `json:"goods_quantity"`
	ItemValue     int    `json:"item_value"`
	Name          string `json:"name"`
	OrderGoodsId  int    `json:"order_goods_id"`
	RealPrice     string `json:"real_price"`
	SellProperty  string `json:"sell_property"`
	UserId        int    `json:"user_id"`
}

/* 送货地址对象 */
type UserThirdAddress struct {
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
}
