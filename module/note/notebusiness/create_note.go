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
	store     CreateNoteStore
	requester common.Requester
}

func NewCreateNoteBiz(store CreateNoteStore, requester common.Requester) *createNoteBiz {
	return &createNoteBiz{store: store, requester: requester}
}

func (biz *createNoteBiz) CreateNewNote(context context.Context, data *notemodel.NoteCreate) error {
	data.UserId = biz.requester.GetUserId()

	if err := biz.store.CreateNote(context, data); err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	return nil
}
