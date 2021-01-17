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
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(context context.Context, id int) error {
	note, err := biz.store.FindNote(context, id)

	if err != nil {
		return common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return common.ErrCannotGetEntity(notemodel.EntityName, errors.New("note note found"))
	}

	if err := biz.store.DeleteNote(id); err != nil {
		return err
	}

	return nil
}
