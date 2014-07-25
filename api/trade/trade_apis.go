// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trade

import (
	"github.com/cnphpbb/open_taobao"
)

/* 卖家查询该笔交易订单的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的订单，但不同状态时订单的相关数据可能会有不同 */
type TradeAmountGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 订单帐务详情需要返回的字段信息，可选值如下：
1. TradeAmount中可指定的fields：
tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee
2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,
order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name
3. order_amounts(返回OrderAmount的所有内容)
4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段) */
func (r *TradeAmountGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 订单交易编号<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：-9223372036854775808 */
func (r *TradeAmountGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeAmountGetRequest) GetResponse(accessToken string) (*TradeAmountGetResponse, []byte, error) {
	var resp TradeAmountGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.amount.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeAmountGetResponse struct {
	TradeAmount *TradeAmount `json:"trade_amount"`
}

type TradeAmountGetResponseResult struct {
	Response *TradeAmountGetResponse `json:"trade_amount_get_response"`
}

/* 关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。 */
type TradeCloseRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 交易关闭原因。可以选择的理由有：
1.未及时付款
2.买家联系不上
3.谢绝还价
4.商品瑕疵
5.协商不一致
6.买家不想买
7.与买家协商一致 */
func (r *TradeCloseRequest) SetCloseReason(value string) {
	r.SetValue("close_reason", value)
}

/* 主订单或子订单编号。 */
func (r *TradeCloseRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeCloseRequest) GetResponse(accessToken string) (*TradeCloseResponse, []byte, error) {
	var resp TradeCloseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.close", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeCloseResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeCloseResponseResult struct {
	Response *TradeCloseResponse `json:"trade_close_response"`
}

/* 获取交易确认收货费用
可以获取主订单或子订单的确认收货费用 */
type TradeConfirmfeeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 是否是子订单。可选值:IS_FATHER(父订单),IS_CHILD(子订单) */
func (r *TradeConfirmfeeGetRequest) SetIsDetail(value string) {
	r.SetValue("is_detail", value)
}

/* 交易编号，或子订单编号 */
func (r *TradeConfirmfeeGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeConfirmfeeGetRequest) GetResponse(accessToken string) (*TradeConfirmfeeGetResponse, []byte, error) {
	var resp TradeConfirmfeeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.confirmfee.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeConfirmfeeGetResponse struct {
	TradeConfirmFee *TradeConfirmFee `json:"trade_confirm_fee"`
}

type TradeConfirmfeeGetResponseResult struct {
	Response *TradeConfirmfeeGetResponse `json:"trade_confirmfee_get_response"`
}

/* 获取单笔交易的详细信息
<br/>1. 只有在交易成功的状态下才能取到交易佣金，其它状态下取到的都是零或空值
<br/>2. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
<br/>3. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>4. 请按需获取字段，减少TOP系统的压力 */
type TradeFullinfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 1.Trade中可以指定返回的fields：seller_nick, buyer_nick, title, type, created, tid, seller_rate,buyer_flag, buyer_rate, status, payment, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment, commission_fee, buyer_memo, seller_memo, alipay_no,alipay_id,buyer_message, pic_path, num_iid, num, price, buyer_alipay_no, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,seller_flag, seller_alipay_no, seller_mobile, seller_phone, seller_name, seller_email, available_confirm_fee, has_post_fee, timeout_action_time, snapshot_url, cod_fee, cod_status, shipping_type, trade_memo, is_3D,buyer_email,buyer_area, trade_from,is_lgtype,is_force_wlb,is_brand_sale,buyer_cod_fee,discount_fee,seller_cod_fee,express_agency_fee,invoice_name,service_orders,credit_cardfee,step_trade_status,step_paid_fee,mark_desc,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source(注：当该授权用户为卖家时不能查看买家buyer_memo,buyer_flag),eticket_ext,send_time, is_daixiao,is_part_consign, arrive_interval, arrive_cut_time, consign_interval,zero_purchase,alipay_point,pcc_af,2.Order中可以指定返回fields：orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.snapshot_url, orders.timeout_action_time，orders.sku_properties_name, orders.item_meal_name, orders.item_meal_id,orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type, orders.is_oversold,orders.end_time,orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no, orders.is_daixiao
3.fields：orders（返回Order的所有内容）
4.flelds：promotion_details(返回promotion_details所有内容，优惠详情),invoice_name(发票抬头),orders.is_www(子订单是否是www订单,orders.store_code(发货的仓库编码)<br>
5. field:service_tags(返回物流标签) */
func (r *TradeFullinfoGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 交易编号<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：1 */
func (r *TradeFullinfoGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeFullinfoGetRequest) GetResponse(accessToken string) (*TradeFullinfoGetResponse, []byte, error) {
	var resp TradeFullinfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.fullinfo.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeFullinfoGetResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeFullinfoGetResponseResult struct {
	Response *TradeFullinfoGetResponse `json:"trade_fullinfo_get_response"`
}

/* 获取单笔交易的部分信息 */
type TradeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要返回的字段。目前支持有：<br>

1. Trade中可以指定返回的fields:seller_nick, buyer_nick, title, type, created, tid, seller_rate, buyer_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment, commission_fee, buyer_memo, seller_memo, alipay_no, buyer_message, pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type， is_daixiao，consign_interval，arrive_interval，arrive_cut_time <br>
2. Order中可以指定返回fields:orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.outer_sku_id, orders.outer_iid, orders.buyer_rate, orders.seller_rate， orders.is_daixiao <br>
3. fields：orders（返回Order中的所有允许返回的字段）,orders.is_wwww(是否是www订单),orders.store_code(仓库代码）<br>
4. field:service_tags(返回物流标签) */
func (r *TradeGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 交易编号 */
func (r *TradeGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeGetRequest) GetResponse(accessToken string) (*TradeGetResponse, []byte, error) {
	var resp TradeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeGetResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeGetResponseResult struct {
	Response *TradeGetResponse `json:"trade_get_response"`
}

/* 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update */
type TradeMemoAddRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0<br /> 支持最大值为：5<br /> 支持最小值为：0 */
func (r *TradeMemoAddRequest) SetFlag(value string) {
	r.SetValue("flag", value)
}

/* 交易备注。最大长度: 1000个字节 */
func (r *TradeMemoAddRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 交易编号 */
func (r *TradeMemoAddRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeMemoAddRequest) GetResponse(accessToken string) (*TradeMemoAddResponse, []byte, error) {
	var resp TradeMemoAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.memo.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeMemoAddResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeMemoAddResponseResult struct {
	Response *TradeMemoAddResponse `json:"trade_memo_add_response"`
}

/* 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能 */
type TradeMemoUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0<br /> 支持最大值为：5<br /> 支持最小值为：0 */
func (r *TradeMemoUpdateRequest) SetFlag(value string) {
	r.SetValue("flag", value)
}

/* 卖家交易备注。最大长度: 1000个字节 */
func (r *TradeMemoUpdateRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 是否对memo的值置空
若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；
若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值 */
func (r *TradeMemoUpdateRequest) SetReset(value string) {
	r.SetValue("reset", value)
}

/* 交易编号 */
func (r *TradeMemoUpdateRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeMemoUpdateRequest) GetResponse(accessToken string) (*TradeMemoUpdateResponse, []byte, error) {
	var resp TradeMemoUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.memo.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeMemoUpdateResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeMemoUpdateResponseResult struct {
	Response *TradeMemoUpdateResponse `json:"trade_memo_update_response"`
}

/* 只能更新发货前子订单的销售属性
只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。
必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先 */
type TradeOrderskuUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 子订单编号（对于单笔订单的交易可以传交易编号）。 */
func (r *TradeOrderskuUpdateRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。 */
func (r *TradeOrderskuUpdateRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}

/* 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。 */
func (r *TradeOrderskuUpdateRequest) SetSkuProps(value string) {
	r.SetValue("sku_props", value)
}

func (r *TradeOrderskuUpdateRequest) GetResponse(accessToken string) (*TradeOrderskuUpdateResponse, []byte, error) {
	var resp TradeOrderskuUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.ordersku.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeOrderskuUpdateResponse struct {
	Order *Order `json:"order"`
}

type TradeOrderskuUpdateResponseResult struct {
	Response *TradeOrderskuUpdateResponse `json:"trade_ordersku_update_response"`
}

/* 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。 */
type TradePostageUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 邮费价格(邮费单位是元） */
func (r *TradePostageUpdateRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 主订单编号 */
func (r *TradePostageUpdateRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradePostageUpdateRequest) GetResponse(accessToken string) (*TradePostageUpdateResponse, []byte, error) {
	var resp TradePostageUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.postage.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradePostageUpdateResponse struct {
	Trade *Trade `json:"trade"`
}

type TradePostageUpdateResponseResult struct {
	Response *TradePostageUpdateResponse `json:"trade_postage_update_response"`
}

/* 延长交易收货时间 */
type TradeReceivetimeDelayRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 延长收货的天数，可选值为：3, 5, 7, 10。<br /> 支持最大值为：10<br /> 支持最小值为：3 */
func (r *TradeReceivetimeDelayRequest) SetDays(value string) {
	r.SetValue("days", value)
}

/* 主订单号 */
func (r *TradeReceivetimeDelayRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeReceivetimeDelayRequest) GetResponse(accessToken string) (*TradeReceivetimeDelayResponse, []byte, error) {
	var resp TradeReceivetimeDelayResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.receivetime.delay", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeReceivetimeDelayResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeReceivetimeDelayResponseResult struct {
	Response *TradeReceivetimeDelayResponse `json:"trade_receivetime_delay_response"`
}

/* 只能更新一笔交易里面的买家收货地址
只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址
更新后的发货地址可以通过taobao.trade.fullinfo.get查到
参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节） */
type TradeShippingaddressUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 收货地址。最大长度为228个字节。<br /> 支持最大长度为：228<br /> 支持的最大列表长度为：228 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverAddress(value string) {
	r.SetValue("receiver_address", value)
}

/* 城市。最大长度为32个字节。如：杭州<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverCity(value string) {
	r.SetValue("receiver_city", value)
}

/* 区/县。最大长度为32个字节。如：西湖区<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverDistrict(value string) {
	r.SetValue("receiver_district", value)
}

/* 移动电话。最大长度为30个字节。<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverMobile(value string) {
	r.SetValue("receiver_mobile", value)
}

/* 收货人全名。最大长度为50个字节。<br /> 支持最大长度为：50<br /> 支持的最大列表长度为：50 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverName(value string) {
	r.SetValue("receiver_name", value)
}

/* 固定电话。最大长度为30个字节。<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverPhone(value string) {
	r.SetValue("receiver_phone", value)
}

/* 省份。最大长度为32个字节。如：浙江<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverState(value string) {
	r.SetValue("receiver_state", value)
}

/* 邮政编码。必须由6个数字组成。<br /> 支持最大长度为：6<br /> 支持的最大列表长度为：6 */
func (r *TradeShippingaddressUpdateRequest) SetReceiverZip(value string) {
	r.SetValue("receiver_zip", value)
}

/* 交易编号。 */
func (r *TradeShippingaddressUpdateRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeShippingaddressUpdateRequest) GetResponse(accessToken string) (*TradeShippingaddressUpdateResponse, []byte, error) {
	var resp TradeShippingaddressUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.shippingaddress.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeShippingaddressUpdateResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeShippingaddressUpdateResponseResult struct {
	Response *TradeShippingaddressUpdateResponse `json:"trade_shippingaddress_update_response"`
}

/* 交易快照查询
目前只支持类型为“旺店标准版(600)”或“旺店入门版(610)”的交易
对于“旺店标准版”类型的交易，返回的snapshot字段为交易快照编号
对于“旺店入门版”类型的交易，返回的snapshot字段为JSON结构的数据(其中的shopPromotion包含了优惠，积分等信息） */
type TradeSnapshotGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要返回的字段列表。现只支持："snapshot"字段 */
func (r *TradeSnapshotGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 交易编号 */
func (r *TradeSnapshotGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *TradeSnapshotGetRequest) GetResponse(accessToken string) (*TradeSnapshotGetResponse, []byte, error) {
	var resp TradeSnapshotGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.snapshot.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeSnapshotGetResponse struct {
	Trade *Trade `json:"trade"`
}

type TradeSnapshotGetResponseResult struct {
	Response *TradeSnapshotGetResponse `json:"trade_snapshot_get_response"`
}

/* 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
<br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
<br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/> <span style="color:red">注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type）</span> */
type TradesSoldGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家昵称 */
func (r *TradesSoldGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss */
func (r *TradesSoldGetRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
func (r *TradesSoldGetRequest) SetExtType(value string) {
	r.SetValue("ext_type", value)
}

/* 需要返回的字段。目前支持有：<br>
1. Trade中可以指定返回的fields:<br>
seller_nick, buyer_nick, title, type, created,  tid, seller_rate,seller_can_rate, buyer_rate,can_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment,  pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,seller_flag,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,buyer_area,has_buyer_message, credit_card_fee, lg_aging_type, lg_aging, step_trade_status,step_paid_fee,mark_desc,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source,send_time,is_daixiao,is_wt,is_part_consign,zero_purchase
<br>
2. Order中可以指定返回fields：orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type, orders.end_time,orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no,orders.is_daixiao<br>
3. fields：orders（返回2中Order的所有内容）
4.fields:service_orders(返回service_order中所有内容) */
func (r *TradesSoldGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1 */
func (r *TradesSoldGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
func (r *TradesSoldGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>
可选值：
RATE_UNBUYER(买家未评)
RATE_UNSELLER(卖家未评)
RATE_BUYER_UNSELLER(买家已评，卖家未评)
RATE_UNBUYER_SELLER(买家未评，卖家已评)
RATE_BUYER_SELLER(买家已评,卖家已评) */
func (r *TradesSoldGetRequest) SetRateStatus(value string) {
	r.SetValue("rate_status", value)
}

/* 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss */
func (r *TradesSoldGetRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
可选值：
TRADE_NO_CREATE_PAY(没有创建支付宝交易)
WAIT_BUYER_PAY(等待买家付款)
WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)
SELLER_CONSIGNED_PART（卖家部分发货）
WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)
TRADE_BUYER_SIGNED(买家已签收,货到付款专用)
TRADE_FINISHED(交易成功)
TRADE_CLOSED(交易关闭)
TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭)
ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY)
ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO)
WAIT_PRE_AUTH_CONFIRM(余额宝0元购合约中) */
func (r *TradesSoldGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
func (r *TradesSoldGetRequest) SetTag(value string) {
	r.SetValue("tag", value)
}

/* 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade, auto_delivery, ec, cod,step的5种交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>
可选值：
fixed(一口价)
auction(拍卖)
guarantee_trade(一口价、拍卖)
step(分阶段付款，万人团，阶梯团订单）
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲)
cod(货到付款)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易)
instant_trade (即时到账)
b2c_cod(大商家货到付款)
hotel_trade(酒店类型交易)
super_market_trade(商超交易)
super_market_cod_trade(商超货到付款交易)
taohua(淘花网交易类型）
waimai(外卖交易类型）
nopaid（即时到帐/趣味猜交易类型）
step (万人团) eticket(电子凭证)
tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）
注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。
pre_auth_type(预授权0元购) */
func (r *TradesSoldGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。 */
func (r *TradesSoldGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}

func (r *TradesSoldGetRequest) GetResponse(accessToken string) (*TradesSoldGetResponse, []byte, error) {
	var resp TradesSoldGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trades.sold.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradesSoldGetResponse struct {
	HasNext      bool `json:"has_next"`
	TotalResults int  `json:"total_results"`
	Trades       struct {
		Trade []*Trade `json:"trade"`
	} `json:"trades"`
}

type TradesSoldGetResponseResult struct {
	Response *TradesSoldGetResponse `json:"trades_sold_get_response"`
}

/* 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_modified - start_modified <= 1天。
<br/>2. 返回的数据结果是以订单的修改时间倒序排列的，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>4. <span style="color:red">使用<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.0.F9TTxy&id=101744">消息服务</a>监听订单变更事件，可以实时获取订单更新数据。</span>
<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。用taobao.trade.fullinfo.get 查订单fields返回type 很容易的能知道订单的类型（type） */
type TradesSoldIncrementGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
func (r *TradesSoldIncrementGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
func (r *TradesSoldIncrementGetRequest) SetExtType(value string) {
	r.SetValue("ext_type", value)
}

/* 需要返回的字段。目前支持有：
1.Trade中可以指定返回的fields:seller_nick, buyer_nick, title, type, created, tid, seller_rate,seller_can_rate, buyer_rate,can_rate,status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment,pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,has_buyer_message,credit_card_fee,step_trade_status,step_paid_fee,mark_desc,send_time,,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source,seller_flag,is_daixiao,is_part_consign,zero_purchase
2.Order中可以指定返回fields：
orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type，orders.end_time,,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no,orders.is_daixiao
3.fields：orders（返回Order的所有内容）
4.fields:service_orders(返回service_order中所有内容) */
func (r *TradesSoldIncrementGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
func (r *TradesSoldIncrementGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
func (r *TradesSoldIncrementGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
func (r *TradesSoldIncrementGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 可选值 TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款)
SELLER_CONSIGNED_PART（卖家部分发货）
WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO)
WAIT_PRE_AUTH_CONFIRM(余额宝0元购合约中) */
func (r *TradesSoldIncrementGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
func (r *TradesSoldIncrementGetRequest) SetTag(value string) {
	r.SetValue("tag", value)
}

/* 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade, auto_delivery, ec, cod,step的5种交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>
可选值：
fixed(一口价)
auction(拍卖)
step（分阶段付款，万人团，阶梯团订单）
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲) cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易)
instant_trade (即时到账)
b2c_cod(大商家货到付款)
hotel_trade(酒店类型交易)
super_market_trade(商超交易),
super_market_cod_trade(商超货到付款交易)
taohua(桃花网交易类型）
waimai(外卖交易类型）
nopaid（即时到帐/趣味猜交易类型）
 eticket(电子凭证)
tmall_i18n（天猫国际）;nopaid（无付款交易）insurance_plus（保险）finance（基金）
注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。
pre_auth_type(预授权0元购) */
func (r *TradesSoldIncrementGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
func (r *TradesSoldIncrementGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}

func (r *TradesSoldIncrementGetRequest) GetResponse(accessToken string) (*TradesSoldIncrementGetResponse, []byte, error) {
	var resp TradesSoldIncrementGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trades.sold.increment.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradesSoldIncrementGetResponse struct {
	HasNext      bool `json:"has_next"`
	TotalResults int  `json:"total_results"`
	Trades       struct {
		Trade []*Trade `json:"trade"`
	} `json:"trades"`
}

type TradesSoldIncrementGetResponseResult struct {
	Response *TradesSoldIncrementGetResponse `json:"trades_sold_increment_get_response"`
}

/* 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
<br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。
<br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。
<br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>4. 使用主动通知监听订单变更事件，可以实时获取订单更新数据。 */
type TradesSoldIncrementvGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
func (r *TradesSoldIncrementvGetRequest) SetEndCreate(value string) {
	r.SetValue("end_create", value)
}

/* 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
func (r *TradesSoldIncrementvGetRequest) SetExtType(value string) {
	r.SetValue("ext_type", value)
}

/* 需要返回的字段。目前支持有：
1.Trade中可以指定返回的fields:seller_nick, buyer_nick, title, type, created, tid, seller_rate, buyer_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment,pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,has_buyer_message,credit_card_fee,step_trade_status,step_paid_fee,mark_desc，is_daixiao,is_part_consign
2.Order中可以指定返回fields：
orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type，orders.end_time, orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no，orders.is_daixiao
3.fields：orders（返回Order的所有内容）
4.fields:service_orders(返回service_order中所有内容) */
func (r *TradesSoldIncrementvGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
func (r *TradesSoldIncrementvGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
func (r *TradesSoldIncrementvGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
func (r *TradesSoldIncrementvGetRequest) SetStartCreate(value string) {
	r.SetValue("start_create", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 可选值 TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)
SELLER_CONSIGNED_PART（卖家部分发货）
WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) */
func (r *TradesSoldIncrementvGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
func (r *TradesSoldIncrementvGetRequest) SetTag(value string) {
	r.SetValue("tag", value)
}

/* 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade, auto_delivery, ec, cod,step的5种交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>
可选值：
fixed(一口价)
auction(拍卖)
step（分阶段付款，万人团，阶梯团订单）
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲) cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易)
instant_trade (即时到账)
b2c_cod(大商家货到付款)
hotel_trade(酒店类型交易)
super_market_trade(商超交易),
super_market_cod_trade(商超货到付款交易)
taohua(桃花网交易类型）
waimai(外卖交易类型）
nopaid（无付款订单）
eticket(电子凭证)
tmall_i18n（天猫国际）
注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。 */
func (r *TradesSoldIncrementvGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
func (r *TradesSoldIncrementvGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}

func (r *TradesSoldIncrementvGetRequest) GetResponse(accessToken string) (*TradesSoldIncrementvGetResponse, []byte, error) {
	var resp TradesSoldIncrementvGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trades.sold.incrementv.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradesSoldIncrementvGetResponse struct {
	HasNext      bool `json:"has_next"`
	TotalResults int  `json:"total_results"`
	Trades       struct {
		Trade []*Trade `json:"trade"`
	} `json:"trades"`
}

type TradesSoldIncrementvGetResponseResult struct {
	Response *TradesSoldIncrementvGetResponse `json:"trades_sold_incrementv_get_response"`
}

/* 商家确认订单，进行发货 */
type TradeWaimaiConfirmRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 未确认发货的订单编号 */
func (r *TradeWaimaiConfirmRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

func (r *TradeWaimaiConfirmRequest) GetResponse(accessToken string) (*TradeWaimaiConfirmResponse, []byte, error) {
	var resp TradeWaimaiConfirmResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.waimai.confirm", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeWaimaiConfirmResponse struct {
	RetCode string `json:"ret_code"`
}

type TradeWaimaiConfirmResponseResult struct {
	Response *TradeWaimaiConfirmResponse `json:"trade_waimai_confirm_response"`
}

/* 可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单 */
type TradeWaimaiGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* true-查询仅按商家维度
false-查询按商家下所属店铺维度 */
func (r *TradeWaimaiGetRequest) SetIsAll(value string) {
	r.SetValue("is_all", value)
}

/* 返回记录数，超过20按20条返回数据<br /> 支持最大值为：20<br /> 支持最小值为：1 */
func (r *TradeWaimaiGetRequest) SetMaxSize(value string) {
	r.SetValue("max_size", value)
}

/* 外卖分店ID */
func (r *TradeWaimaiGetRequest) SetStoreId(value string) {
	r.SetValue("store_id", value)
}

func (r *TradeWaimaiGetRequest) GetResponse(accessToken string) (*TradeWaimaiGetResponse, []byte, error) {
	var resp TradeWaimaiGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.waimai.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeWaimaiGetResponse struct {
	Result *ListResult `json:"result"`
}

type TradeWaimaiGetResponseResult struct {
	Response *TradeWaimaiGetResponse `json:"trade_waimai_get_response"`
}

/* 获取淘宝外卖全部订单 */
type TradeWaimaiOrdersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* true 仅有支付宝订单,false 包括所有类型订单(货到付款,支付券等) */
func (r *TradeWaimaiOrdersGetRequest) SetIsAllOrder(value string) {
	r.SetValue("is_all_order", value)
}

/* true-查询仅按商家维度
false-查询按商家下所属店铺维度 */
func (r *TradeWaimaiOrdersGetRequest) SetIsAllShop(value string) {
	r.SetValue("is_all_shop", value)
}

/* 买家昵称/订单ID 搜索 */
func (r *TradeWaimaiOrdersGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 订单状态 待确认订单2 , 退款中订单4 , 已发货12 关闭20 交易成功21 */
func (r *TradeWaimaiOrdersGetRequest) SetOrderStatus(value string) {
	r.SetValue("order_status", value)
}

/* 页码<br /> 支持的最大列表长度为：100 */
func (r *TradeWaimaiOrdersGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回记录数，超过20按20条返回数据<br /> 支持最大值为：20<br /> 支持最小值为：1 */
func (r *TradeWaimaiOrdersGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 外卖分店ID */
func (r *TradeWaimaiOrdersGetRequest) SetShopId(value string) {
	r.SetValue("shop_id", value)
}

func (r *TradeWaimaiOrdersGetRequest) GetResponse(accessToken string) (*TradeWaimaiOrdersGetResponse, []byte, error) {
	var resp TradeWaimaiOrdersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.waimai.orders.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeWaimaiOrdersGetResponse struct {
	Result *ListResult `json:"result"`
}

type TradeWaimaiOrdersGetResponseResult struct {
	Response *TradeWaimaiOrdersGetResponse `json:"trade_waimai_orders_get_response"`
}

/* 商家拒绝订单，产生退款操作 */
type TradeWaimaiRefuseRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家付款未发货订单号 */
func (r *TradeWaimaiRefuseRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* 拒单理由 */
func (r *TradeWaimaiRefuseRequest) SetReason(value string) {
	r.SetValue("reason", value)
}

func (r *TradeWaimaiRefuseRequest) GetResponse(accessToken string) (*TradeWaimaiRefuseResponse, []byte, error) {
	var resp TradeWaimaiRefuseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trade.waimai.refuse", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TradeWaimaiRefuseResponse struct {
	RetCode string `json:"ret_code"`
}

type TradeWaimaiRefuseResponseResult struct {
	Response *TradeWaimaiRefuseResponse `json:"trade_waimai_refuse_response"`
}
