package snippet

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler/snippethandle"
)

func GetSnippetList(ctx *gin.Context) {
	var req snippethandle.GetSnippetListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		Biu.Failed(ctx, constance.ParamErr, constance.ParamErrMsg)
		return
	}

	resp, code, msg, err := snippethandle.GetSnippetList(ctx, req)
	if err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, constance.ParamServerErrMsg)
		return
	}

	Biu.ReturnList(ctx, code, msg, resp)
}
