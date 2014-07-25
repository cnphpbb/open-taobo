// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jipiao

import (
	"github.com/cnphpbb/open_taobao"
)

/* 根据淘宝系统订单号获取订单详情信息 */
type JipiaoAgentOrderDetailRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 淘宝订单id列表，当前支持列表长度为1，即当前只支持单个订单详情查询 */
func (r *JipiaoAgentOrderDetailRequest) SetOrderIds(value string) {
	r.SetValue("order_ids", value)
}

func (r *JipiaoAgentOrderDetailRequest) GetResponse(accessToken string) (*JipiaoAgentOrderDetailResponse, []byte, error) {
	var resp JipiaoAgentOrderDetailResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.agent.order.detail", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoAgentOrderDetailResponse struct {
	ErrorMessage string `json:"error_message"`
	IsSuccess    bool   `json:"is_success"`
	Orders       struct {
		TripOrder []*TripOrder `json:"trip_order"`
	} `json:"orders"`
}

type JipiaoAgentOrderDetailResponseResult struct {
	Response *JipiaoAgentOrderDetailResponse `json:"jipiao_agent_order_detail_response"`
}

/* 获取订单对应的产品快照信息 */
type JipiaoAgentOrderProductSnapshotRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 订单号 */
func (r *JipiaoAgentOrderProductSnapshotRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

func (r *JipiaoAgentOrderProductSnapshotRequest) GetResponse(accessToken string) (*JipiaoAgentOrderProductSnapshotResponse, []byte, error) {
	var resp JipiaoAgentOrderProductSnapshotResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.agent.order.product.snapshot", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoAgentOrderProductSnapshotResponse struct {
	Data         string `json:"data"`
	ErrorMessage string `json:"error_message"`
	IsSuccess    bool   `json:"is_success"`
}

type JipiaoAgentOrderProductSnapshotResponseResult struct {
	Response *JipiaoAgentOrderProductSnapshotResponse `json:"jipiao_agent_order_product_snapshot_response"`
}

/* 卖家根据条件查询淘宝订单id列表 */
type JipiaoAgentOrderSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整） */
func (r *JipiaoAgentOrderSearchRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整） */
func (r *JipiaoAgentOrderSearchRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 是否需要行程单，true表示需要行程单；false表示不许要 */
func (r *JipiaoAgentOrderSearchRequest) SetHasItinerary(value string) {
	r.SetValue("has_itinerary", value)
}

/* 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext */
func (r *JipiaoAgentOrderSearchRequest) SetPage(value string) {
	r.SetValue("page", value)
}

/* 订单状态，默认为空，查询所有状态的订单
1:待确认
2:待出票
3:强制成功
4:未付款
5:预订成功
6:已失效 */
func (r *JipiaoAgentOrderSearchRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 航程类型： 0.单程；1.往返 */
func (r *JipiaoAgentOrderSearchRequest) SetTripType(value string) {
	r.SetValue("trip_type", value)
}

func (r *JipiaoAgentOrderSearchRequest) GetResponse(accessToken string) (*JipiaoAgentOrderSearchResponse, []byte, error) {
	var resp JipiaoAgentOrderSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.agent.order.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoAgentOrderSearchResponse struct {
	ErrorMessage string             `json:"error_message"`
	IsSuccess    bool               `json:"is_success"`
	SearchResult *SearchOrderResult `json:"search_result"`
}

type JipiaoAgentOrderSearchResponseResult struct {
	Response *JipiaoAgentOrderSearchResponse `json:"jipiao_agent_order_search_response"`
}

/* 淘宝机票代理商回填票号/成功订单 */
type JipiaoAgentOrderTicketRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 淘宝系统订单id */
func (r *JipiaoAgentOrderTicketRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* 成功订单参数：列表元素结构为——
1.航班号（注：是订单里的航班号，非实际承运航班号）;
2.旧乘机人姓名;
3.新乘机人姓名;
4.票号 （乘机人、航段对应的票号）
说明：
1.往返订单，2段航班对应1个票号的，需要2条success_info记录，分别对应去程、回程；
2.有时用户输入的乘机人姓名输入错误或者有生僻字，代理商必须输入新的名字以保证验真通过；即旧乘机人姓名和新乘机人姓名不需要变化时可以相同 */
func (r *JipiaoAgentOrderTicketRequest) SetSuccessInfo(value string) {
	r.SetValue("success_info", value)
}

func (r *JipiaoAgentOrderTicketRequest) GetResponse(accessToken string) (*JipiaoAgentOrderTicketResponse, []byte, error) {
	var resp JipiaoAgentOrderTicketResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.agent.order.ticket", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoAgentOrderTicketResponse struct {
	ErrorMessage    string `json:"error_message"`
	IsOrderSuccess  bool   `json:"is_order_success"`
	IsSuccess       bool   `json:"is_success"`
	IsTicketSuccess bool   `json:"is_ticket_success"`
}

type JipiaoAgentOrderTicketResponseResult struct {
	Response *JipiaoAgentOrderTicketResponse `json:"jipiao_agent_order_ticket_response"`
}

/* 根据条件大批量更新产品，目前只支持删除,每十分钟只允许调用一次,更新记录数不能超过10W个 */
type JipiaoPoliciesstatusUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 航空公司二字码<br /> 支持最大长度为：5<br /> 支持的最大列表长度为：5 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetAirline(value string) {
	r.SetValue("airline", value)
}

/* 到达机场三字码,此项必需与出发机场同时为空或不为空<br /> 支持最大长度为：3<br /> 支持的最大列表长度为：3 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetArrAirport(value string) {
	r.SetValue("arr_airport", value)
}

/* 出发机场三字码,此项必需与到达机场同时为空或不为空<br /> 支持最大长度为：3<br /> 支持的最大列表长度为：3 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetDepAirport(value string) {
	r.SetValue("dep_airport", value)
}

/* 外部产品id集,最多支持1000个,后续调大,其中的out_product_id含有空格将不会处理 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetOutProductIds(value string) {
	r.SetValue("out_product_ids", value)
}

/* 产品id集,最多支持1000个，后续调大，其中单个的policy_id含有留空格或不是数字将会忽略不处理 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetPolicyIds(value string) {
	r.SetValue("policy_ids", value)
}

/* 发布日期 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetPublishDate(value string) {
	r.SetValue("publish_date", value)
}

/* 发布来源, 通过接口taobao.jipiao.policy.process添加的政策会自动加上source为TOP,代理商后台页面录入的source为PC,excel上传的source为UPLOAD,通过接口taobao.jipiao.policies.fulladd,taobao.jipiao.policies.add的自定义source也可以<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 0：按policy_ids更新；1：按out_product_ids更新；2:按其它条件更新 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetType(value string) {
	r.SetValue("type", value)
}

func (r *JipiaoPoliciesstatusUpdateRequest) GetResponse(accessToken string) (*JipiaoPoliciesstatusUpdateResponse, []byte, error) {
	var resp JipiaoPoliciesstatusUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policiesstatus.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPoliciesstatusUpdateResponse struct {
	InvokeId  string `json:"invoke_id"`
	IsSuccess bool   `json:"is_success"`
}

type JipiaoPoliciesstatusUpdateResponseResult struct {
	Response *JipiaoPoliciesstatusUpdateResponse `json:"jipiao_policiesstatus_update_response"`
}

/* 国内机票代理商行程单信息回填 */
type TripJipiaoAgentItinerarySendRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 物流公司代码CODE，如不清楚，请找运营提供<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (r *TripJipiaoAgentItinerarySendRequest) SetCompanyCode(value string) {
	r.SetValue("company_code", value)
}

/* 邮寄单号，长度不能超过32字节<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *TripJipiaoAgentItinerarySendRequest) SetExpressCode(value string) {
	r.SetValue("express_code", value)
}

/* 淘宝系统行程单唯一键 */
func (r *TripJipiaoAgentItinerarySendRequest) SetItineraryId(value string) {
	r.SetValue("itinerary_id", value)
}

/* 行程单号<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *TripJipiaoAgentItinerarySendRequest) SetItineraryNo(value string) {
	r.SetValue("itinerary_no", value)
}

/* 邮寄日期，格式yyyy-mm-dd */
func (r *TripJipiaoAgentItinerarySendRequest) SetSendDate(value string) {
	r.SetValue("send_date", value)
}

func (r *TripJipiaoAgentItinerarySendRequest) GetResponse(accessToken string) (*TripJipiaoAgentItinerarySendResponse, []byte, error) {
	var resp TripJipiaoAgentItinerarySendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.itinerary.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentItinerarySendResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentItinerarySendResponseResult struct {
	Response *TripJipiaoAgentItinerarySendResponse `json:"trip_jipiao_agent_itinerary_send_response"`
}

/* 国内机票代理商确认订单接口 */
type TripJipiaoAgentOrderConfirmRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 国内机票订单id */
func (r *TripJipiaoAgentOrderConfirmRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* hk（占座）时需要的信息列表，元素结构：乘机人姓名;pnr (以分号进行分隔) */
func (r *TripJipiaoAgentOrderConfirmRequest) SetPnrInfo(value string) {
	r.SetValue("pnr_info", value)
}

func (r *TripJipiaoAgentOrderConfirmRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderConfirmResponse, []byte, error) {
	var resp TripJipiaoAgentOrderConfirmResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.confirm", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderConfirmResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderConfirmResponseResult struct {
	Response *TripJipiaoAgentOrderConfirmResponse `json:"trip_jipiao_agent_order_confirm_response"`
}

/* 国内机票代理商失败订单接口 */
type TripJipiaoAgentOrderFailRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 失败类型为0，说明备注原因<br /> 支持最大长度为：200<br /> 支持的最大列表长度为：200 */
func (r *TripJipiaoAgentOrderFailRequest) SetFailMemo(value string) {
	r.SetValue("fail_memo", value)
}

/* 失败原因：1．客户要求失败订单；2．此舱位已售完（经济舱或特舱）；3．剩余座位少于用户购买数量；4．特价管理里录入的特价票已经售完；5．假舱（请及时通过旺旺或者电话反馈给淘宝工作人员）；0．其它（请在备注中说明原因） */
func (r *TripJipiaoAgentOrderFailRequest) SetFailType(value string) {
	r.SetValue("fail_type", value)
}

/* 国内机票订单id */
func (r *TripJipiaoAgentOrderFailRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

func (r *TripJipiaoAgentOrderFailRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderFailResponse, []byte, error) {
	var resp TripJipiaoAgentOrderFailResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.fail", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderFailResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderFailResponseResult struct {
	Response *TripJipiaoAgentOrderFailResponse `json:"trip_jipiao_agent_order_fail_response"`
}

/* 根据淘宝机票政策id搜索订单 */
type TripJipiaoAgentOrderFindRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderFindRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderFindRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext */
func (r *TripJipiaoAgentOrderFindRequest) SetPage(value string) {
	r.SetValue("page", value)
}

/* 淘宝机票政策id */
func (r *TripJipiaoAgentOrderFindRequest) SetPolicyId(value string) {
	r.SetValue("policy_id", value)
}

func (r *TripJipiaoAgentOrderFindRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderFindResponse, []byte, error) {
	var resp TripJipiaoAgentOrderFindResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.find", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderFindResponse struct {
	SearchResult *SearchOrderResult `json:"search_result"`
}

type TripJipiaoAgentOrderFindResponseResult struct {
	Response *TripJipiaoAgentOrderFindResponse `json:"trip_jipiao_agent_order_find_response"`
}

/* 根据淘宝系统订单号获取订单详情信息 */
type TripJipiaoAgentOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 淘宝政策id列表，当前支持列表长度为1，即当前只支持单个订单详情查询 */
func (r *TripJipiaoAgentOrderGetRequest) SetOrderIds(value string) {
	r.SetValue("order_ids", value)
}

func (r *TripJipiaoAgentOrderGetRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderGetResponse, []byte, error) {
	var resp TripJipiaoAgentOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderGetResponse struct {
	Orders struct {
		AtOrder []*AtOrder `json:"at_order"`
	} `json:"orders"`
}

type TripJipiaoAgentOrderGetResponseResult struct {
	Response *TripJipiaoAgentOrderGetResponse `json:"trip_jipiao_agent_order_get_response"`
}

/* 国内机票代理商手工hk订单（未付款前，手工填写pnr） */
type TripJipiaoAgentOrderHkRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 国内机票订单id */
func (r *TripJipiaoAgentOrderHkRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* hk（占座）时需要的信息信息列表，元素结构：乘机人姓名;pnr (以分号进行分隔). */
func (r *TripJipiaoAgentOrderHkRequest) SetPnrInfo(value string) {
	r.SetValue("pnr_info", value)
}

func (r *TripJipiaoAgentOrderHkRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderHkResponse, []byte, error) {
	var resp TripJipiaoAgentOrderHkResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.hk", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderHkResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderHkResponseResult struct {
	Response *TripJipiaoAgentOrderHkResponse `json:"trip_jipiao_agent_order_hk_response"`
}

/* 根据条件查询淘宝订单id列表 */
type TripJipiaoAgentOrderSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderSearchRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderSearchRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 是否需要行程单，true表示需要行程单；false表示不许要 （必须设置，且不能为null） */
func (r *TripJipiaoAgentOrderSearchRequest) SetHasItinerary(value string) {
	r.SetValue("has_itinerary", value)
}

/* 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext */
func (r *TripJipiaoAgentOrderSearchRequest) SetPage(value string) {
	r.SetValue("page", value)
}

/* 订单状态，默认为空，查询所有状态的订单 */
func (r *TripJipiaoAgentOrderSearchRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 航程类型： 0.单程和普通往返；2.多程（暂时没有使用）；3.特价往返 */
func (r *TripJipiaoAgentOrderSearchRequest) SetTripType(value string) {
	r.SetValue("trip_type", value)
}

func (r *TripJipiaoAgentOrderSearchRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderSearchResponse, []byte, error) {
	var resp TripJipiaoAgentOrderSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderSearchResponse struct {
	SearchResult *SearchOrderResult `json:"search_result"`
}

type TripJipiaoAgentOrderSearchResponseResult struct {
	Response *TripJipiaoAgentOrderSearchResponse `json:"trip_jipiao_agent_order_search_response"`
}

/* 淘宝机票代理商成功/解冻订单 */
type TripJipiaoAgentOrderSuccessRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 淘宝系统订单id */
func (r *TripJipiaoAgentOrderSuccessRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* 成功订单参数：列表元素结构为——旧乘机人姓名;新乘机人姓名;pnr;票号 (以分号进行分隔)。说明：有时用户输入的乘机人姓名输入错误或者有生僻字，代理商必须输入新的名字以保证验真通过；即旧乘机人姓名和新乘机人姓名不需要变化时可以相同 */
func (r *TripJipiaoAgentOrderSuccessRequest) SetSuccessInfo(value string) {
	r.SetValue("success_info", value)
}

func (r *TripJipiaoAgentOrderSuccessRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderSuccessResponse, []byte, error) {
	var resp TripJipiaoAgentOrderSuccessResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.success", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderSuccessResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderSuccessResponseResult struct {
	Response *TripJipiaoAgentOrderSuccessResponse `json:"trip_jipiao_agent_order_success_response"`
}
