package notebusiness

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type UpdateNoteStore interface {
	FindNote(context context.Context, id int) (*notemodel.Note, error)
	Update(ctx context.Context, id int, data *notemodel.NoteUpdate) error
}

type updateNoteBiz struct {
	store UpdateNoteStore
}

func NewUpdateNoteBiz(store UpdateNoteStore) *updateNoteBiz {
	return &updateNoteBiz{store: store}
}

func (biz *updateNoteBiz) UpdateNote(context context.Context, id int, data *notemodel.NoteUpdate) error {
	note, err := biz.store.FindNote(context, id)

	if err != nil {
		return common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	if note.Status == 0 {
		return common.ErrCannotGetEntity(notemodel.EntityName, errors.New("note note found"))
	}

	if err := biz.store.Update(context, note.Id, data); err != nil {
		return common.ErrCannotUpdateEntity(notemodel.EntityName, err)
	}

	return nil
}
