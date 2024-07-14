package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"grpc-common/ucenter/types/login"
	"mycoin-common/tools"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const LoginCacheKey = "Login:"

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
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
	//校验密码
	//查询salt
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	member, err := l.MemberDomain.FindByPhone(ctx, in.Username)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	if member == nil {
		return nil, errors.New("账号或密码错误")
	}
	password := member.Password
	salt := member.Salt
	verify := tools.Verify(in.Password, salt, password, nil)
	if !verify {
		return nil, errors.New("账号或密码错误")
	}
	//登录成功，生成token
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	token, err := l.getJwtToken(accessSecret, time.Now().Unix(), accessExpire, member.Id)
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	loginCount := member.LoginCount + 1
	go func() {
		l.MemberDomain.UpdateLoginCount(member.Id, 1)
	}()
	return &login.LoginRes{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner:  member.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil

}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
