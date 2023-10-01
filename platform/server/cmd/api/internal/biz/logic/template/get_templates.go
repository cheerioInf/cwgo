/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package template

import (
	"context"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/biz/model/template"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
)

const (
	successMsgGetTemplates = "" // TODO: to be filled...
)

type GetTemplatesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplatesLogic {
	return &GetTemplatesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplatesLogic) GetTemplates(req *template.GetTemplatesReq) (res *template.GetTemplatesRes) {
	// TODO: to be filled...

	return &template.GetTemplatesRes{
		Code: 0,
		Msg:  successMsgGetTemplates,
	}
}