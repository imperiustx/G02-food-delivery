package notebusiness

import (
	"fooddlv/module/note/notemodel"
)

type ListNoteStorage interface {
	ListNote() ([]notemodel.Note, error)
}

type listNote struct {
	store ListNoteStorage
}

func NewListNoteBiz(store ListNoteStorage) *listNote {
	return &listNote{store: store}
}

func (biz *listNote) ListAllNote() ([]notemodel.Note, error) {
	return biz.store.ListNote()
}
