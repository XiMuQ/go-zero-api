package ucenter

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-api/common/errorx"

	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PathParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPathParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PathParamLogic {
	return &PathParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PathParamLogic) PathParam(req *types.PathReq) (resp *types.BaseModelJson, err error) {
	// todo: add your logic here and delete this line
	if req.Id > 0 {
		resp := &types.BaseModelJson{}
		copier.Copy(resp, req)
		return resp, nil
	} else {
		return nil, errorx.NewDefaultError(errorx.ParamErrorCode)
	}
	return nil, nil
}
