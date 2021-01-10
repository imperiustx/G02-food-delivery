package notebusiness

import (
	"errors"
	"fooddlv/module/note/notemodel"
)

type DeleteNoteStore interface {
	FindNote(id int) (*notemodel.Note, error)
	DeleteNote(id int) error
}

type deleteNoteBiz struct {
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(id int) error {
	note, err := biz.store.FindNote(id)

	//if note == nil {
	//	return errors.New("note not found")
	//}

	if err != nil {
		return err
	}

	if note.Status == 0 {
		return errors.New("note note found")
	}

	if err := biz.store.DeleteNote(id); err != nil {
		return err
	}

	return nil
}
