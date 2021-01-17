package notebusiness

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type CreateNoteStore interface {
	CreateNote(context context.Context, data *notemodel.NoteCreate) error
}

type createNoteBiz struct {
	store CreateNoteStore
}

func NewCreateNoteBiz(store CreateNoteStore) *createNoteBiz {
	return &createNoteBiz{store: store}
}

func (biz *createNoteBiz) CreateNewNote(context context.Context, data *notemodel.NoteCreate) error {
	if err := biz.store.CreateNote(context, data); err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	return nil
}
