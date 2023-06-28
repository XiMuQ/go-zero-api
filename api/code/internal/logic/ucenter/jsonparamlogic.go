package ucenter

import (
	"context"
	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
)

type JsonParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JsonParamLogic {
	return &JsonParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JsonParamLogic) JsonParam(req *types.BaseModelJson) (resp *types.BaseModelJson, err error) {
	// todo: add your logic here and delete this line
	if req.Id > 0 {
		resp := &types.BaseModelJson{}
		copier.Copy(resp, req)
		return resp, nil
	}
	return nil, nil
}
