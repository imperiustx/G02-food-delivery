package notebusiness

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type GetNoteStore interface {
	FindNote(context context.Context, id int) (*notemodel.Note, error)
}

type getNoteBiz struct {
	store GetNoteStore
}

func NewGetNoteBiz(store GetNoteStore) *getNoteBiz {
	return &getNoteBiz{store: store}
}

func (biz *getNoteBiz) GetNote(context context.Context, id int) (*notemodel.Note, error) {
	note, err := biz.store.FindNote(context, id)

	if err != nil {
		return nil, common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	return note, nil
}
