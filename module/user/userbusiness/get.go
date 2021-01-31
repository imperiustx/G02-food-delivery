package userbusiness

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/user/usermodel"
)

type GetUserStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {

	user, err := biz.store.FindUser(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	return user, nil
}
