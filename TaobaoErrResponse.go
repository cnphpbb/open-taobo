// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"strconv"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

type TaobaoTopErrResponse struct {
	Error            string       `json:"error"`
	ErrorDescription string       `json:"error_description"`
	ErrResponse      *ErrResponse `json:"error_response"`
}

func (t *TaobaoTopErrResponse) GetErr() string {
	if t.Error != "" {
		return "[" + t.Error + "]" + t.ErrorDescription
	}
	if t.ErrResponse != nil {
		if t.ErrResponse.Code > 0 {
			return "code:: [" + strconv.Itoa(t.ErrResponse.Code) + "] Msg:: " + t.ErrResponse.Msg + " sub_code:: " + t.ErrResponse.SubCode + " sub_msg:: " + t.ErrResponse.SubMsg
		}
	}
	return ""
}
