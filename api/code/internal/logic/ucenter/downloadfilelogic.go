package ucenter

import (
	"context"
	"errors"
	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer http.ResponseWriter) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.FileShowReq) error {
	// todo: add your logic here and delete this line
	SavePath := l.svcCtx.Config.UploadFile.SavePath //上传文件的存储路径
	filePath := SavePath + req.FileUrl
	//fileName := filepath.Base(filePath)
	fileName := "go-zero.png"

	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		return errors.New("文件不存在")
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("读取文件失败")
	}
	//如果是下载，则需要在Header中设置这两个参数
	//l.writer.Header().Add("Content-Type", "application/octet-stream")
	//l.writer.Header().Add("Content-Disposition", "attachment; filename= "+fileName)

	l.writer.Header().Add("Content-Type", "application/octet-stream")
	l.writer.Header().Add("Content-Disposition", "attachment; filename= "+fileName)
	l.writer.Write(bytes)
	return nil
}
