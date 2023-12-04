package api

import (
	"EdgeSys/pkg/tool"
	"fmt"
	"net/http"
	"os"
	"path"

	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
)

type UploadApi struct{}

const filePath = "uploads/file"

// UploadImage 图片上传
func (up *UploadApi) UploadImage(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("file")
	biz.ErrIsNil(err, "请传入文件")
	local := &tool.Local{Path: filePath}
	link, fileName, err := local.UploadFile(fileHeader)
	biz.ErrIsNil(err, "文件上传失败")
	rc.ResData = map[string]string{"fileName": fileName, "filePath": link}
}

func (up *UploadApi) GetImage(rc *restfulx.ReqCtx) {
	actual := path.Join(filePath, restfulx.PathParam(rc, "subpath"))
	http.ServeFile(
		rc.Response.ResponseWriter,
		rc.Request.Request,
		actual)
}

func (up *UploadApi) DeleteImage(rc *restfulx.ReqCtx) {
	fileName := restfulx.QueryParam(rc, "fileName")
	biz.NotEmpty(fileName, "请传要删除的图片名")
	err := os.Remove(fmt.Sprintf("%s/%s", filePath, fileName))
	biz.ErrIsNil(err, "文件删除失败")
}
