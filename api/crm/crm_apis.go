// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crm

import (
	"github.com/cnphpbb/open_taobao"
)

/* 卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息 */
type CrmGradeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

func (r *CrmGradeGetRequest) GetResponse(accessToken string) (*CrmGradeGetResponse, []byte, error) {
	var resp CrmGradeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grade.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGradeGetResponse struct {
	GradePromotions struct {
		GradePromotion []*GradePromotion `json:"grade_promotion"`
	} `json:"grade_promotions"`
}

type CrmGradeGetResponseResult struct {
	Response *CrmGradeGetResponse `json:"crm_grade_get_response"`
}

/* 设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。 */
type CrmGradeSetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br>
amount参数的填写规范：升级到下一个级别的需要的交易额，单位为分,必须全部填写.例如10000,20000,30000，其中10000表示非会员升级到普通的所需的交易额，20000表示普通升级到高级所需的交易额，层级等级中最高等级的下一个等级默认为0。会员等级越高，所需交易额必须越高。<br /> 支持最小值为：0 */
func (r *CrmGradeSetRequest) SetAmount(value string) {
	r.SetValue("amount", value)
}

/* 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br>
count参数的填写规范：
升级到下一个级别的需要的交易量,必须全部填写. 以逗号分隔,例如100,200,300，其中100表示非会员升级到普通会员交易量。层级等级中最高等级的下一个等级的交易量默认为0。会员等级越高，交易量必须越高。<br /> 支持最小值为：0 */
func (r *CrmGradeSetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 会员级别折扣率。会员等级越高，折扣必须越低。
950即9.5折，888折即8.88折。出于安全原因，折扣现最低只能设置到700即7折。<br /> 支持最大值为：1000<br /> 支持最小值为：700 */
func (r *CrmGradeSetRequest) SetDiscount(value string) {
	r.SetValue("discount", value)
}

/* 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP<br /> 支持最大值为：4<br /> 支持最小值为：1 */
func (r *CrmGradeSetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 是否设置达到某一会员等级的交易量和交易额，必填。4个级别都需要设置，如入参为true,true,true,false时，表示设置达到高级会员、VIP会员的交易量或者交易额，不设置达到至尊会员的交易量和交易额 */
func (r *CrmGradeSetRequest) SetHierarchy(value string) {
	r.SetValue("hierarchy", value)
}

func (r *CrmGradeSetRequest) GetResponse(accessToken string) (*CrmGradeSetResponse, []byte, error) {
	var resp CrmGradeSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grade.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGradeSetResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGradeSetResponseResult struct {
	Response *CrmGradeSetResponse `json:"crm_grade_set_response"`
}

/* 商家通过该接口吸纳线上店铺会员。 */
type CrmGrademktMemberAddRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 会员nick */
func (r *CrmGrademktMemberAddRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 系统属性，json格式 */
func (r *CrmGrademktMemberAddRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 会员属性-json format
生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberAddRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}

func (r *CrmGrademktMemberAddRequest) GetResponse(accessToken string) (*CrmGrademktMemberAddResponse, []byte, error) {
	var resp CrmGrademktMemberAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberAddResponse struct {
	Model bool `json:"model"`
}

type CrmGrademktMemberAddResponseResult struct {
	Response *CrmGrademktMemberAddResponse `json:"crm_grademkt_member_add_response"`
}

/* 创建商品等级营销明细 */
type CrmGrademktMemberDetailCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 扩展字段 */
func (r *CrmGrademktMemberDetailCreateRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberDetailCreateRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}

func (r *CrmGrademktMemberDetailCreateRequest) GetResponse(accessToken string) (*CrmGrademktMemberDetailCreateResponse, []byte, error) {
	var resp CrmGrademktMemberDetailCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.detail.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberDetailCreateResponse struct {
	Module bool `json:"module"`
}

type CrmGrademktMemberDetailCreateResponseResult struct {
	Response *CrmGrademktMemberDetailCreateResponse `json:"crm_grademkt_member_detail_create_response"`
}

/* 删除商品等级营销明细 */
type CrmGrademktMemberDetailDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 扩展字段 */
func (r *CrmGrademktMemberDetailDeleteRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberDetailDeleteRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}

func (r *CrmGrademktMemberDetailDeleteRequest) GetResponse(accessToken string) (*CrmGrademktMemberDetailDeleteResponse, []byte, error) {
	var resp CrmGrademktMemberDetailDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.detail.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberDetailDeleteResponse struct {
	Module bool `json:"module"`
}

type CrmGrademktMemberDetailDeleteResponseResult struct {
	Response *CrmGrademktMemberDetailDeleteResponse `json:"crm_grademkt_member_detail_delete_response"`
}

/* 商家通过该接口查询等级营销活动 */
type CrmGrademktMemberDetailQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 扩展字段 */
func (r *CrmGrademktMemberDetailQueryRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberDetailQueryRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}

func (r *CrmGrademktMemberDetailQueryRequest) GetResponse(accessToken string) (*CrmGrademktMemberDetailQueryResponse, []byte, error) {
	var resp CrmGrademktMemberDetailQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.detail.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberDetailQueryResponse struct {
	Model string `json:"model"`
}

type CrmGrademktMemberDetailQueryResponseResult struct {
	Response *CrmGrademktMemberDetailQueryResponse `json:"crm_grademkt_member_detail_query_response"`
}

/* 商家通过该接口设置等级活动 */
type CrmGrademktMemberGradeactivityInitRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 扩展字段 */
func (r *CrmGrademktMemberGradeactivityInitRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 活动名称，不传默认为“等级营销” */
func (r *CrmGrademktMemberGradeactivityInitRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}

func (r *CrmGrademktMemberGradeactivityInitRequest) GetResponse(accessToken string) (*CrmGrademktMemberGradeactivityInitResponse, []byte, error) {
	var resp CrmGrademktMemberGradeactivityInitResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.gradeactivity.init", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberGradeactivityInitResponse struct {
	Module bool `json:"module"`
}

type CrmGrademktMemberGradeactivityInitResponseResult struct {
	Response *CrmGrademktMemberGradeactivityInitResponse `json:"crm_grademkt_member_gradeactivity_init_response"`
}

/* 商家通过该接口查询线上店铺会员。 */
type CrmGrademktMemberQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 会员nick */
func (r *CrmGrademktMemberQueryRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 系统属性，json格式 */
func (r *CrmGrademktMemberQueryRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 会员属性-json format
生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberQueryRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}

func (r *CrmGrademktMemberQueryRequest) GetResponse(accessToken string) (*CrmGrademktMemberQueryResponse, []byte, error) {
	var resp CrmGrademktMemberQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberQueryResponse struct {
	Module string `json:"module"`
}

type CrmGrademktMemberQueryResponseResult struct {
	Response *CrmGrademktMemberQueryResponse `json:"crm_grademkt_member_query_response"`
}

/* 卖家创建一个新的分组，接口返回一个创建成功的分组的id */
type CrmGroupAddRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 分组名称，每个卖家最多可以拥有100个分组<br /> 支持最大长度为：15<br /> 支持的最大列表长度为：15 */
func (r *CrmGroupAddRequest) SetGroupName(value string) {
	r.SetValue("group_name", value)
}

func (r *CrmGroupAddRequest) GetResponse(accessToken string) (*CrmGroupAddResponse, []byte, error) {
	var resp CrmGroupAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupAddResponse struct {
	GroupId   int  `json:"group_id"`
	IsSuccess bool `json:"is_success"`
}

type CrmGroupAddResponseResult struct {
	Response *CrmGroupAddResponse `json:"crm_group_add_response"`
}

/* 将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。 */
type CrmGroupAppendRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 添加的来源分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (r *CrmGroupAppendRequest) SetFromGroupId(value string) {
	r.SetValue("from_group_id", value)
}

/* 添加的目标分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (r *CrmGroupAppendRequest) SetToGroupId(value string) {
	r.SetValue("to_group_id", value)
}

func (r *CrmGroupAppendRequest) GetResponse(accessToken string) (*CrmGroupAppendResponse, []byte, error) {
	var resp CrmGroupAppendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.append", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupAppendResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupAppendResponseResult struct {
	Response *CrmGroupAppendResponse `json:"crm_group_append_response"`
}

/* 将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。 */
type CrmGroupDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 要删除的分组id<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (r *CrmGroupDeleteRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

func (r *CrmGroupDeleteRequest) GetResponse(accessToken string) (*CrmGroupDeleteResponse, []byte, error) {
	var resp CrmGroupDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupDeleteResponseResult struct {
	Response *CrmGroupDeleteResponse `json:"crm_group_delete_response"`
}

/* 将一个分组下的所有会员移动到另一个分组，会员从原分组中删除
注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。 */
type CrmGroupMoveRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要移动的分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (r *CrmGroupMoveRequest) SetFromGroupId(value string) {
	r.SetValue("from_group_id", value)
}

/* 目的分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (r *CrmGroupMoveRequest) SetToGroupId(value string) {
	r.SetValue("to_group_id", value)
}

func (r *CrmGroupMoveRequest) GetResponse(accessToken string) (*CrmGroupMoveResponse, []byte, error) {
	var resp CrmGroupMoveResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.move", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupMoveResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupMoveResponseResult struct {
	Response *CrmGroupMoveResponse `json:"crm_group_move_response"`
}

/* 修改一个已经存在的分组，接口返回分组的修改是否成功 */
type CrmGroupUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 分组的id<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (r *CrmGroupUpdateRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 新的分组名，分组名称不能包含|或者：<br /> 支持最大长度为：15<br /> 支持的最大列表长度为：15 */
func (r *CrmGroupUpdateRequest) SetNewGroupName(value string) {
	r.SetValue("new_group_name", value)
}

func (r *CrmGroupUpdateRequest) GetResponse(accessToken string) (*CrmGroupUpdateResponse, []byte, error) {
	var resp CrmGroupUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupUpdateResponseResult struct {
	Response *CrmGroupUpdateResponse `json:"crm_group_update_response"`
}

/* 查询卖家的分组，返回查询到的分组列表，分页返回分组 */
type CrmGroupsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1<br /> 支持最大值为：1000000<br /> 支持最小值为：1<br /> 支持的最大列表长度为：3 */
func (r *CrmGroupsGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条<br /> 支持最大值为：100<br /> 支持最小值为：1<br /> 支持的最大列表长度为：3 */
func (r *CrmGroupsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

func (r *CrmGroupsGetRequest) GetResponse(accessToken string) (*CrmGroupsGetResponse, []byte, error) {
	var resp CrmGroupsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.groups.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupsGetResponse struct {
	Groups struct {
		Group []*Group `json:"group"`
	} `json:"groups"`
	TotalResult int `json:"total_result"`
}

type CrmGroupsGetResponseResult struct {
	Response *CrmGroupsGetResponse `json:"crm_groups_get_response"`
}

/* 检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组
若分组上有任务则该属性不能被操作。 */
type CrmGrouptaskCheckRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 分组id */
func (r *CrmGrouptaskCheckRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

func (r *CrmGrouptaskCheckRequest) GetResponse(accessToken string) (*CrmGrouptaskCheckResponse, []byte, error) {
	var resp CrmGrouptaskCheckResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grouptask.check", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrouptaskCheckResponse struct {
	IsFinished bool `json:"is_finished"`
}

type CrmGrouptaskCheckResponseResult struct {
	Response *CrmGrouptaskCheckResponse `json:"crm_grouptask_check_response"`
}

/* 获取买家身上的标签，不返回标签的总人数 */
type CrmMemberGroupGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 会员Nick */
func (r *CrmMemberGroupGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

func (r *CrmMemberGroupGetRequest) GetResponse(accessToken string) (*CrmMemberGroupGetResponse, []byte, error) {
	var resp CrmMemberGroupGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.member.group.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMemberGroupGetResponse struct {
	Groups struct {
		Group []*Group `json:"group"`
	} `json:"groups"`
}

type CrmMemberGroupGetResponseResult struct {
	Response *CrmMemberGroupGetResponse `json:"crm_member_group_get_response"`
}

/* 设置会员等级 */
type CrmMembergradeSetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家昵称 */
func (r *CrmMembergradeSetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 买家会员级别有四种1：普通会员。2：高级会员。 3VIP会员。 4：至尊VIP<br /> 支持最大值为：4<br /> 支持最小值为：1 */
func (r *CrmMembergradeSetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

func (r *CrmMembergradeSetRequest) GetResponse(accessToken string) (*CrmMembergradeSetResponse, []byte, error) {
	var resp CrmMembergradeSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.membergrade.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembergradeSetResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMembergradeSetResponseResult struct {
	Response *CrmMembergradeSetResponse `json:"crm_membergrade_set_response"`
}

/* 编辑会员的基本资料，接口返回会员信息修改是否成功 */
type CrmMemberinfoUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家昵称<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *CrmMemberinfoUpdateRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 城市.
请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线 */
func (r *CrmMemberinfoUpdateRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 交易关闭金额，单位：分 */
func (r *CrmMemberinfoUpdateRequest) SetCloseTradeAmount(value string) {
	r.SetValue("close_trade_amount", value)
}

/* 交易关闭次数 */
func (r *CrmMemberinfoUpdateRequest) SetCloseTradeCount(value string) {
	r.SetValue("close_trade_count", value)
}

/* 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip

只有正常会员才给予升级，对于status blacklist的会员升级无效<br /> 支持最大值为：4<br /> 支持最小值为：1<br /> 支持的最大列表长度为：32 */
func (r *CrmMemberinfoUpdateRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 分组的id集合字符串 */
func (r *CrmMemberinfoUpdateRequest) SetGroupIds(value string) {
	r.SetValue("group_ids", value)
}

/* 宝贝件数 */
func (r *CrmMemberinfoUpdateRequest) SetItemNum(value string) {
	r.SetValue("item_num", value)
}

/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.
请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线. */
func (r *CrmMemberinfoUpdateRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (r *CrmMemberinfoUpdateRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易金额，单位：分 */
func (r *CrmMemberinfoUpdateRequest) SetTradeAmount(value string) {
	r.SetValue("trade_amount", value)
}

/* 交易笔数 */
func (r *CrmMemberinfoUpdateRequest) SetTradeCount(value string) {
	r.SetValue("trade_count", value)
}

func (r *CrmMemberinfoUpdateRequest) GetResponse(accessToken string) (*CrmMemberinfoUpdateResponse, []byte, error) {
	var resp CrmMemberinfoUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.memberinfo.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMemberinfoUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMemberinfoUpdateResponseResult struct {
	Response *CrmMemberinfoUpdateResponse `json:"crm_memberinfo_update_response"`
}

/* 查询卖家的会员，进行基本的查询，返回符合条件的会员列表 */
type CrmMembersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家的昵称<br /> 支持最大长度为：1000<br /> 支持的最大列表长度为：1000 */
func (r *CrmMembersGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000<br /> 支持最大值为：1000<br /> 支持最小值为：1 */
func (r *CrmMembersGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员。如果不传入值则默认为全部等级。<br /> 支持最大值为：4<br /> 支持最小值为：-1<br /> 支持的最大列表长度为：32 */
func (r *CrmMembersGetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 最迟上次交易时间 */
func (r *CrmMembersGetRequest) SetMaxLastTradeTime(value string) {
	r.SetValue("max_last_trade_time", value)
}

/* 最大交易额，单位为元 */
func (r *CrmMembersGetRequest) SetMaxTradeAmount(value string) {
	r.SetValue("max_trade_amount", value)
}

/* 最大交易量<br /> 支持最小值为：0 */
func (r *CrmMembersGetRequest) SetMaxTradeCount(value string) {
	r.SetValue("max_trade_count", value)
}

/* 最早上次交易时间 */
func (r *CrmMembersGetRequest) SetMinLastTradeTime(value string) {
	r.SetValue("min_last_trade_time", value)
}

/* 最小交易额,单位为元 */
func (r *CrmMembersGetRequest) SetMinTradeAmount(value string) {
	r.SetValue("min_trade_amount", value)
}

/* 最小交易量<br /> 支持最小值为：0 */
func (r *CrmMembersGetRequest) SetMinTradeCount(value string) {
	r.SetValue("min_trade_count", value)
}

/* 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (r *CrmMembersGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

func (r *CrmMembersGetRequest) GetResponse(accessToken string) (*CrmMembersGetResponse, []byte, error) {
	var resp CrmMembersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersGetResponse struct {
	Members struct {
		BasicMember []*BasicMember `json:"basic_member"`
	} `json:"members"`
	TotalResult int `json:"total_result"`
}

type CrmMembersGetResponseResult struct {
	Response *CrmMembersGetResponse `json:"crm_members_get_response"`
}

/* 为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过 */
type CrmMembersGroupBatchaddRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 会员的id（一次最多传入10个）<br /> 支持最小值为：1 */
func (r *CrmMembersGroupBatchaddRequest) SetBuyerIds(value string) {
	r.SetValue("buyer_ids", value)
}

/* 分组id<br /> 支持最小值为：1 */
func (r *CrmMembersGroupBatchaddRequest) SetGroupIds(value string) {
	r.SetValue("group_ids", value)
}

func (r *CrmMembersGroupBatchaddRequest) GetResponse(accessToken string) (*CrmMembersGroupBatchaddResponse, []byte, error) {
	var resp CrmMembersGroupBatchaddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.group.batchadd", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersGroupBatchaddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMembersGroupBatchaddResponseResult struct {
	Response *CrmMembersGroupBatchaddResponse `json:"crm_members_group_batchadd_response"`
}

/* 批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。 */
type CrmMembersGroupsBatchdeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家的Id集合<br /> 支持最小值为：1 */
func (r *CrmMembersGroupsBatchdeleteRequest) SetBuyerIds(value string) {
	r.SetValue("buyer_ids", value)
}

/* 会员需要删除的分组<br /> 支持最小值为：1 */
func (r *CrmMembersGroupsBatchdeleteRequest) SetGroupIds(value string) {
	r.SetValue("group_ids", value)
}

func (r *CrmMembersGroupsBatchdeleteRequest) GetResponse(accessToken string) (*CrmMembersGroupsBatchdeleteResponse, []byte, error) {
	var resp CrmMembersGroupsBatchdeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.groups.batchdelete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersGroupsBatchdeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMembersGroupsBatchdeleteResponseResult struct {
	Response *CrmMembersGroupsBatchdeleteResponse `json:"crm_members_groups_batchdelete_response"`
}

/* 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100 */
type CrmMembersIncrementGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1<br /> 支持最小值为：1 */
func (r *CrmMembersIncrementGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间. */
func (r *CrmMembersIncrementGetRequest) SetEndModify(value string) {
	r.SetValue("end_modify", value)
}

/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员<br /> 支持最大值为：4<br /> 支持最小值为：-1<br /> 支持的最大列表长度为：32 */
func (r *CrmMembersIncrementGetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 每页显示的会员数，page_size的值不能超过100，最小值要大于1<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (r *CrmMembersIncrementGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 卖家修改会员信息的时间起点. */
func (r *CrmMembersIncrementGetRequest) SetStartModify(value string) {
	r.SetValue("start_modify", value)
}

func (r *CrmMembersIncrementGetRequest) GetResponse(accessToken string) (*CrmMembersIncrementGetResponse, []byte, error) {
	var resp CrmMembersIncrementGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.increment.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersIncrementGetResponse struct {
	Members struct {
		BasicMember []*BasicMember `json:"basic_member"`
	} `json:"members"`
	TotalResult int `json:"total_result"`
}

type CrmMembersIncrementGetResponseResult struct {
	Response *CrmMembersIncrementGetResponse `json:"crm_members_increment_get_response"`
}

/* 会员列表的高级查询，接口返回符合条件的会员列表.<br>
注：建议获取09年以后的数据，09年之前的数据不是很完整 */
type CrmMembersSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家昵称<br /> 支持最大长度为：1000<br /> 支持的最大列表长度为：1000 */
func (r *CrmMembersSearchRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 城市.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索. */
func (r *CrmMembersSearchRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页<br /> 支持最大值为：1000<br /> 支持最小值为：1 */
func (r *CrmMembersSearchRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 会员等级，0：店铺客户，1：普通客户，2：高级会员，3：VIP会员, 4：至尊VIP会员<br /> 支持最大值为：4<br /> 支持最小值为：-1<br /> 支持的最大列表长度为：32 */
func (r *CrmMembersSearchRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 分组id<br /> 支持的最大列表长度为：19 */
func (r *CrmMembersSearchRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 最大平均客单价，单位为元.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索. */
func (r *CrmMembersSearchRequest) SetMaxAvgPrice(value string) {
	r.SetValue("max_avg_price", value)
}

/* 最大交易关闭笔数.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索.<br /> 支持最小值为：0 */
func (r *CrmMembersSearchRequest) SetMaxCloseTradeNum(value string) {
	r.SetValue("max_close_trade_num", value)
}

/* 最大交易宝贝件数<br /> 支持最小值为：0 */
func (r *CrmMembersSearchRequest) SetMaxItemNum(value string) {
	r.SetValue("max_item_num", value)
}

/* 最迟上次交易时间 */
func (r *CrmMembersSearchRequest) SetMaxLastTradeTime(value string) {
	r.SetValue("max_last_trade_time", value)
}

/* 最大交易额，单位为元 */
func (r *CrmMembersSearchRequest) SetMaxTradeAmount(value string) {
	r.SetValue("max_trade_amount", value)
}

/* 最大交易量<br /> 支持最小值为：0 */
func (r *CrmMembersSearchRequest) SetMaxTradeCount(value string) {
	r.SetValue("max_trade_count", value)
}

/* 最少平均客单价，单位为元.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索. */
func (r *CrmMembersSearchRequest) SetMinAvgPrice(value string) {
	r.SetValue("min_avg_price", value)
}

/* 最小交易关闭的笔数.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索.<br /> 支持最小值为：0 */
func (r *CrmMembersSearchRequest) SetMinCloseTradeNum(value string) {
	r.SetValue("min_close_trade_num", value)
}

/* 最小交易宝贝件数<br /> 支持最小值为：0 */
func (r *CrmMembersSearchRequest) SetMinItemNum(value string) {
	r.SetValue("min_item_num", value)
}

/* 最早上次交易时间 */
func (r *CrmMembersSearchRequest) SetMinLastTradeTime(value string) {
	r.SetValue("min_last_trade_time", value)
}

/* 最小交易额，单位为元 */
func (r *CrmMembersSearchRequest) SetMinTradeAmount(value string) {
	r.SetValue("min_trade_amount", value)
}

/* 最小交易量<br /> 支持最小值为：0 */
func (r *CrmMembersSearchRequest) SetMinTradeCount(value string) {
	r.SetValue("min_trade_count", value)
}

/* 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (r *CrmMembersSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35.
注:请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.<br /> 支持最大值为：1000000<br /> 支持最小值为：1 */
func (r *CrmMembersSearchRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 关系来源，1交易成功，2未成交，3卖家手动吸纳<br /> 支持最大值为：3<br /> 支持最小值为：1<br /> 支持的最大列表长度为：32 */
func (r *CrmMembersSearchRequest) SetRelationSource(value string) {
	r.SetValue("relation_source", value)
}

func (r *CrmMembersSearchRequest) GetResponse(accessToken string) (*CrmMembersSearchResponse, []byte, error) {
	var resp CrmMembersSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersSearchResponse struct {
	Members struct {
		CrmMember []*CrmMember `json:"crm_member"`
	} `json:"members"`
	TotalResult int `json:"total_result"`
}

type CrmMembersSearchResponseResult struct {
	Response *CrmMembersSearchResponse `json:"crm_members_search_response"`
}

/* 给千牛旺旺插件单独接口获取卖家设置的标签 */
type CrmQnLabelsSellerGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

func (r *CrmQnLabelsSellerGetRequest) GetResponse(accessToken string) (*CrmQnLabelsSellerGetResponse, []byte, error) {
	var resp CrmQnLabelsSellerGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.qn.labels.seller.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmQnLabelsSellerGetResponse struct {
	Groups struct {
		Group []*Group `json:"group"`
	} `json:"groups"`
}

type CrmQnLabelsSellerGetResponseResult struct {
	Response *CrmQnLabelsSellerGetResponse `json:"crm_qn_labels_seller_get_response"`
}

/* 此接口用于取消VIP优惠 */
type CrmShopvipCancelRequest struct {
	open_taobao.TaobaoMethodRequest
}

func (r *CrmShopvipCancelRequest) GetResponse(accessToken string) (*CrmShopvipCancelResponse, []byte, error) {
	var resp CrmShopvipCancelResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.shopvip.cancel", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmShopvipCancelResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmShopvipCancelResponseResult struct {
	Response *CrmShopvipCancelResponse `json:"crm_shopvip_cancel_response"`
}

/* 获取天猫卖家设置的等级权益 */
type TmallCrmEquityGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

func (r *TmallCrmEquityGetRequest) GetResponse(accessToken string) (*TmallCrmEquityGetResponse, []byte, error) {
	var resp TmallCrmEquityGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.crm.equity.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallCrmEquityGetResponse struct {
	GradeEquitys struct {
		GradeEquity []*GradeEquity `json:"grade_equity"`
	} `json:"grade_equitys"`
}

type TmallCrmEquityGetResponseResult struct {
	Response *TmallCrmEquityGetResponse `json:"tmall_crm_equity_get_response"`
}

/* 天猫卖家设置权益 */
type TmallCrmEquitySetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 不免邮区域，只在包邮条件设置的时候生效。要和等级一一对应。包邮条件为false的时候不起作用。 */
func (r *TmallCrmEquitySetRequest) SetExcludeArea(value string) {
	r.SetValue("exclude_area", value)
}

/* 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP<br /> 支持最大值为：4<br /> 支持最小值为：1 */
func (r *TmallCrmEquitySetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 返几倍天猫积分，可以不设置。如果设置则要和等级一一对应。不设置代表清空。 */
func (r *TmallCrmEquitySetRequest) SetPoint(value string) {
	r.SetValue("point", value)
}

/* 是否包邮，可以不设置，如果设置则要和等级一一对应。不设置代表清空 */
func (r *TmallCrmEquitySetRequest) SetPostage(value string) {
	r.SetValue("postage", value)
}

func (r *TmallCrmEquitySetRequest) GetResponse(accessToken string) (*TmallCrmEquitySetResponse, []byte, error) {
	var resp TmallCrmEquitySetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.crm.equity.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallCrmEquitySetResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallCrmEquitySetResponseResult struct {
	Response *TmallCrmEquitySetResponse `json:"tmall_crm_equity_set_response"`
}
