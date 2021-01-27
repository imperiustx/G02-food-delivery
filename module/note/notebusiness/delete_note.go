package notebusiness

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type DeleteNoteStore interface {
	FindNote(context context.Context, id int) (*notemodel.Note, error)
	DeleteNote(id int) error
}

type deleteNoteBiz struct {
	store     DeleteNoteStore
	requester common.Requester
}

func NewDeleteNoteBiz(store DeleteNoteStore, requester common.Requester) *deleteNoteBiz {
	return &deleteNoteBiz{store: store, requester: requester}
}

func (biz *deleteNoteBiz) DeleteNote(context context.Context, id int) error {
	note, err := biz.store.FindNote(context, id)

	if err != nil {
		return common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return common.ErrCannotGetEntity(notemodel.EntityName, errors.New("note note found"))
	}

	isAuthor := biz.requester.GetUserId() == note.UserId
	isAdmin := biz.requester.GetRole() == "admin"

	if !isAuthor && !isAdmin {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.DeleteNote(id); err != nil {
		return common.ErrCannotDeleteEntity(notemodel.EntityName, err)
	}

	return nil
}
