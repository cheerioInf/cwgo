// Code generated by Kitex v0.6.1. DO NOT EDIT.

package baseservice

import (
	"context"
	base "github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/base"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return baseServiceServiceInfo
}

var baseServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BaseService"
	handlerType := (*base.BaseService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newBaseServiceRegisterArgs, newBaseServiceRegisterResult, false),
		"Login":    kitex.NewMethodInfo(loginHandler, newBaseServiceLoginArgs, newBaseServiceLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "base",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServiceRegisterArgs)
	realResult := result.(*base.BaseServiceRegisterResult)
	success, err := handler.(base.BaseService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServiceRegisterArgs() interface{} {
	return base.NewBaseServiceRegisterArgs()
}

func newBaseServiceRegisterResult() interface{} {
	return base.NewBaseServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServiceLoginArgs)
	realResult := result.(*base.BaseServiceLoginResult)
	success, err := handler.(base.BaseService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServiceLoginArgs() interface{} {
	return base.NewBaseServiceLoginArgs()
}

func newBaseServiceLoginResult() interface{} {
	return base.NewBaseServiceLoginResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *base.RegisterReq) (r *base.RegisterRes, err error) {
	var _args base.BaseServiceRegisterArgs
	_args.Req = req
	var _result base.BaseServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *base.LoginReq) (r *base.LoginRes, err error) {
	var _args base.BaseServiceLoginArgs
	_args.Req = req
	var _result base.BaseServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}