package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	UploadFile UploadFile
}

type UploadFile struct {
	MaxFileNum  int64
	MaxFileSize int64
	SavePath    string
}
