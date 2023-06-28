package ucenter

import (
	"context"
	"errors"
	"net/http"
	"os"

	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewPreviewFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer http.ResponseWriter) *PreviewFileLogic {
	return &PreviewFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *PreviewFileLogic) PreviewFile(req *types.FileShowReq) error {
	// todo: add your logic here and delete this line
	SavePath := l.svcCtx.Config.UploadFile.SavePath //上传文件的存储路径
	filePath := SavePath + req.FileUrl

	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		return errors.New("文件不存在")
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("读取文件失败")
	}
	l.writer.Write(bytes)
	return nil
}
