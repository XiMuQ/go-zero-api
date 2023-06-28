// Code generated by goctl. DO NOT EDIT.
package types

type BaseModelJson struct {
	Id   int64  `json:"id"`            // id
	Name string `json:"name,optional"` // 名称
	Data string `json:"data,optional"` // 数据
}

type BaseModelForm struct {
	Id   int64  `form:"id"`            // id
	Name string `form:"name,optional"` // 名称
	Data string `form:"data,optional"` // 数据
}

type PathReq struct {
	Id   int64  `path:"id"`   // id
	Name string `path:"name"` // 名称
}

type FileUploadReq struct {
	Id       int64   `form:"id"`                // 父级-id
	Type     int64   `form:"type,optional"`     // 类型 1：证书文件；2：私钥文件
	FileList []*byte `form:"fileList,optional"` // 文件列表
}

type FileShowReq struct {
	Id      int64  `form:"id"`               // 文件-id
	FileUrl string `form:"fileUrl,optional"` // 文件地址
}

type UserModel struct {
	Account  string `json:"account"`                      // 账号
	Password string `json:"password"`                     // 密码
	Gender   int64  `json:"gender,options=1:2:3"`         // 性别 1：未设置；2：男性；3：女性
	Current  int64  `json:"current,optional,default=1"`   // 当前页码
	PageSize int64  `json:"pageSize,optional,default=20"` // 每页数量
}

type UserLoginResp struct {
	Id           int64  `json:"id"`           // 用户id
	Account      string `json:"account"`      // 账号
	Username     string `json:"username"`     // 登录账号
	Gender       int64  `json:"gender"`       // 性别 1：未设置；2：男性；3：女性
	Avatar       string `json:"avatar"`       // 头像
	AccessToken  string `json:"token"`        // token
	AccessExpire int64  `json:"accessExpire"` // 过期时间戳
	RefreshAfter int64  `json:"refreshAfter"` //
}