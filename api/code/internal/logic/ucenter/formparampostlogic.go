package ucenter

import (
	"context"
	"github.com/jinzhu/copier"

	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FormParamPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFormParamPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FormParamPostLogic {
	return &FormParamPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FormParamPostLogic) FormParamPost(req *types.BaseModelForm) (resp *types.BaseModelJson, err error) {
	// todo: add your logic here and delete this line
	if req.Id > 0 {
		resp := &types.BaseModelJson{}
		copier.Copy(resp, req)
		return resp, nil
	}
	return nil, nil
}
