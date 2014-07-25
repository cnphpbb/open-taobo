// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package widget

import (
	"github.com/cnphpbb/open_taobao"
)

/* 获取当前浏览器用户在淘宝登陆状态。如果传了session并且此用户已在淘宝登陆，返回nick和userId。仅支持widget入口调用。调用入口为/widget。签名方法简化为Hmac-md5,hmac(secret+‘app_key' ＋app_key +'timestamp' + timestamp+secret, secret)。timestamp为60分钟内有效 */
type WidgetLoginstatusGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 指定判断当前浏览器登陆用户是否此昵称用户，并且返回是否登陆。如果用户不一致返回未登陆，如果用户一致且已登录返回已登陆 */
func (r *WidgetLoginstatusGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

func (r *WidgetLoginstatusGetRequest) GetResponse(accessToken string) (*WidgetLoginstatusGetResponse, []byte, error) {
	var resp WidgetLoginstatusGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.widget.loginstatus.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WidgetLoginstatusGetResponse struct {
	IsLogin bool   `json:"is_login"`
	Nick    string `json:"nick"`
	UserId  string `json:"user_id"`
}

type WidgetLoginstatusGetResponseResult struct {
	Response *WidgetLoginstatusGetResponse `json:"widget_loginstatus_get_response"`
}
