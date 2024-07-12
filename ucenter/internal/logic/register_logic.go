package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mycoin-common/tools"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER:"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	//校验人机是否通过,远程调用没有注册，先注释
	//isVerify := l.CaptchaDomain.Verify(
	//	in.Captcha.Server,
	//	l.svcCtx.Config.Captcha.Vid,
	//	l.svcCtx.Config.Captcha.Key,
	//	in.Captcha.Token,
	//	2,
	//	in.Ip)
	//
	//if !isVerify {
	//	return nil, errors.New("人机校验不通过")
	//}
	logx.Info("人机验证通过")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	//校验验证码
	redisCode := ""
	err := l.svcCtx.Cache.GetCtx(ctx, RegisterCacheKey+in.Phone, &redisCode)
	if err != nil {
		return nil, errors.New("验证码不可用或者验证码已过期")
	}
	if in.Code != redisCode {
		return nil, errors.New("验证码不正确")
	}
	//检查手机号是否注册
	mem, err := l.MemberDomain.FindByPhone(ctx, in.Phone)
	if err != nil {
		return nil, errors.New("服务异常，请联系管理员")
	}
	if mem != nil {
		return nil, errors.New("手机号已经被注册")
	}
	//存入数据库
	err = l.MemberDomain.Register(
		ctx,
		in.Username,
		in.Phone,
		in.Password,
		in.Country,
		in.Promotion,
		in.SuperPartner,
	)
	if err != nil {
		return nil, errors.New("注册失败")
	}

	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(req *register.CodeReq) (*register.NoRes, error) {
	code := tools.Rand4Num()
	//假设调用短信平台发送验证码成功
	go func() {
		logx.Info("调用短信平台发送验证码成功")
	}()
	logx.Infof("验证码为: %s \n", code)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 5*time.Minute)
	if err != nil {
		return nil, errors.New("验证码存入cache失败")
	}
	return &register.NoRes{}, nil
}
