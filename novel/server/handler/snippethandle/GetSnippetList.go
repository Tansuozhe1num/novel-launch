package snippethandle

import (
	"github.com/gin-gonic/gin"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/middleware/db"
	models "novel-launch/novel/modells/db"
	"novel-launch/novel/server/dao"
)

type GetSnippetListRequest struct {
	Size int `form:"size" json:"size"`
}

func GetSnippetList(ctx *gin.Context, req GetSnippetListRequest) ([]models.FunnySnippet, int, string, error) {
	snippetDAO := dao.NewFunnySnippetDAO(db.Get())

	snippets, err := snippetDAO.GetRandomSnippetList(ctx, req.Size)
	if err != nil {
		return nil, constance.ParamDBErr, constance.ParamDBErrMsg, err
	}

	return snippets, 0, "", nil
}
