package notebusiness

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type ListNoteStorage interface {
	ListNote(ctx context.Context, filter *notemodel.Filter, paging *common.Paging, moreKeys ...string) ([]notemodel.Note, error)
}

type listNote struct {
	store ListNoteStorage
}

func NewListNoteBiz(store ListNoteStorage) *listNote {
	return &listNote{store: store}
}

func (biz *listNote) ListAllNote(ctx context.Context, filter *notemodel.Filter, paging *common.Paging) ([]notemodel.Note, error) {
	data, err := biz.store.ListNote(ctx, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(notemodel.EntityName, err)
	}

	return data, nil
}
