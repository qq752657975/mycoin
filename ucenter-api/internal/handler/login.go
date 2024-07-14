package handler

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"mycoin-common"
	"mycoin-common/tools"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

func NewLoginHandler(svcCtx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{
		svcCtx,
	}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//需要获取ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	resp, err := l.Login(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *LoginHandler) CheckLogin(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-auth-token")
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	resp := l.CheckLogin(token)
	result := common.NewResult().Deal(resp, nil)
	httpx.OkJsonCtx(r.Context(), w, result)
}
