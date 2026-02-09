package category

import (
	"context"
	constance "novel-launch/novel/Biu/const"
	models "novel-launch/novel/modells/db"
	"novel-launch/novel/server/dao"
)

func GetDefaultCategories(ctx context.Context) (code int, msg string, categories []models.Category, err error) {
	categories, err = dao.GetDefaultCategories(ctx)
	if err != nil {
		return constance.ParamServerErr, constance.ParamServerErrMsg, nil, err
	}

	return constance.Success, constance.SuccessMsg, categories, nil
}
