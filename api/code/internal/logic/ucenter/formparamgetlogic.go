package ucenter

import (
	"context"
	"github.com/jinzhu/copier"

	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FormParamGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFormParamGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FormParamGetLogic {
	return &FormParamGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FormParamGetLogic) FormParamGet(req *types.BaseModelForm) (resp *types.BaseModelJson, err error) {
	// todo: add your logic here and delete this line
	if req.Id > 0 {
		resp := &types.BaseModelJson{}
		copier.Copy(resp, req)
		return resp, nil
	}
	return nil, nil
}
