package domain

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"mycoin-common/tools"
)

type vaptchaReq struct {
	Id        string `json:"id"`
	Secretkey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

type CaptchaDomain struct {
}

func (c CaptchaDomain) Verify(server string, vid string, key string, token string, scene int, ip string) bool {
	//发送一个post请求
	respBytes, err := tools.Post(server, &vaptchaReq{
		Id:        vid,
		Secretkey: key,
		Token:     token,
		Scene:     scene,
		Ip:        ip,
	})

	if err != nil {
		logx.Errorf("CaptchaDomain Verify post err : %s", err.Error())
		return false
	}
	var vaptchaRsp *vaptchaRsp
	err = json.Unmarshal(respBytes, &vaptchaRsp)
	if err != nil {
		logx.Errorf("CaptchaDomain Verify Unmarshal respBytes err : %s", err.Error())
		return false
	}
	if vaptchaRsp != nil && vaptchaRsp.Success == 1 {
		logx.Info("CaptchaDomain Verify no success")
		return true
	}
	return false
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}
