package tbsdk

////////////////////////////////////////////////////////////////////////////////
// (生成淘口令) 提供isv生成淘口令接口，isv提交口令内容、logo、url等参数，生成淘口令关键key如：￥SADadW￥，后续进行文案包装组装用于传播
// http://open.taobao.com/doc2/apiDetail.htm?apiId=26520&scopeId=11998
type OpenTbkTpwdParam struct {
	// 必须参数
	text            string    			`json:"text"`                          // 必须 口令弹框内容
	url             string                    	`json:"url"`                           // 必须 口令跳转目标页

	// 可选参数
	user_id         string                    	`json:"user_id"`           // 可选 生成口令的淘宝用户ID
	logo            string                    	`json:"logo"`            // 可选 口令弹框logoURL
	ext           	string                  	`json:"ext"`            // 可选 扩展字段JSON格式
}

func (this OpenTbkTpwdParam) APIName() string {
	return "taobao.wireless.share.tpwd.create"
}

func (this OpenTbkTpwdParam) Params() map[string]string {
	return nil
}

func (this OpenTbkTpwdParam) ExtJSONParamName() string {
	return "wireless_share_tpwd_create_response"
}

func (this OpenTbkTpwdParam) ExtJSONParamValue() string {
	return marshal(this)
}
