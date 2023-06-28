package ucenter

import (
	"context"
	"fmt"
	"go-zero-api/common/utils"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	uuid "github.com/satori/go.uuid"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(request *http.Request, req *types.FileUploadReq) (resp *types.BaseModelJson, err error) {
	// todo: add your logic here and delete this line
	SavePath := l.svcCtx.Config.UploadFile.SavePath //上传文件的存储路径
	utils.CreateDir(SavePath)
	files := request.MultipartForm.File["fileList"]
	res := &types.BaseModelJson{
		Id:   req.Id,
		Data: "操作成功",
	}
	typeId := fmt.Sprintf("%d", req.Type)
	// 遍历所有文件
	for _, fileHeader := range files {
		//获取文件大小
		fileSize := fileHeader.Size
		//获取文件名称带后缀
		fileNameWithSuffix := path.Base(fileHeader.Filename)
		//获取文件的后缀(文件类型)
		fileType := path.Ext(fileNameWithSuffix)
		//生成UUID防止文件被覆盖
		uuidName := typeId + "_" + strings.Replace(uuid.NewV4().String(), "-", "", -1)

		saveName := uuidName + fileType
		saveFullPath := SavePath + saveName
		logx.Infof("upload file: %+v, file size: %d", fileNameWithSuffix, fileSize)
		file, err := fileHeader.Open()
		tempFile, err := os.Create(saveFullPath)
		if err != nil {
			return nil, err
		}
		io.Copy(tempFile, file)
		//关闭文件
		file.Close()
		tempFile.Close()
	}
	return res, nil
}
