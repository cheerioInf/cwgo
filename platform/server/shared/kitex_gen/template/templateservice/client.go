// Code generated by Kitex v0.6.1. DO NOT EDIT.

package templateservice

import (
	"context"
	template "github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/template"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddTemplate(ctx context.Context, req *template.AddTemplateReq, callOptions ...callopt.Option) (r *template.AddTemplateRes, err error)
	DeleteTemplate(ctx context.Context, req *template.DeleteTemplateReq, callOptions ...callopt.Option) (r *template.DeleteTemplateRes, err error)
	UpdateTemplate(ctx context.Context, req *template.UpdateTemplateReq, callOptions ...callopt.Option) (r *template.UpdateTemplateRes, err error)
	GetTemplates(ctx context.Context, req *template.GetTemplateItemsReq, callOptions ...callopt.Option) (r *template.GetTemplatesRes, err error)
	AddTemplateItem(ctx context.Context, req *template.AddTemplateItemReq, callOptions ...callopt.Option) (r *template.AddTemplateItemRes, err error)
	DeleteTemplateItem(ctx context.Context, req *template.DeleteTemplateItemReq, callOptions ...callopt.Option) (r *template.DeleteTemplateRes, err error)
	UpdateTemplateItem(ctx context.Context, req *template.UpdateTemplateItemReq, callOptions ...callopt.Option) (r *template.UpdateTemplateItemRes, err error)
	GetTemplateItems(ctx context.Context, req *template.GetTemplatesReq, callOptions ...callopt.Option) (r *template.GetTemplateItemsRes, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTemplateServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTemplateServiceClient struct {
	*kClient
}

func (p *kTemplateServiceClient) AddTemplate(ctx context.Context, req *template.AddTemplateReq, callOptions ...callopt.Option) (r *template.AddTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTemplate(ctx, req)
}

func (p *kTemplateServiceClient) DeleteTemplate(ctx context.Context, req *template.DeleteTemplateReq, callOptions ...callopt.Option) (r *template.DeleteTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTemplate(ctx, req)
}

func (p *kTemplateServiceClient) UpdateTemplate(ctx context.Context, req *template.UpdateTemplateReq, callOptions ...callopt.Option) (r *template.UpdateTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTemplate(ctx, req)
}

func (p *kTemplateServiceClient) GetTemplates(ctx context.Context, req *template.GetTemplateItemsReq, callOptions ...callopt.Option) (r *template.GetTemplatesRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTemplates(ctx, req)
}

func (p *kTemplateServiceClient) AddTemplateItem(ctx context.Context, req *template.AddTemplateItemReq, callOptions ...callopt.Option) (r *template.AddTemplateItemRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTemplateItem(ctx, req)
}

func (p *kTemplateServiceClient) DeleteTemplateItem(ctx context.Context, req *template.DeleteTemplateItemReq, callOptions ...callopt.Option) (r *template.DeleteTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTemplateItem(ctx, req)
}

func (p *kTemplateServiceClient) UpdateTemplateItem(ctx context.Context, req *template.UpdateTemplateItemReq, callOptions ...callopt.Option) (r *template.UpdateTemplateItemRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTemplateItem(ctx, req)
}

func (p *kTemplateServiceClient) GetTemplateItems(ctx context.Context, req *template.GetTemplatesReq, callOptions ...callopt.Option) (r *template.GetTemplateItemsRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTemplateItems(ctx, req)
}